package model

import "time"

// CommunityList:用于返回社区列表使用
type CommunityList struct {
	ID   int64  `json:"id,string" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

// CommunityDetailRequest:请求社区信息
type CommunityDetailRequest struct {
	ID int64 `form:"id" binding:"required"` // 社区ID*
}

// CommunityDetail:用于返回社区详情
type CommunityDetail struct {
	ID           int64     `json:"id,string" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
