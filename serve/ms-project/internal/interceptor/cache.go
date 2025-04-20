package interceptor

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-grpc/project"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-user/pkg/model"
	"time"
)

type Interceptor struct {
	cache    repo.Cache
	cacheMap map[string]any
}

func NewInterceptor() *Interceptor {
	cacheMap := make(map[string]any)
	cacheMap["/project.service.v1.ProjectService/FindProjectByMemId"] = &project.MyProjectResponse{}
	return &Interceptor{cache: dao.Rc, cacheMap: cacheMap}
}

func (i *Interceptor) CacheInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		respType := i.cacheMap[info.FullMethod]
		if respType == nil {
			return handler(ctx, req)
		}
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		marshal, _ := json.Marshal(req)
		cacheKey := encrypts.Md5(string(marshal))
		respJson, _ := i.cache.Get(c, info.FullMethod+"::"+cacheKey)
		if respJson != "" {
			_ = json.Unmarshal([]byte(respJson), &respType)
			zap.L().Info(info.FullMethod + "\n read cache \n")
			return respType, nil
		}
		resp, err = handler(ctx, req)
		bytes, _ := json.Marshal(resp)
		_ = i.cache.Put(c, info.FullMethod+"::"+cacheKey, string(bytes), model.CacheExpire*time.Minute)
		zap.L().Info(info.FullMethod + "\n writer cache \n")
		return
	})
}
