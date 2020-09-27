/******
** @创建时间 : 2020/9/14 17:27
** @作者 : SongZhiBin
******/
package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pb "Happy/model/pmodel/user"
	pbVote "Happy/model/pmodel/vote"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// Vote:投票
// @Summary 用户投票接口
// @Description 用于用户投票接口 内部调用grpc接口
// @Tags 投票相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body model.ParamVoteDate true "获取帖子的模式"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseStruct
// @Router /Vote [post]
func Vote(c *gin.Context) {
	// 1.获取请求参数
	v := new(model.ParamVoteDate)
	// 2.校验有效性(使用validator来进行校验)
	err := ParameterVerification(c, v)
	if err != nil {
		return
	}
	// 校验完成 生成客户端调用grpc
	cc := pbVote.NewVoteClient(GrpcConnAuth)
	r, err := cc.Vote(c, &pbVote.VoteRequest{
		PostID: v.PostID,
		Mode:   (pbVote.VoteRequest_State)(v.Direction),
	})
	if err != nil {
		zap.L().Error("pbVote.Vote", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return

}
