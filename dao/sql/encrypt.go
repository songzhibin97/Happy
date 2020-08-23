/******
** @创建时间 : 2020/8/16 15:32
** @作者 : SongZhiBin
******/
package sql

import (
	"crypto/md5"
	"encoding/hex"
)

// 加密

const Sweet = "Happy"

func GetEncrypt(s string) string {
	h := md5.New()
	h.Write([]byte(Sweet))
	return hex.EncodeToString(h.Sum([]byte(s)))
}
