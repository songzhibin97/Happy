/******
** @创建时间 : 2020/8/22 09:35
** @作者 : SongZhiBin
******/
package gcontroller

import (
	"Happy/dao/redis"
	sqls "Happy/dao/sql"
	"Happy/middleware"
	"Happy/model/gmodel"
	"Happy/model/model"
	pbUser "Happy/model/pmodel/user"
	"Happy/pkg/email"
	"Happy/pkg/jwt"
	"Happy/pkg/randomCode"
	"Happy/settings"
	"context"
	"database/sql"
	"go.uber.org/zap"
)

// grpcController

// UserServer:要实现对应对应pd的接口
type UserServer struct{}

// 用户相关
// SignUpHandler:注册
// todo: 为对邮箱等做限制校验等...

// Register:注册
func (s *UserServer) Register(ctx context.Context, request *pbUser.RegisterRequest) (*pbUser.Response, error) {
	// 校验请求参数
	res, err := _verification(request)
	if err != nil {
		return res, nil
	}
	// 3.数据库判断是否有该用户存在
	err = sqls.IsExist(request.UserName)
	if err != nil {
		zap.L().Info("IsExist Error", zap.Error(err))
		if err == model.CodeUserExist.Err() {
			return gmodel.ResponseError(model.CodeUserExist), nil
		}
		// 否则是查询错误
		return gmodel.ResponseWithMsg(model.CodeServerBusy, err.Error()), nil
	}
	// 这里使用邮箱地址作为redis的ip key

	// 4.增加校验 验证码邮箱
	code, ok := redis.GetVerificationCode(request.Email)
	// 如果已经存在
	if !ok || code != request.VerificationCode {
		return gmodel.ResponseError(model.CodeInvalidVerificationCode), nil
	}
	// 5.校验完成 插入数据库
	ok = sqls.InsertUser(request.UserName, request.Password)
	if !ok {
		zap.L().Info("用户创建失败")
		return gmodel.ResponseWithMsg(model.CodeServerBusy, "用户创建失败"), nil
	}
	return gmodel.ResponseSuccess(map[string]string{"success": "ok"}), nil
}

// Login:登录
func (s *UserServer) Login(ctx context.Context, request *pbUser.LoginRequest) (*pbUser.Response, error) {
	// 校验请求参数
	res, err := _verification(request)
	if err != nil {
		return res, nil
	}
	// 数据库判断该用户是否存在
	// 3.数据库判断是否有该用户存在
	user, err := sqls.IsUserValid(request.UserName, sqls.GetEncrypt(request.Password))
	if err != nil {
		zap.L().Info("IsUserValid Error", zap.Error(err))
		// 如果是用户已经存在返回对应的错误信息
		if err == model.CodeInvalidPassword.Err() || err == sql.ErrNoRows {
			return gmodel.ResponseError(model.CodeInvalidPassword), nil
		}
		// 否则是查询错误
		return gmodel.ResponseError(model.CodeServerBusy), nil
	}
	// 校验成功 返回jwt
	// 判断当前jwt的状态
	if settings.GetString("JWT.Mode") == "refresh" {
		aToken, rToken, err := jwt.GetACRFToken(user.UID)
		if err != nil {
			zap.L().Error("GetACRFToken Error", zap.Error(err))
			return gmodel.ResponseWithMsg(model.CodeServerBusy, err.Error()), nil

		}
		//gin.H{
		//	middleware.AccessToken:  aToken,
		//	middleware.RefreshToken: rToken,
		//}
		redis.SetToken(int64(user.UID), aToken)
		return gmodel.ResponseSuccess(
			map[string]string{middleware.AccessToken: aToken, middleware.RefreshToken: rToken},
		), nil
	} else {
		aToken, err := jwt.GetJWT(user.UID)
		if err != nil {
			zap.L().Error("GetJWT Error", zap.Error(err))
			return gmodel.ResponseWithMsg(model.CodeServerBusy, err.Error()), nil
		}
		redis.SetToken(int64(user.UID), aToken)
		//gin.H{
		//	middleware.AccessToken: aToken,
		//}
		return gmodel.ResponseSuccess(map[string]string{middleware.AccessToken: aToken}), nil
	}
}

// Verification:获取验证码
func (s *UserServer) Verification(ctx context.Context, request *pbUser.VerificationRequest) (*pbUser.Response, error) {
	// 校验请求参数
	res, err := _verification(request)
	if err != nil {
		return res, nil
	}
	// 判断redis中有没有缓存过验证码
	if redis.GetCurrentLimit(request.Email) {
		// 如果有缓存的话证明已经发送过邮件了 需要等
		zap.L().Info("Request Bus:", zap.String("email", request.Email))
		return gmodel.ResponseError(model.CodeFrequentRequests), nil
	}
	// 证明有效请求
	// 获取验证码
	code := randomCode.GetCode(randomCode.CodeModeMixing)
	// 发送邮件
	go email.GE.Send(email.GE.CreateTemp(request.Email, "Happy", code))
	// 写入验证码
	redis.SetCurrentLimit(request.Email, code)
	return gmodel.ResponseSuccess(map[string]string{"success": "ok"}), nil
}
