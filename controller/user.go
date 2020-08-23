/******
** @创建时间 : 2020/8/15 22:36
** @作者 : SongZhiBin
******/
package controller

import (
	"Happy/dao/redis"
	"Happy/dao/sql"
	"Happy/middleware"
	"Happy/model"
	ip2 "Happy/pkg/ip"
	"Happy/pkg/jwt"
	"Happy/settings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户相关

// SignUpHandler:注册
// todo: 为对邮箱等做限制校验等...
func SignUpHandler(c *gin.Context) {
	// 1.获取请求参数
	u := new(model.RegisterForm)
	// 2.校验有效性(使用validator来进行校验)
	if err := c.ShouldBind(u); err != nil {
		// 校验失败
		// 判断error是否是校验失败的error
		errs, ok := isVerifyError(err)
		if !ok {
			// 如果不是校验失败的错误就返回异常 标记服务器异常
			zap.L().Error("isVerifyError", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err)
			return
		}
		// 是参数校验的错误返回对应错误
		zap.L().Info("VerifyError", zap.Any("error", errs))
		model.ResponseErrorWithMsg(c, model.CodeInvalidParams, errs)
		return
	}
	// 3.数据库判断是否有该用户存在
	err := sql.IsExist(u.UserName)
	if err != nil {
		zap.L().Info("IsExist Error", zap.Error(err))
		// 如果是用户已经存在返回对应的错误信息
		if err == model.CodeUserExist.Err() {
			model.ResponseErrorWithMsg(c, model.CodeUserExist, err.Error())
			return
		}
		// 否则是查询错误
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	// 获取ip
	ip := ip2.GetIP(c)
	// 4.增加校验 验证码邮箱
	code, ok := redis.GetVerificationCode(ip)
	if !ok || code != u.VerificationCode {
		model.ResponseError(c, model.CodeInvalidVerificationCode)
		return
	}
	// 5.校验完成 插入数据库
	ok = sql.InsertUser(u)
	if !ok {
		zap.L().Info("用户创建失败")
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, "用户创建失败")
		return
	}
	model.ResponseSuccess(c, "ok")
	return
}

// LoginHandler:登录
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数
	u := new(model.LoginGet)
	// 2.校验有效性(使用validator来进行校验)
	if err := c.ShouldBind(u); err != nil {
		// 校验失败
		// 判断error是否是校验失败的error
		errs, ok := isVerifyError(err)
		if !ok {
			// 如果不是校验失败的错误就返回异常 标记服务器异常
			zap.L().Error("isVerifyError", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
			return
		}
		// 是参数校验的错误返回对应错误
		zap.L().Info("VerifyError", zap.Any("error", errs))
		model.ResponseErrorWithMsg(c, model.CodeInvalidParams, errs)
		return
	}
	// 3.数据库判断是否有该用户存在
	user, err := sql.IsUserValid(u.UserName, sql.GetEncrypt(u.Password))
	if err != nil {
		zap.L().Info("IsUserValid Error", zap.Error(err))
		// 如果是用户已经存在返回对应的错误信息
		if err == model.CodeInvalidPassword.Err() {
			model.ResponseError(c, model.CodeInvalidPassword)
			return
		}
		// 否则是查询错误
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	// 校验成功 返回jwt
	// 判断当前jwt的状态
	if settings.GetString("JWT.Mode") == "refresh" {
		aToken, rToken, err := jwt.GetACRFToken(user.UID)
		if err != nil {
			zap.L().Error("GetACRFToken Error", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
			return
		}
		model.ResponseSuccess(c, gin.H{
			middleware.AccessToken:  aToken,
			middleware.RefreshToken: rToken,
		})
		redis.SetToken(int64(user.UID), aToken)
		return
	} else {
		aToken, err := jwt.GetJWT(user.UID)
		if err != nil {
			zap.L().Error("GetJWT Error", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
			return
		}
		model.ResponseSuccess(c, gin.H{
			middleware.AccessToken: aToken,
		})
		redis.SetToken(int64(user.UID), aToken)
		return
	}
}
