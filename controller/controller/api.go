package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pbUser "Happy/model/pmodel/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 存放一些公用api

// VerificationCode验证码相关
// @Summary 发送验证码
// @Description 用于发送验证码的接口 内部调用grpc接口
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param object query model.Email false "发送参数"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseStruct
// @Router /VerificationCode [post]
func VerificationCode(c *gin.Context) {
	e := new(model.Email)
	// 2.校验有效性(使用validator来进行校验)
	err := ParameterVerification(c, e)
	if err != nil {
		return
	}
	cc := pbUser.NewUserClient(GrpcConnNoAuth)
	r, err := cc.Verification(c, &pbUser.VerificationRequest{
		Email: e.Addr,
	})
	if err != nil {
		zap.L().Error("pbUser.NewUserClient", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse(r))
	return
}
