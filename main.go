/******
** @创建时间 : 2020/8/11 17:23
** @作者 : SongZhiBin
******/
package main

import (
	"Happy/dao/redis"
	"Happy/dao/sql"
	"Happy/logger"
	"Happy/router"
	"Happy/settings"
	"github.com/fvbock/endless"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Go Web 通用模板
/*
项目架构:
	- controller // 调度模块 用于存放路由调度函数
	- dao // 数据库相关
		- redis
			- redis.go
		- sql
			- sql.go
	- log // 存放日志项
	- logger // 初始化zap
		- logger.go
	- logic // 存放业务逻辑
	- pkg // 存放第三方包
	- model // 存放模型
	- settings // 用于初始化viper
		- settings.go
	- router // 存放路由信息
		- router.go
	- config.ini // 配置文件
	- go.mod
	- main.go // 主函数
*/
func main() {
	// 1.加载配置
	err := settings.Init()
	if err != nil {
		zap.L().Panic("Settings Init Error", zap.Error(err))
		return
	}
	// 2.初始化日志模块
	err = logger.Init()
	if err != nil {
		zap.L().Panic("Logger Init Error", zap.Error(err))
		return
	}
	defer logger.Close()

	// 3.初始化sql模块
	err = sql.Init()
	if err != nil {
		zap.L().Panic("SQL Init Error", zap.Error(err))
		return
	}
	defer sql.Close()

	// 4.初始化Redis模块
	err = redis.Init()
	if err != nil {
		zap.L().Panic("Redis Init Error", zap.Error(err))
		return
	}
	defer redis.Close()
	// 5.注册路由
	r := router.SetUp()
	// 6.启动服务(注册优雅关机/重启)
	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	if err := endless.ListenAndServe(":"+viper.GetString("APP.Port"), r); err != nil {
		zap.L().Error("listen error", zap.Error(err))
	}
	zap.L().Debug("Server exiting")
}
