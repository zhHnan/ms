package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"hnz.com/ms_serve/ms-api/api/midd"
	srv "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-user/config"
	"hnz.com/ms_serve/ms-user/router"
	"hnz.com/ms_serve/ms-user/tracing"
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
	grpc := router.RegisterGrpc()
	router.RegisterEtcdServer()
	stop := func() {
		grpc.Stop()
	}
	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name, stop)
}
