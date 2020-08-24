/******
** @创建时间 : 2020/8/15 22:40
** @作者 : SongZhiBin
******/
package model

// 用户相关

// RegisterForm:注册使用
type RegisterForm struct {
	UserName         string `json:"username" binding:"required,gte=4,lt=20"`
	Password         string `json:"password" binding:"required,gte=6,lt=20"`
	ConfirmPassword  string `json:"confirm_password" binding:"required,eqfield=Password,gte=6,lt=20"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verification_code" binding:"required"`
}

// LoginGet:登录请求使用
type LoginGet struct {
	UserName string `json:"username" binding:"required,gte=4,lt=20"`
	Password string `json:"password" binding:"required,gte=6,lt=20"`
}

// User:用与校验是否登陆成功
type User struct {
	UID      int    `db:"user_id"`
	Username string `db:"username"`
}

// Email:邮箱接口
type Email struct {
	Addr string `json:"email" binding:"required,email"`
}
