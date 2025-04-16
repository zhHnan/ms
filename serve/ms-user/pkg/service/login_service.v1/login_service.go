package login_service_v1

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"hnz.com/ms_serve/common"
	"hnz.com/ms_serve/common/errs"
	"hnz.com/ms_serve/ms-user/pkg/dao"
	"hnz.com/ms_serve/ms-user/pkg/model"
	"hnz.com/ms_serve/ms-user/pkg/repo"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	fmt.Println("run LoginService New...")
	if dao.Rc == nil {
		fmt.Println("dao.Rc is nil, please initialize redis cache first")
	}
	return &LoginService{
		cache: dao.Rc,
	}
}

func (l *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
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

		err := l.cache.Put(c, "REGISTER_"+mobile, code, time.Minute*5)
		if err != nil {
			log.Println("验证码存入redis失败！", err)
		}
		log.Printf("手机号和验证码存入到redis中：REGISTER_%s:%s", mobile, code)
	}()
	return &CaptchaResponse{Code: code}, nil
}
