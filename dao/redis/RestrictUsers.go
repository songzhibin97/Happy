/******
** @创建时间 : 2020/8/17 21:14
** @作者 : SongZhiBin
******/
package redis

import "strconv"

// 利用redis建立一个多端登录的限制

// Hash的一些方法

// SetToken:设置Token
func SetToken(uid int64, accessToken string) {
	// 根据uid设置与用户对应的唯一标识
	HashAdd(rdb, KeyUserToken, strconv.FormatInt(uid, 10), accessToken)
}

// GetToken:获取Token跟目前的token进行对比
func GetToken(uid int64, accessToken string) bool {
	// 获取uid对应的token 跟accessToken进行对比
	return HashContrast(rdb, KeyUserToken, strconv.FormatInt(uid, 10), accessToken)
}
