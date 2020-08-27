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
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"net/http"
)

// CommunityList:获取社区列表
func CommunityList(c *gin.Context) {
	cc := pbCommunity.NewCommunityClient(GrpcConnAuth)
	// 获取请求头的token
	token, ok := c.Get("accessJWT")
	if !ok {
		zap.L().Error("Get(middleware.AccessToken)")
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, "获取上下文Token失败")
		return
	}
	md := metadata.Pairs("authorization", fmt.Sprintf("bearer %s", token))
	ctx := metadata.NewOutgoingContext(c, md)
	r, err := cc.CommunityList(ctx, &pbCommunity.CommunityListRequest{})
	if err != nil {
		zap.L().Error("pbCommunity.CommunityListRequest", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return
}

// CommunityDetail:获取社区列表
func CommunityDetail(c *gin.Context) {
	// 参数校验
	rq := new(model.CommunityDetailRequest)
	// 2.校验有效性(使用validator来进行校验)
	ParameterVerification(c, rq)

	cc := pbCommunity.NewCommunityClient(GrpcConnAuth)
	r, err := cc.CommunityDetail(c, &pbCommunity.CommunityDetailRequest{ID: int64(rq.ID)})
	if err != nil {
		zap.L().Error("pbCommunity.CommunityListRequest", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return
}
