package redis

import (
	"Happy/settings"
	"fmt"
)

// 存放一些共用的Redis Key

const (
	// KeyBlacklistPF:uid 黑名单
	KeyBlacklistPF = "Blacklist:" // String

	// KeyCurrentLimitPF:email 限流
	KeyCurrentLimitPF = "CurrentLimit:" // String

	// KeyUserToken: 记录最近一次用户登录有效Token 用作多端登录限制
	KeyUserToken = "UToken:" // Hash

	// 投票

	// KeyPosTime:PID 帖子的发帖时间权值
	KeyPostHash = "Post:Hash:" // Hash
	// {"t":常量,x:"点赞数",y:"踩数"}
	KeyPostHashT = "T"
	KeyPostHashX = "X"
	KeyPostHashY = "Y"
	KeyPostAuth  = "PostAuth"

	// KeyPostScore 帖子及投票分数
	KeyPostScore = "Post:Score" // ZSet

	// KeyPostVotePF:uid 用户投票情况
	KeyPostVotePF = "Post:Voted:" // ZSet
)

// Tool

// Partial:偏函数
// 用于拼接一个合法的key
func Partial(key string) string {
	return fmt.Sprintf("%s:%s", settings.GetString("APP.Name"), key)
}

// PartSplice:部分拼接
func PartSplice(front, end string) string {
	return fmt.Sprintf("%s%s", front, end)
}
