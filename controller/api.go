/******
** @创建时间 : 2020/8/20 20:21
** @作者 : SongZhiBin
******/
package controller

import (
	"Happy/dao/redis"
	"Happy/model"
	"Happy/pkg/email"
	"Happy/pkg/randomCode"
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"go.uber.org/zap"
)

// 存放一些公用api

// VerificationCode验证码相关
func VerificationCode(c *gin.Context) {
	// var r *http.Request
	// 获取ip地址
	ip := exnet.ClientPublicIP(c.Request)
	if ip == "" {
		ip = exnet.ClientIP(c.Request)
	}

	// 判断redis中有没有缓存过验证码
	if redis.GetCurrentLimit(ip) {
		// 如果有缓存的话证明已经发送过邮件了 需要等
		zap.L().Info("Request Bus:", zap.String("IP", ip))
		model.ResponseError(c, model.CodeFrequentRequests)
		return
	}
	// 证明有效请求
	// 获取验证码
	code := randomCode.GetCode(randomCode.CodeModeMixing)
	// 发送邮件
	go email.GE.Send(email.GE.CreateTemp("718428482@qq.com", "Seer", code))
	// 写入验证码
	redis.SetCurrentLimit(ip, code)
	model.ResponseSuccess(c, "ok")
	return
}
