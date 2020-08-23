/******
** @创建时间 : 2020/8/17 21:14
** @作者 : SongZhiBin
******/
package redis

import "strconv"

// 利用redis建立一个多端登录的限制
const UserToken = "UToken"

// SetToken:设置Token
func SetToken(uid int64, accessToken string) {
	// 根据uid设置与用户对应的唯一标识
	rdb.HSet(UserToken, strconv.FormatInt(uid, 10), accessToken)
}

// GetToken:获取Token跟目前的token进行对比
func GetToken(uid int64, accessToken string) bool {
	// 获取uid对应的token 跟accessToken进行对比
	old, err := rdb.HGet(UserToken, strconv.FormatInt(uid, 10)).Result()
	if err != nil {
		return false
	}
	return old == accessToken
}
