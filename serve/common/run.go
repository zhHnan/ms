package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, addr, serverName string) {
	// 启动http服务
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	// 启动http服务
	go func() {
		log.Printf("%s server running in %s \n", serverName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	// SIGINT 用户发出INTR信号(ctrl + c)触发
	// SIGTERM 结束程序(可被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("shutting down project %s server... \n", serverName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s server shutdown, caused by %v:\n", serverName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("waiting timeout...")
	}
	log.Printf("%s server stop success... \n", serverName)

}
