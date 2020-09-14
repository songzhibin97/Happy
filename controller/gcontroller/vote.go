/******
** @创建时间 : 2020/9/14 09:37
** @作者 : SongZhiBin
******/
package gcontroller

import (
	"Happy/dao/redis"
	"Happy/middleware"
	"Happy/model/gmodel"
	"Happy/model/model"
	pbVote "Happy/model/pmodel/vote"
	"context"
	"go.uber.org/zap"
)

type Vote struct{}

// Vote:投票
func (v Vote) Vote(ctx context.Context, request *pbVote.VoteRequest) (*pbVote.Response, error) {
	// 1.校验参数
	res, err := _verification(request)
	if err != nil {
		return (*pbVote.Response)(res), nil
	}

	// 获取上下文的用户id
	iid := ctx.Value(middleware.ConTextUserID)
	uid, ok := iid.(int64)
	if !ok {
		zap.L().Error("转换用户ID失败 ctx.Value(middleware.ConTextUserID)")
		return (*pbVote.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "ID转化失败")), nil
	}
	// 2.判断PostId是否有效
	ok = redis.PostIdExist(request.PostID)
	if !ok {
		// 不存在
		return (*pbVote.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "文章ID无效")), nil
	}
	if redis.JudgeAuth(request.PostID, uid) {
		return (*pbVote.Response)(gmodel.ResponseError(model.CodeAuthNoVote)), nil
	}
	// 3.判断UserID是否投过票 进行case分支
	pow, ok := redis.GetUserVote(request.PostID, uid)
	dir := int32(request.GetMode() - 1)
	// 更新用户表
	redis.AddUserVote(request.PostID, uid, int64(dir))
	if !ok {
		// 说明没投过票
		// 这里做一下投票操作 更新一下分数
		// 这里先做一个特殊判断 如果 == 0 代表不投票直接完成

		// 投票操作
		if dir == 0 {
			return (*pbVote.Response)(gmodel.ResponseSuccess(nil)), nil
		}
		if dir < 0 {
			// 投反对票
			redis.ChangeY(request.PostID, 1)
		} else {
			redis.ChangeX(request.PostID, 1)
		}
		if !redis.ChangeScore(request.PostID) {
			return (*pbVote.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "分数更新失败")), nil
		}
		return (*pbVote.Response)(gmodel.ResponseSuccess(nil)), nil
	}
	// 这里做一个简单的判断 如果获取的投票状态与现在的投票状态一直 不做任何处理
	if dir == int32(pow) {
		return (*pbVote.Response)(gmodel.ResponseError(model.CodeRepeatVoting)), nil
	}
	switch pow {
	case -1:
		// 原来是反对票
		if dir == 0 {
			// 取消反对票
			redis.ChangeY(request.PostID, -1)
		}
		if dir == 1 {
			// 转投赞成
			redis.ChangeY(request.PostID, -1)
			redis.ChangeX(request.PostID, 1)
		}
	case 0:
		// 没投票
		if dir == -1 {
			redis.ChangeY(request.PostID, 1)
		}
		if dir == 1 {
			redis.ChangeX(request.PostID, 1)
		}
	case 1:
		// 原来是赞成票
		if dir == 0 {
			redis.ChangeX(request.PostID, -1)
		}
		if dir == -1 {
			redis.ChangeX(request.PostID, -1)
			redis.ChangeY(request.PostID, 1)
		}
	}
	// 刷新分数
	if !redis.ChangeScore(request.PostID) {
		return (*pbVote.Response)(gmodel.ResponseWithMsg(model.CodeServerBusy, "分数更新失败")), nil
	}
	return (*pbVote.Response)(gmodel.ResponseSuccess(nil)), nil
}
