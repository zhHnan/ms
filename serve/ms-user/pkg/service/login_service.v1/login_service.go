package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-user/internal/dao"
	"hnz.com/ms_serve/ms-user/internal/data/member"
	"hnz.com/ms_serve/ms-user/internal/data/organization"
	"hnz.com/ms_serve/ms-user/internal/repo"
	"hnz.com/ms_serve/ms-user/pkg/model"
	"log"
	"time"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache            repo.Cache
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
}

func New() *LoginService {
	return &LoginService{
		cache:            dao.Rc,
		memberRepo:       dao.NewMemberDao(),
		organizationRepo: dao.NewOrganizationDao(),
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

		//if l.cache == nil {
		//	log.Println("cache 未初始化")
		//	return
		//}

		err := l.cache.Put(c, model.RegisterKey+mobile, code, time.Minute*5)
		if err != nil {
			log.Println("验证码存入redis失败！", err)
		}
		log.Printf("手机号和验证码存入到redis中：REGISTER_%s:%s", mobile, code)
	}()
	return &login.CaptchaResponse{Code: code}, nil
}

func (ls *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	// 检验参数
	// 校验验证码
	redisCode, err := ls.cache.Get(ctx, model.RegisterKey+msg.Mobile)
	if err != nil {
		zap.L().Error("验证码校验失败！", zap.Error(err))
		return nil, errs.GrpcError(model.RedisError)
	}
	if redisCode != msg.Captcha {
		return nil, errs.GrpcError(model.CaptchaError)
	}
	// 校验业务逻辑（邮箱是否注册、账号是否注册、手机号是否注册）
	exist, err := ls.memberRepo.GetMemberByEmail(ctx, msg.Email)
	if err != nil {
		zap.L().Error("register database get error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if exist {
		return nil, errs.GrpcError(model.EmailExist)
	}
	exist, err = ls.memberRepo.GetMemberByAccount(ctx, msg.Name)
	if err != nil {
		zap.L().Error("register database get error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if exist {
		return nil, errs.GrpcError(model.NameExist)
	}
	exist, err = ls.memberRepo.GetMemberByMobile(ctx, msg.Mobile)
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
	err = ls.memberRepo.SaveMember(ctx, mem)
	if err != nil {
		zap.L().Error("register database save error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	// 存入组织
	org := &organization.Organization{
		Name:       mem.Name + "个人组织",
		MemberId:   mem.Id,
		CreateTime: time.Now().UnixMilli(),
		Personal:   int32(model.Personal),
		Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
	}
	err = ls.organizationRepo.SaveOrganization(ctx, org)
	if err != nil {
		zap.L().Error("register SaveOrganization db err", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	// 返回结果
	return &login.RegisterResponse{}, nil
}
