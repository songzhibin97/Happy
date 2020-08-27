/******
** @创建时间 : 2020/8/26 20:21
** @作者 : SongZhiBin
******/
package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pbCommunity "Happy/model/pmodel/community"
	pb "Happy/model/pmodel/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// CommunityList:获取社区列表
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
