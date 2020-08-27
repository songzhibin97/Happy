/******
** @创建时间 : 2020/8/17 16:45
** @作者 : SongZhiBin
******/
package middleware

import (
	"Happy/controller/controller"
	"Happy/model/gmodel"
	"Happy/model/model"
	pbJwt "Happy/model/pmodel/jwt"
	pbUser "Happy/model/pmodel/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 关于jwt的中间件

const (
	ConTextUserID = "UID"
	AccessToken   = "accessJWT"
	RefreshToken  = "refreshJWT"
)

// VerificationJWT:判断JWT是否有效
func VerificationJWT(c *gin.Context) {
	// 从请求头获取jwt
	access := c.Request.Header.Get(AccessToken)
	if access == "" {
		// 未携带参数
		model.ResponseError(c, model.CodeInvalidToken)
		c.Abort()
		return
	}
	cc := pbJwt.NewJWTClient(controller.GrpcConnNoAuth)
	r, err := cc.VerificationJWT(c, &pbJwt.VerificationJWTRequest{
		Access: access,
	})
	if err != nil {
		zap.L().Error("pbUser.NewUserClient", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	switch r.State {
	case pbJwt.VerificationJWTResponse_NotPass:
		// 未通过校验
		model.ResponseError(c, model.CodeInvalidToken)
		c.Abort()
		return
	case pbJwt.VerificationJWTResponse_ExpiredJump:
		// 未通过校验带跳转
		model.ResponseError(c, model.CodeAccessExpired)
		c.Abort()
		return
	case pbJwt.VerificationJWTResponse_Expired:
		// 未通过校验不带跳转
		model.ResponseError(c, model.CodeJWTExpired)
		c.Abort()
		return
	case pbJwt.VerificationJWTResponse_MultiTerminalLogin:
		model.ResponseError(c, model.CodeMultiTerminalLogin)
		c.Abort()
		return
	case pbJwt.VerificationJWTResponse_Pass:
		// 上下文存入
		c.Set(AccessToken, access)
		c.Set(ConTextUserID, r.Uid)
		c.Next()
	}
}

// VerificationRefreshJWT:用于校验刷新token
func VerificationRefreshJWT(c *gin.Context) {
	access := c.Request.Header.Get(AccessToken)
	refresh := c.Request.Header.Get(RefreshToken)

	cc := pbJwt.NewJWTClient(controller.GrpcConnNoAuth)
	r, err := cc.VerificationRefreshJWT(c, &pbJwt.VerificationRefreshJWTRequest{
		Access:  access,
		Refresh: refresh,
	})
	if err != nil {
		zap.L().Error("pbUser.NewUserClient", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pbUser.Response)(r)))
	return
}
