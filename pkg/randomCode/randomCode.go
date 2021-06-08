package randomCode

import (
	"math/rand"
	"time"
)

// 生成随机码

// CodeType:用于区分验证码模式类型
type CodeType int

// ReRange:返回有效范围
type ReRange [2]int

const CodeLen = 8 // 秘钥长度

var RandomList = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F',
	'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

const (
	CodeModeDigital CodeType = iota // 纯数字
	CodeModeLetter                  // 纯字母
	CodeModeMixing                  // 混合
)

// GetRange:获取验证码范围
func GetRange(c CodeType) *ReRange {
	switch c {
	case CodeModeDigital:
		return &ReRange{0, 9}
	case CodeModeLetter:
		return &ReRange{10, 62}
	case CodeModeMixing:
		return &ReRange{0, 62}
	default:
		return &ReRange{0, 62}
	}
}

func GetCode(c CodeType) string {
	r := GetRange(c)
	res := ""
	for i := 0; i < CodeLen; i++ {
		// Random
		ro := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := ro.Intn(r[1]-r[0]) + r[0]
		res = res + string(RandomList[index])
	}
	return res
}
