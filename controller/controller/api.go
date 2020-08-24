/******
** @创建时间 : 2020/8/20 20:21
** @作者 : SongZhiBin
******/
package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pbUser "Happy/model/pmodel/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 存放一些公用api

// VerificationCode验证码相关
func VerificationCode(c *gin.Context) {
	e := new(model.Email)
	cc := pbUser.NewUserClient(GrpcConn)
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
