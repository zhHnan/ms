package login_service_v1

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"hnz.com/ms_serve/user/pkg/dao"
	"hnz.com/ms_serve/user/pkg/repo"
	"log"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func NewLoginService() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (l *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	mobile := msg.Mobile
	if mobile == "" {
		return nil, errors.New("手机号不合法！")
	}
	// 生成验证码
	code := "123456"
	go func() {
		// 发送验证码
		time.Sleep(time.Second * 2)
		zap.L().Info("验证码发送成功！")
		c, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		err := l.cache.Put(c, "REGISTER_"+mobile, code, time.Minute*5)
		if err != nil {
			log.Println("验证码存入redis失败！", err)
		}
		log.Printf("手机号和验证码存入到redis中：REGISTER_%s:%s", mobile, code)
	}()
	return &CaptchaResponse{}, nil
}
