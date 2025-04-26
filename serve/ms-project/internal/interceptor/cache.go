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

func (i *Interceptor) CacheInterceptor() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		respType := i.cacheMap[info.FullMethod]
		if respType == nil {
			return handler(ctx, req)
		}

		// 记录请求开始时间
		start := time.Now()

		// 将请求参数转换为 JSON 字符串
		reqJSON, err := json.Marshal(req)
		if err != nil {
			zap.L().Error("Failed to marshal request", zap.Error(err))
			reqJSON = []byte("{}")
		}

		// 打印请求信息
		zap.L().Info("gRPC request",
			zap.String("method", info.FullMethod),
			zap.String("request", string(reqJSON)),
		)

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

		// 将响应结果转换为 JSON 字符串
		respJSON, err := json.Marshal(resp)
		if err != nil {
			zap.L().Error("Failed to marshal response", zap.Error(err))
			respJSON = []byte("{}")
		}

		// 打印响应信息和处理时间
		duration := time.Since(start)
		zap.L().Info("gRPC response",
			zap.String("method", info.FullMethod),
			zap.String("response", string(respJSON)),
			zap.Error(err),
			zap.Duration("duration", duration),
		)

		bytes, _ := json.Marshal(resp)
		_ = i.cache.Put(c, info.FullMethod+"::"+cacheKey, string(bytes), model.CacheExpire*time.Minute)
		zap.L().Info(info.FullMethod + "\n writer cache \n")

		return resp, err
	}
}
