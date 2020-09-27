/******
** @创建时间 : 2020/8/27 14:20
** @作者 : SongZhiBin
******/
package model

import "time"

// 关于Post结构体信息

// CreatePost:创建帖子的一些字段
type CreatePost struct {
	CommunityID int64  `json:"community_id" binding:"required"` // 社区ID*
	Title       string `json:"title" binding:"required"`        // 标题*
	Content     string `json:"content,omitempty"`               // 内容
}

// ParamPost:用于Get请求解析字段
type ParamPost struct {
	ID   int64 `from:"id" binding:"required"`    // 对应模式的ID(社区ID/用户ID)*
	Mode int   `from:"mode" binding:"oneof=0 1"` // 查询模式*
	Page int   `from:"page"`                     // 对应页码*
	Max  int   `from:"max" binding:"max=100"`    // 每页最大数量*
}

// ParamPostID:获取帖子ID
type ParamPostID struct {
	ID int64 `form:"id" binding:"required"` // 帖子ID*
}

// Post:关于帖子的一些字段
type Post struct {
	ID          int64     `json:"post_id,string" db:"post_id" binding:"required"`
	AuthorID    int64     `json:"author_id,string" db:"author_id" binding:"required"`
	CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	*Votes
}

// PostDetail:帖子详情
type ApiPostDetail struct {
	AuthorName string `json:"author_name" db:"author_name"`
	*Post
	*CommunityDetail `json:"community" db:"community"`
}
