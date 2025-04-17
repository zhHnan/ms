package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-user/internal/dao"
	"hnz.com/ms_serve/ms-user/internal/repo"
	"hnz.com/ms_serve/ms-user/pkg/model"
	"log"
	"time"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache      repo.Cache
	memberRepo repo.MemberRepo
}

func New() *LoginService {
	return &LoginService{
		cache:      dao.Rc,
		memberRepo: dao.NewMemberDao(),
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
	// 返回结果
	return &login.RegisterResponse{}, nil
}
