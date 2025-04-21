package login_service_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/jwts"
	"hnz.com/ms_serve/ms-common/times"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-user/config"
	"hnz.com/ms_serve/ms-user/internal/dao"
	"hnz.com/ms_serve/ms-user/internal/data/member"
	"hnz.com/ms_serve/ms-user/internal/data/organization"
	"hnz.com/ms_serve/ms-user/internal/database"
	"hnz.com/ms_serve/ms-user/internal/database/tran"
	"hnz.com/ms_serve/ms-user/internal/repo"
	"hnz.com/ms_serve/ms-user/pkg/model"
	"log"
	"strconv"
	"strings"
	"time"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache            repo.Cache
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
	transaction      tran.Transaction
}

func New() *LoginService {
	return &LoginService{
		cache:            dao.Rc,
		memberRepo:       dao.NewMemberDao(),
		organizationRepo: dao.NewOrganizationDao(),
		transaction:      dao.NewTrans(),
	}
}

func (l *LoginService) GetCaptcha(ctx context.Context, msg *login.CaptchaMessage) (*login.CaptchaResponse, error) {
	mobile := msg.Mobile
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	// 生成验证码
	code := "123456"
	go func() {
		// 发送验证码
		time.Sleep(time.Second * 2)
		zap.L().Info("验证码发送成功！")
		c, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		err := l.cache.Put(c, model.RegisterKey+mobile, code, time.Minute*5)
		if err != nil {
			log.Println("验证码存入redis失败！", err)
		}
		log.Printf("手机号和验证码存入到redis中：REGISTER_%s:%s", mobile, code)
	}()
	return &login.CaptchaResponse{Code: code}, nil
}

func (l *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	// 检验参数
	// 校验验证码
	redisCode, err := l.cache.Get(ctx, model.RegisterKey+msg.Mobile)
	if errors.Is(err, redis.Nil) {
		return nil, errs.GrpcError(model.CaptchaNotFound)
	}
	if err != nil {
		zap.L().Error("验证码校验失败！", zap.Error(err))
		return nil, errs.GrpcError(model.RedisError)
	}
	if redisCode != msg.Captcha {
		return nil, errs.GrpcError(model.CaptchaError)
	}
	// 校验业务逻辑（邮箱是否注册、账号是否注册、手机号是否注册）
	exist, err := l.memberRepo.GetMemberByEmail(ctx, msg.Email)
	if err != nil {
		zap.L().Error("register database get error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if exist {
		return nil, errs.GrpcError(model.EmailExist)
	}
	exist, err = l.memberRepo.GetMemberByAccount(ctx, msg.Name)
	if err != nil {
		zap.L().Error("register database get error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if exist {
		return nil, errs.GrpcError(model.NameExist)
	}
	exist, err = l.memberRepo.GetMemberByMobile(ctx, msg.Mobile)
	if err != nil {
		zap.L().Error("register database get error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if exist {
		return nil, errs.GrpcError(model.MobileExist)
	}
	// 执行业务 -- 将数据插入到member表中 生成一个数据存入到组织表organization中
	pwd := encrypts.Md5(msg.Password)
	mem := &member.Member{
		Account:       msg.Name,
		Password:      pwd,
		Name:          msg.Name,
		Mobile:        msg.Mobile,
		Email:         msg.Email,
		CreateTime:    time.Now().UnixMilli(),
		Status:        model.Normal,
		LastLoginTime: time.Now().UnixMilli(),
	}
	// 事务
	err = l.transaction.Action(func(conn database.DBConn) error {
		err = l.memberRepo.SaveMember(conn, ctx, mem)
		if err != nil {
			zap.L().Error("register database save error！", zap.Error(err))
			return errs.GrpcError(model.DataBaseError)
		}
		// 存入组织
		org := &organization.Organization{
			Name:       mem.Name + "个人组织",
			MemberId:   mem.Id,
			CreateTime: time.Now().UnixMilli(),
			Personal:   int32(model.Personal),
			Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
		}
		err = l.organizationRepo.SaveOrganization(conn, ctx, org)
		if err != nil {
			zap.L().Error("register SaveOrganization db err", zap.Error(err))
			return errs.GrpcError(model.DataBaseError)
		}
		return nil
	})
	// 返回结果
	return &login.RegisterResponse{}, err
}
func (l *LoginService) Login(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	c := context.Background()
	// 数据库查询账号密码是否正确
	pwd := encrypts.Md5(msg.Password)
	mem, err := l.memberRepo.FindMember(c, msg.Account, pwd)
	log.Printf("msg: %v", msg)
	if err != nil {
		zap.L().Error("login database get error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if mem == nil {
		return nil, errs.GrpcError(model.AccountOrPasswordError)
	}
	membMessage := &login.MemberMessage{}
	err = copier.Copy(membMessage, mem)
	membMessage.Code, err = encrypts.Encrypt(strconv.FormatInt(mem.Id, 10), model.AESKey)
	membMessage.LastLoginTime = times.FormatByMill(mem.LastLoginTime)
	membMessage.CreateTime = times.FormatByMill(mem.CreateTime)
	// 根据用户id 查询用户信息--组织
	orgs, err := l.organizationRepo.FindOrganizationByMemId(c, mem.Id)
	if err != nil {
		zap.L().Error("login database get error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	for _, v := range orgsMessage {
		v.Code, _ = encrypts.Encrypt(strconv.FormatInt(v.Id, 10), model.AESKey)
		v.OwnerCode = membMessage.Code
		v.CreateTime = times.FormatByMill(organization.ToMap(orgs)[v.Id].CreateTime)
	}
	if len(orgs) > 0 {
		membMessage.OrganizationCode, _ = encrypts.Encrypt(strconv.FormatInt(orgs[0].Id, 10), model.AESKey)
	}
	// 使用jwt生成token
	memIdStr := strconv.FormatInt(mem.Id, 10)
	accessExp := time.Duration(config.Cfg.Jc.AccessExp*3600*24) * time.Second
	refreshExp := time.Duration(config.Cfg.Jc.RefreshExp*3600*24) * time.Second
	jwtToken := jwts.CreateToken(memIdStr, config.Cfg.Jc.AccessSecret, config.Cfg.Jc.RefreshSecret, accessExp, refreshExp)
	tokenList := &login.TokenMessage{
		AccessToken:    jwtToken.AccessToken,
		RefreshToken:   jwtToken.RefreshToken,
		AccessTokenExp: int64(jwtToken.AccessExp),
		TokenType:      "bearer",
	}
	// todo 存入缓存中
	go func() {
		memJson, _ := json.Marshal(mem)
		l.cache.Put(c, model.Member+"::"+memIdStr, string(memJson), accessExp)
		orgsJson, _ := json.Marshal(orgs)
		l.cache.Put(c, model.MemberOrganization+"::"+memIdStr, string(orgsJson), accessExp)
	}()
	return &login.LoginResponse{
		Member:           membMessage,
		OrganizationList: orgsMessage,
		TokenList:        tokenList,
	}, nil
}
func (l *LoginService) TokenVerify(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	token := msg.Token
	if strings.Contains(token, "bearer") {
		token = strings.Replace(token, "bearer ", "", 1)
	}
	parseToken, err := jwts.ParseToken(token, config.Cfg.Jc.AccessSecret)
	if err != nil {
		zap.L().Error("login token verify error！", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	// todo 放于redis中
	memJson, err := l.cache.Get(context.Background(), model.Member+"::"+parseToken)
	if err != nil {
		zap.L().Error("TokenVerify cache error！", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	if memJson == "" {
		zap.L().Error("TokenVerify cache token expired error！", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	membyId := &member.Member{}
	_ = json.Unmarshal([]byte(memJson), membyId)

	memMessage := &login.MemberMessage{}
	_ = copier.Copy(memMessage, membyId)
	memMessage.Code, _ = encrypts.Encrypt(strconv.FormatInt(membyId.Id, 10), model.AESKey)

	orgJson, err := l.cache.Get(context.Background(), model.MemberOrganization+"::"+parseToken)
	if err != nil {
		zap.L().Error("TokenVerify cache error！", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	if orgJson == "" {
		zap.L().Error("TokenVerify cache token expired error！", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	var orgs []*organization.Organization
	_ = json.Unmarshal([]byte(orgJson), &orgs)

	if len(orgs) > 0 {
		memMessage.OrganizationCode, _ = encrypts.Encrypt(strconv.FormatInt(orgs[0].Id, 10), model.AESKey)
	}
	return &login.LoginResponse{Member: memMessage}, nil
}
func (l *LoginService) MyOrgList(ctx context.Context, msg *login.UserMessage) (*login.OrgListResponse, error) {
	memId := msg.MemId
	orgs, err := l.organizationRepo.FindOrganizationByMemId(ctx, memId)
	if err != nil {
		zap.L().Error("MyOrgList FindOrganizationByMemId err", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	for _, org := range orgsMessage {
		org.Code, _ = encrypts.Encrypt(strconv.FormatInt(org.Id, 10), model.AESKey)
	}
	return &login.OrgListResponse{OrganizationList: orgsMessage}, nil
}
func (l *LoginService) FindMemberInfoById(ctx context.Context, msg *login.UserMessage) (*login.MemberMessage, error) {
	memberById, err := l.memberRepo.FindMemberById(context.Background(), msg.MemId)
	if err != nil {
		zap.L().Error("TokenVerify db FindMemberById error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	memMsg := &login.MemberMessage{}
	_ = copier.Copy(memMsg, memberById)
	memMsg.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey)
	orgs, err := l.organizationRepo.FindOrganizationByMemId(context.Background(), memberById.Id)
	if err != nil {
		zap.L().Error("TokenVerify db FindMember error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if len(orgs) > 0 {
		memMsg.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	memMsg.CreateTime = times.FormatByMill(memberById.CreateTime)
	return memMsg, nil
}
func (l *LoginService) FindMemberByIds(ctx context.Context, msg *login.UserMessage) (*login.MemberListResponse, error) {
	fmt.Printf("\n userMessage--memberIds:【%s】\n", msg.MemberIds)
	memberByIds, err := l.memberRepo.FindMemberByIds(context.Background(), msg.MemberIds)
	if err != nil {
		zap.L().Error("TokenVerify db FindMemberByIds error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if memberByIds == nil || len(memberByIds) <= 0 {
		return &login.MemberListResponse{MemberList: nil}, nil
	}
	mMap := make(map[int64]*member.Member)
	for _, v := range memberByIds {
		mMap[v.Id] = v
	}
	var memMsgs []*login.MemberMessage
	_ = copier.Copy(&memMsgs, memberByIds)
	for _, v := range memMsgs {
		m := mMap[v.Id]
		v.CreateTime = times.FormatByMill(m.CreateTime)
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
	}
	return &login.MemberListResponse{MemberList: memMsgs}, nil
}
