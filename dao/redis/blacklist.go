/******
** @创建时间 : 2020/8/17 21:50
** @作者 : SongZhiBin
******/
package redis

import (
	"Happy/pkg/jwt"
	"strconv"
)

const (
	// TimeOut引用jwt的Timeout
	BlacklistTimeOut = jwt.RefreshDuration
)

// 黑名单系列

// SetBlacklist:加入黑名单
func SetBlacklist(uid int64) {
	StrAdd(rdb, KeyBlacklistPF, strconv.FormatInt(uid, 10), strconv.FormatInt(uid, 10), BlacklistTimeOut)

}

// GetBlacklist:查询是否在黑名单中
func GetBlacklist(uid int64) bool {
	return IsStr(rdb, KeyBlacklistPF, strconv.FormatInt(uid, 10))
}
