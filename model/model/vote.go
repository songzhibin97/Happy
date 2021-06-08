package model

// 投票相关模型

// ParamVoteDate:参数
type ParamVoteDate struct {
	PostID    int64 `json:"post_id,string" binding:"required"`      // 帖子id
	Direction int   `json:"direction,string" binding:"oneof=0 1 2"` // 投票 反对 取消 赞成
}

// Votes:点赞和踩
type Votes struct {
	Like   int `json:"like"`
	Unlike int `json:"unlike"`
}

// GVoted:返回HSet中存储PostID的信息
type GVoted struct {
	T int64 // 文章创建时间 - 服务创建时间
	X int   // 点赞数
	Y int   // 踩
}
