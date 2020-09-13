/******
** @创建时间 : 2020/9/13 15:16
** @作者 : SongZhiBin
******/
package model

// 投票相关模型

type ParamVoteDate struct {
	PostID    int64 `json:"post_id,string" binding:"required"`               // 帖子id
	Direction int   `json:"direction,string" binding:"required,oneof=0 1 2"` // 投票 反对 取消 赞成
}
