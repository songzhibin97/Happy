/******
** @创建时间 : 2020/8/15 22:36
** @作者 : SongZhiBin
******/
package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pbUser "Happy/model/pmodel/user"
	"Happy/settings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

// 用户相关
var GrpcConnNoAuth *grpc.ClientConn
var GrpcConnAuth *grpc.ClientConn

// 回收资源
func CloseConn() {
	_ = GrpcConnNoAuth.Close()
	_ = GrpcConnAuth.Close()
}

// Init:初始化grpc连接
func Init() {
	var err error
	GrpcConnNoAuth, err = grpc.Dial(":"+strconv.Itoa(settings.GetInt("GRPC.NoAuthPort")), grpc.WithInsecure())
	GrpcConnAuth, err = grpc.Dial(":"+strconv.Itoa(settings.GetInt("GRPC.AuthPort")), grpc.WithInsecure())
	if err != nil {
		zap.L().Error("grpc.Dial Error", zap.Error(err))
		//model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
}

// SignUpHandler:注册
// todo: 未为对邮箱等做限制校验等...
func SignUpHandler(c *gin.Context) {
	// 1.获取请求参数
	u := new(model.RegisterForm)
	// 2.校验有效性(使用validator来进行校验)
	err := ParameterVerification(c, u)
	if err != nil {
		return
	}
	cc := pbUser.NewUserClient(GrpcConnNoAuth)
	r, err := cc.Register(c, &pbUser.RegisterRequest{
		UserName:         u.UserName,
		Password:         u.Password,
		ConfirmPassword:  u.ConfirmPassword,
		UserInfo:         "",
		Email:            u.Email,
		VerificationCode: u.VerificationCode,
	})
	if err != nil {
		zap.L().Error("pbUser.NewUserClient", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse(r))
	return

}

// LoginHandler:登录
func LoginHandler(c *gin.Context) {
	u := new(model.LoginGet)
	// 2.校验有效性(使用validator来进行校验)
	err := ParameterVerification(c, u)
	if err != nil {
		return
	}
	// 改用grpc内部调用
	// grpc.WithInsecure() 安全参数 可传可不传
	cc := pbUser.NewUserClient(GrpcConnNoAuth)
	r, err := cc.Login(c, &pbUser.LoginRequest{
		UserName: u.UserName,
		Password: u.Password,
	})
	if err != nil {
		zap.L().Error("pbUser.NewUserClient", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	zap.L().Info("Login Response", zap.Any("Response", r))
	c.JSON(http.StatusOK, gmodel.GinResponse(r))
	return
}
