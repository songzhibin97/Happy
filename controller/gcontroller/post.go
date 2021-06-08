package gcontroller

import (
	"Happy/dao/redis"
	sqls "Happy/dao/sql"
	"Happy/middleware"
	"Happy/model/gmodel"
	"Happy/model/model"
	pbPost "Happy/model/pmodel/post"
	"Happy/pkg/snowflake"
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"go.uber.org/zap"
)

// Post:帖子相关
type Post struct{}

// todo 还需要做的事情 : PostList获取帖子列表 增加flag 根据多种不同热度获取
// todo 主要工作在redis上的优化

// PostList:获取帖子列表
func (p Post) PostList(ctx context.Context, request *pbPost.GetPostListRequest) (*pbPost.Response, error) {
	// 1.校验参数
	res, err := _verification(request)
	if err != nil {
		return (*pbPost.Response)(res), nil
	}
	// 判断state是什么模式进入的 根据社区区分还是根据用户区分
	var postList []*model.Post

	switch request.Model {
	case pbPost.GetPostListRequest_Community:
		// 以社区筛选
		cid := request.GetCommunityID()
		postList, err = sqls.GetPostListToCid(cid, request.Page, request.Max)
	case pbPost.GetPostListRequest_Author:
		// 以用户筛选
		uid := request.GetAuthorID()
		postList, err = sqls.GetPostListToUid(uid, request.Page, request.Max)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return (*pbPost.Response)(gmodel.ResponseError(model.CodeGetListEmpty)), nil
		}
		zap.L().Error("sql.GetPostList Error", zap.Error(err))
		return (*pbPost.Response)(gmodel.ResponseError(model.CodeServerBusy)), nil
	}
	jaPost, err := json.Marshal(postList)
	if err != nil {
		return (*pbPost.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "Json Marshal Error")), nil
	}
	return (*pbPost.Response)(gmodel.ResponseSuccess(map[string]string{"Post_list": string(jaPost)})), nil
}

// CreatePost:新建帖子
func (p Post) CreatePost(ctx context.Context, request *pbPost.CreatePostRequest) (*pbPost.Response, error) {
	// 1.校验参数
	res, err := _verification(request)
	if err != nil {
		return (*pbPost.Response)(res), nil
	}
	// 获取上下文的用户id
	iid := ctx.Value(middleware.ConTextUserID)
	uid, ok := iid.(int64)
	if !ok {
		zap.L().Error("转换用户ID失败 ctx.Value(middleware.ConTextUserID)")
		return (*pbPost.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "ID转化失败")), nil
	}
	// 生成全局id
	postId, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("生成全局Error", zap.Error(err))
		return (*pbPost.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "全局ID生成失败")), nil
	}
	// 插入数据库
	err = sqls.CreatePost(int64(postId), uid, request.CommunityID, request.Title, request.Content)
	if err != nil {
		zap.L().Error("sqls.CreatePost Error", zap.Error(err))
		return (*pbPost.Response)(gmodel.ResponseError(model.CodePostError)), nil
	}
	// 新增:成功创建post后建立对应redis对列
	redis.CreatePost(int64(postId), uid)
	// 成功 返回postId
	return (*pbPost.Response)(gmodel.ResponseSuccess(map[string]string{"post_id": strconv.FormatUint(postId, 10)})), nil
}

// GetPostDetail:获取帖子详情
func (p Post) GetPostDetail(ctx context.Context, request *pbPost.GetPostDetailRequest) (*pbPost.Response, error) {
	// 1.参数校验
	res, err := _verification(request)
	if err != nil {
		return (*pbPost.Response)(res), nil
	}
	// 2.根据postId获取帖子信息
	aPost, err := sqls.GetApiPostDetail(request.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return (*pbPost.Response)(gmodel.ResponseError(model.CodeGetListEmpty)), nil
		}
		zap.L().Error("sqls.GetApiPostDetail Error", zap.Error(err))
		return (*pbPost.Response)(gmodel.ResponseError(model.CodeServerBusy)), nil
	}
	jaPost, err := json.Marshal(aPost)
	if err != nil {
		return (*pbPost.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "Json Marshal Error")), nil
	}
	return (*pbPost.Response)(gmodel.ResponseSuccess(map[string]string{"PostDetail": string(jaPost)})), nil
}
