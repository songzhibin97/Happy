package controller

import (
	"Happy/model/gmodel"
	"Happy/model/model"
	pbPost "Happy/model/pmodel/post"
	pb "Happy/model/pmodel/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 帖子相关
const (
	Mode string = "mode"
	ID   string = "id"
	Page string = "page"
	Max  string = "max"
)

// PostList:获取帖子列表
// @Summary 获取帖子列表
// @Description 用于获取帖子列表接口 内部调用grpc接口
// @Tags 帖子相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query model.ParamPost true "获取帖子的模式"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseStruct
// @Router /PostList [get]
func PostList(c *gin.Context) {
	//cc := pbCommunity.NewCommunityClient(GrpcConnAuth)
	// 1.获取请求参数
	v := new(model.ParamPost)
	// 2.校验有效性(使用validator来进行校验)
	err := ParameterVerification(c, v)
	if err != nil {
		return
	}
	cc := pbPost.NewPostClient(GrpcConnAuth)
	// 获取请求头的token
	ctx, err := GetToken(c)
	if err != nil {
		return
	}
	switch v.Mode {
	case 0:
		r, err := cc.PostList(ctx, &pbPost.GetPostListRequest{
			Model: (pbPost.GetPostListRequest_State)(v.Mode),
			ID:    &pbPost.GetPostListRequest_CommunityID{CommunityID: int64(v.ID)},
			Page:  int64(v.Page),
			Max:   int64(v.Max),
		})
		if err != nil {
			zap.L().Error("PostList", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
			return
		}
		model.ResponseSuccess(c, gmodel.GinResponse((*pb.Response)(r)))
	case 1:
		r, err := cc.PostList(ctx, &pbPost.GetPostListRequest{
			Model: (pbPost.GetPostListRequest_State)(v.Mode),
			ID:    &pbPost.GetPostListRequest_AuthorID{AuthorID: int64(v.ID)},
			Page:  int64(v.Page),
			Max:   int64(v.Max),
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
// @Summary 新建帖子
// @Description 用于新建帖子接口 内部调用grpc接口
// @Tags 帖子相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body model.CreatePost true "创建帖子参数"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseStruct
// @Router /CreatePost [post]
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
// @Summary 获取帖子详情
// @Description 用于获取帖子详情接口 内部调用grpc接口
// @Tags 帖子相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query model.ParamPostID true "帖子ID"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseStruct
// @Router /GetPostDetail [get]
func GetPostDetail(c *gin.Context) {
	// 参数校验
	rq := new(model.ParamPostID)
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
	r, err := cc.GetPostDetail(ctx, &pbPost.GetPostDetailRequest{
		PostID: rq.ID,
	})
	if err != nil {
		zap.L().Error("GetPostDetail", zap.Error(err))
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
		return
	}
	c.JSON(http.StatusOK, gmodel.GinResponse((*pb.Response)(r)))
	return
}
