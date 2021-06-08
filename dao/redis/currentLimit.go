package redis

import "time"

// 限流 放置恶意调用开放接口

const (
	CurrentLimitTimeOut = time.Minute
)

// 限流系列

// SetCurrentLimit:加入访问限流
func SetCurrentLimit(ip string, code string) {
	StrAdd(rdb, KeyCurrentLimitPF, ip, code, CurrentLimitTimeOut)
}

// GetCurrentLimit:查询是否已经被限流
func GetCurrentLimit(ip string) bool {
	return IsStr(rdb, KeyCurrentLimitPF, ip)
}

// GetVerificationCode:获取验证码
func GetVerificationCode(ip string) (string, bool) {
	return GetStr(rdb, KeyCurrentLimitPF, ip)
}
