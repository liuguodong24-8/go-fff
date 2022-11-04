package main

import (
	"blockExplore/internal/config"
	"blockExplore/internal/router"
	"blockExplore/internal/service"
	"blockExplore/internal/util"
	"blockExplore/pkgs/logger"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	var conf string
	usage()
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	flag.Parse()
	if conf == "" {
		panic("请指定配置文件位置")
	}
	config.Load(conf)

}

func main() {
	//注册日志
	if err := util.RegisterLogger(logger.Config{
		Channel:    config.Setting.Log.Channel,
		Level:      logger.Level(config.Setting.Log.Level),
		OutputFile: config.Setting.Log.Output,
		WithStack:  config.Setting.Log.Stack,
	}); nil != err {
		panic(fmt.Sprintf("注册日志失败:%s", err.Error()))
	}
	// 捕获异常信息
	defer util.CatchException()

	//注册路由
	router.Init()
	srv := &http.Server{Addr: fmt.Sprintf(":%d", config.Setting.App.Port), Handler: router.Router}
	util.Logger.Info("启动服务")
	go func() {
		if err := srv.ListenAndServe(); nil != err && http.ErrServerClosed != err {
			panic("启动服务失败")
		}
	}()
	err := service.SetMongoDb()
	if err != nil {
		util.Logger.Error("链接mongodb失败")
	}

	err = service.SetEthClient()
	if err != nil {
		util.Logger.Error("链接以太坊失败")
	}

	// 接收终端信号来关闭服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	util.Logger.Info("关闭服务")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); nil != err {
		util.Logger.Error("关闭服务异常，强制关闭")
	}
	util.Logger.Info("关闭服务完成")

}

// usage 返回使用方法
func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: go run cmd/server/main.go -c {配置文件}
`)
}
