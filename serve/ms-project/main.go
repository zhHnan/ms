package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"hnz.com/ms_serve/ms-api/api/midd"
	srv "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-project/config"
	"hnz.com/ms_serve/ms-project/router"
	"hnz.com/ms_serve/ms-project/tracing"
	"log"
)

func main() {
	r := gin.Default()
	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	r.Use(midd.RequestLog())
	r.Use(otelgin.Middleware("ms-api"))
	//从配置中读取日志配置，初始化日志
	router.InitRouter(r)
	// 初始化rpc
	router.InitUserRpc()
	grpc := router.RegisterGrpc()
	router.RegisterEtcdServer()
	ka := config.InitKafkaWriter()
	stop := func() {
		grpc.Stop()
		ka()
	}
	// 开启pprof
	pprof.Register(r)
	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name, stop)
}
