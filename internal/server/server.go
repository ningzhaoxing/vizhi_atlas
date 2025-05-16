package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vizhi_atlas/initialize"
	"vizhi_atlas/internal/pkg/globals"
	"vizhi_atlas/internal/pkg/logger"
)

func Run() {
	if err := initialize.InitApp(); err != nil {
		panic(err)
	}
	initRouter(globals.E)

	svr := &http.Server{
		Addr:    fmt.Sprint(globals.C.App.Host, ":", globals.C.App.Port),
		Handler: globals.E,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			fmt.Println("server-->", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		fmt.Printf("stutdown err %v \n", err)
	}
	fmt.Println("shutdown-->ok")

	// 记录应用关闭日志
	logger.Info("应用程序正常关闭")
}
