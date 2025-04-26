package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	_ "hnz.com/ms_serve/ms-api/api"
	"hnz.com/ms_serve/ms-api/api/midd"
	"hnz.com/ms_serve/ms-api/config"
	"hnz.com/ms_serve/ms-api/router"
	"hnz.com/ms_serve/ms-api/tracing"
	srv "hnz.com/ms_serve/ms-common"
	"log"
	"net/http"
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
	r.StaticFS("/upload", http.Dir("./upload"))
	//从配置中读取日志配置，初始化日志
	config.Cfg.InitZapLog()
	router.InitRouter(r)

	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name, nil)
}
