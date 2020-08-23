/******
** @创建时间 : 2020/8/17 12:00
** @作者 : SongZhiBin
******/
package jwt

import (
	"Happy/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	// 糖
	Sweet           = "Happy"
	TokenDuration   = time.Minute * 10
	RefreshDuration = time.Hour * 24 * 14
)

// Auth:Jwt认证结构体
type Auth struct {
	Uid int `json:"user_id"` // 记录唯一标识 uid
	jwt.StandardClaims
}

// 关于jwt的方法函数

// isValid:判断是否有效
func isValid(t int64) bool {
	if t < time.Now().Unix() {
		return false
	}
	return true
}

// GetJWT:进行声明
func GetJWT(uid int) (string, error) {
	auth := Auth{
		uid, // 自定义字段
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),                    // 发出时间
			ExpiresAt: time.Now().Add(TokenDuration).Unix(), // 过期时间
			Issuer:    "Happy",                              // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth)
	return token.SignedString([]byte(Sweet))
}

// GetReJWT:refreshJWT
func GetReJWT() (string, error) {

	J := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),                    // 发出时间
		ExpiresAt: time.Now().Add(TokenDuration).Unix(), // 过期时间
		Issuer:    "Happy",                              // 签发人
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, J)
	return token.SignedString([]byte(Sweet))
}

// ParseJWT:进行解析
func ParseJWT(jwtToken string) (*Auth, error) {
	auth := new(Auth)
	// 解析token
	token, err := jwt.ParseWithClaims(jwtToken, auth, func(token *jwt.Token) (interface{}, error) {
		return []byte(Sweet), nil
	})
	if err != nil {
		// 解析失败
		return nil, err
	}
	// 判断是否验证通过
	if !token.Valid {
		return nil, model.CodeJWTVerificationFailed.Err()
	}
	// 判断时间是否有效
	if !isValid(auth.ExpiresAt) {
		return nil, model.CodeJWTExpired.Err()
	}
	return auth, nil
}

// GetACRFToken:同时生成access token 和 refresh token
func GetACRFToken(uid int) (string, string, error) {
	// 先生产一个access token
	aToken, err := GetJWT(uid)
	if err != nil {
		return "", "", err
	}
	// 在生成一个refresh token
	reToken, err := GetReJWT()
	if err != nil {
		return "", "", err
	}
	return aToken, reToken, err
}

// ParseRFToken:解析刷新refreshToken
func ParseRFToken(aToken, rfToken string) (*Auth, error) {
	auth, err := ParseJWT(aToken)
	if err == nil {
		// 如果 err为空表示accessToken 在有效期内
		return auth, nil
	}
	if err != model.CodeJWTExpired.Err() {
		// 如果错误不是过期 就直接返回失败
		return nil, model.CodeJWTVerificationFailed.Err()
	}
	// 表示accessToken已经过期,去判断refreshToken是否在有效期内
	// 判断refreshToken是否有效
	rf := new(jwt.StandardClaims)
	_, err = jwt.ParseWithClaims(rfToken, rf, func(token *jwt.Token) (interface{}, error) {
		return []byte(Sweet), nil
	})
	if err != nil {
		// 如果解析错误直接返回错误
		return nil, model.CodeJWTVerificationFailed.Err()
	}
	// 判断时间是否有效
	if !isValid(rf.ExpiresAt) {
		// refresh已经过期返回
		return nil, model.CodeJWTExpired.Err()
	}
	// 返回原来的auth 用于生成新的token
	return auth, nil
}
