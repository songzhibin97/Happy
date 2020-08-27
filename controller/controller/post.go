/******
** @创建时间 : 2020/8/27 18:59
** @作者 : SongZhiBin
******/
package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pbPost "Happy/model/pmodel/post"
	pb "Happy/model/pmodel/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 帖子相关
const (
	Mode string = "mode"
	ID   string = "id"
	Page string = "page"
	Max  string = "max"
)

// PostList:获取帖子列表
func PostList(c *gin.Context) {
	//cc := pbCommunity.NewCommunityClient(GrpcConnAuth)
	mode := c.DefaultQuery(Mode, "0") // mode获取帖子的模式 0为社区 1为个人
	iMode, err := StoA(mode, c)
	if err != nil {
		return
	}
	id := c.DefaultQuery(ID, "0") // id 社区id或用户id
	iId, err := StoA(id, c)
	if err != nil {
		return
	}
	page := c.DefaultQuery(Page, "1")
	iPage, err := StoA(page, c)
	if err != nil {
		return
	}
	max := c.DefaultQuery(Max, "10")
	iMax, err := StoA(max, c)
	if err != nil {
		return
	}
	if iPage <= 0 {
		iPage = 1 // page不能小于0
	}
	if iMax <= 0 {
		iMax = 1 // max不能小于0
	}
	cc := pbPost.NewPostClient(GrpcConnAuth)
	// 获取请求头的token
	ctx, err := GetToken(c)
	if err != nil {
		return
	}
	switch iMode {
	case 0:
		r, err := cc.PostList(ctx, &pbPost.GetPostListRequest{
			Model: (pbPost.GetPostListRequest_State)(iMode),
			ID:    &pbPost.GetPostListRequest_CommunityID{CommunityID: int64(iId)},
			Page:  int64(iPage),
			Max:   int64(iMax),
		})
		if err != nil {
			zap.L().Error("PostList", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
			return
		}
		model.ResponseSuccess(c, gmodel.GinResponse((*pb.Response)(r)))
	case 1:
		r, err := cc.PostList(ctx, &pbPost.GetPostListRequest{
			Model: (pbPost.GetPostListRequest_State)(iMode),
			ID:    &pbPost.GetPostListRequest_AuthorID{AuthorID: int64(iId)},
			Page:  int64(iPage),
			Max:   int64(iMax),
		})
		if err != nil {
			zap.L().Error("PostList", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
			return
		}
		model.ResponseSuccess(c, gmodel.GinResponse((*pb.Response)(r)))
	}
}

// CreatePost:新建帖子
func CreatePost(c *gin.Context) {
	// 参数校验
	rq := new(model.CreatePost)
	// 2.校验有效性(使用validator来进行校验)
	err := ParameterVerification(c, rq)
	if err != nil {
		return
	}
	cc := pbPost.NewPostClient(GrpcConnAuth)
	// 获取请求头的token
	ctx, err := GetToken(c)
	if err != nil {
		return
	}
	r, err := cc.CreatePost(ctx, &pbPost.CreatePostRequest{
		CommunityID: rq.CommunityID,
		Title:       rq.Title,
		Content:     rq.Content,
	})
	if err != nil {
		zap.L().Error("cc.CreatePost Error", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return

}

// GetPostDetail:获取帖子详情
func GetPostDetail(c *gin.Context) {
	// 获取帖子id
	id := c.Query(ID)
	if id == "" {
		model.ResponseError(c, model.CodeInvalidParams)
		return
	}
	iid, err := StoA(id, c)
	if err != nil {
		return
	}
	cc := pbPost.NewPostClient(GrpcConnAuth)
	// 获取请求头的token
	ctx, err := GetToken(c)
	if err != nil {
		return
	}
	r, err := cc.GetPostDetail(ctx, &pbPost.GetPostDetailRequest{
		PostID: int64(iid),
	})
	if err != nil {
		zap.L().Error("GetPostDetail", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return
}
