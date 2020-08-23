/******
** @创建时间 : 2020/8/20 09:05
** @作者 : SongZhiBin
******/
package email

import (
	"go.uber.org/zap"
	"net/smtp"
)

// 关于验证码的第三方插件

/*
常用邮箱：
QQ 邮箱
POP3 服务器地址：qq.com（端口：995）
SMTP 服务器地址：smtp.qq.com（端口：25/465/587）

163 邮箱：
POP3 服务器地址：pop.163.com（端口：110）
SMTP 服务器地址：smtp.163.com（端口：25）

126 邮箱：
POP3 服务器地址：pop.126.com（端口：110）
SMTP 服务器地址：smtp.126.com（端口：25）
*/

// 定义了一些邮箱服务器常量
const (
	QQHost         = "smtp.qq.com"
	QQServerAddr   = "smtp.qq.com:25"
	_163Host       = "smtp.163.com"
	_163ServerAddr = "smtp.163.com:25"
	_126Host       = "smtp.126.com"
	_125ServerAddr = "smtp.126.com:25"
)

// EmailAuth:邮箱认证
type EmailAuth struct {
	Host     string // 服务器地址
	Server   string // 服务器地址:端口
	Auth     string // 用户名
	Password string // 秘钥(非密码)
}

// Message:邮件详情
type Message struct {
	To      string // 收件人地址
	Message []byte // 发送内容
}

var GE = EmailAuth{
	Host:     QQHost,
	Server:   QQServerAddr,
	Auth:     "718428482@qq.com",
	Password: "qprdvdxqnhhvbcdg",
}

// Send:发送邮件
func (e *EmailAuth) Send(m *Message) {
	auth := smtp.PlainAuth("", e.Auth, e.Password, e.Host)
	err := smtp.SendMail(e.Server, auth, e.Auth, []string{m.To}, m.Message)
	if err != nil {
		zap.L().Error("SendMail Error", zap.Error(err))
		return
	}
}
