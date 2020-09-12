/******
** @创建时间 : 2020/8/15 22:40
** @作者 : SongZhiBin
******/
package model

// 用户相关

// RegisterForm:注册使用
type RegisterForm struct {
	UserName         string `json:"username" binding:"required,gte=4,lt=20"`                          // 用户名*
	Password         string `json:"password" binding:"required,gte=6,lt=20"`                          // 密码*
	ConfirmPassword  string `json:"confirm_password" binding:"required,eqfield=Password,gte=6,lt=20"` // re密码*
	Email            string `json:"email" binding:"required,email"`                                   // 邮箱*
	VerificationCode string `json:"verification_code" binding:"required"`                             // 验证码*
}

// LoginGet:登录请求使用
type LoginGet struct {
	UserName string `json:"username" binding:"required,gte=4,lt=20"` // 用户名*
	Password string `json:"password" binding:"required,gte=6,lt=20"` // 密码*
}

// User:用与校验是否登陆成功
type User struct {
	UID      int64  `db:"user_id"`
	Username string `db:"username"`
}

// Email:邮箱接口
type Email struct {
	Addr string `json:"email" binding:"required,email"` // 发送地址*
}
