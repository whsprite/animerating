package main

import (
	"Goweb_cli/dao/mysql"
	"Goweb_cli/logger"
	"Goweb_cli/pkg/snowflake"
	"Goweb_cli/routes"
	"Goweb_cli/settings"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	//1. 加载配置
	if err := settings.Init(); err != nil {
		zap.L().Error("settings.Init() failed", zap.Error(err))
	}
	//2.配置日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		zap.L().Error("logger.Init() failed", zap.Error(err))
	}
	defer zap.L().Sync()
	//3.配置mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		zap.L().Error("mysql.Init() failed", zap.Error(err))
	}
	if err := snowflake.Init(settings.Conf.StartTime, 1); err != nil {
		zap.L().Error("settings.Init() failed", zap.Error(err))
	}
	defer mysql.Close()
	//4.注册路由
	r := routes.SetUp()

	//5.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen:", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
