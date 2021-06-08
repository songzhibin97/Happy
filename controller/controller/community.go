package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pbCommunity "Happy/model/pmodel/community"
	pb "Happy/model/pmodel/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommunityList:获取社区列表
// @Summary 获取社区列表
// @Description 用于获取获取社区列表接口 内部调用grpc接口
// @Tags 社区相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseStruct
// @Router /communityList [get]
func CommunityList(c *gin.Context) {
	cc := pbCommunity.NewCommunityClient(GrpcConnAuth)
	// 获取请求头的token
	ctx, err := GetToken(c)
	if err != nil {
		return
	}
	r, err := cc.CommunityList(ctx, &pbCommunity.CommunityListRequest{})
	if err != nil {
		zap.L().Error("pbCommunity.CommunityListRequest", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return
}

// CommunityDetail:获取社区详情
// @Summary 获取社区详情
// @Description 用于获取社区详情接口 内部调用grpc接口
// @Tags 社区相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string ture "Bearer 用户令牌"
// @Param object query model.CommunityDetailRequest ture "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseStruct
// @Router /communityDetail [get]
func CommunityDetail(c *gin.Context) {
	// 参数校验
	rq := new(model.CommunityDetailRequest)
	// 2.校验有效性(使用validator来进行校验)
	err := ParameterVerification(c, rq)
	if err != nil {
		return
	}
	cc := pbCommunity.NewCommunityClient(GrpcConnAuth)
	//// 获取请求头的token
	//ctx, err := GetToken(c)
	//if err != nil {
	//	return
	//}
	r, err := cc.CommunityDetail(c, &pbCommunity.CommunityDetailRequest{ID: int64(rq.ID)})
	if err != nil {
		zap.L().Error("pbCommunity.CommunityListRequest", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return
}