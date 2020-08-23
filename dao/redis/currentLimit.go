/******
** @创建时间 : 2020/8/20 17:11
** @作者 : SongZhiBin
******/
package redis

import "time"

// 限流 放置恶意调用开放接口

const (
	CurrentLimit        = "CurrentLimit_"
	CurrentLimitTimeOut = time.Minute
)

// 限流系列

// SetCurrentLimit:加入访问限流
func SetCurrentLimit(ip string, code string) {
	StrAdd(rdb, CurrentLimit, ip, code, CurrentLimitTimeOut)
}

// GetCurrentLimit:查询是否已经被限流
func GetCurrentLimit(ip string) bool {
	return IsStr(rdb, CurrentLimit, ip)
}

// GetVerificationCode:获取验证码
func GetVerificationCode(ip string) (string, bool) {
	return GetStr(rdb, CurrentLimit, ip)
}
