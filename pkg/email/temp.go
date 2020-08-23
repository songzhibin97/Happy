/******
** @创建时间 : 2020/8/20 10:16
** @作者 : SongZhiBin
******/
package email

import (
	"Happy/settings"
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"html/template"
)

// 邮件MessageTemp

func (e *EmailAuth) CreateTemp(to string, title string, code string) *Message {
	//提取邮件页面
	t, _ := template.ParseFiles("/Users/songzhibin/go/src/Songzhibin/Happy/pkg/email/email.html")
	//设置邮件格式
	//const headers = "MIMI-version: 1.0;\nContent-Type: text/html"
	headers := fmt.Sprintf("%s-version: %s;\nContent-Type: text/html", settings.GetString("App.Name"),
		settings.GetString("App.Version"))
	//
	////设置邮件基本信息
	var msg bytes.Buffer
	// 设置邮箱头信息:
	// From:xxx\n Subject:xxx\n Head\n\n
	// Content
	msg.Write([]byte(fmt.Sprintf("From:%s\nSubject:%s邮箱验证码:\n%s\n\n", "official", title, headers)))
	if err := t.Execute(&msg, struct {
		Title string
		Code  string
	}{
		Title: title,
		Code:  code,
	}); err != nil {
		zap.L().Error("template Execute", zap.Error(err))
	}
	str3 := msg.Bytes()
	return &Message{
		To:      to,
		Message: str3,
	}

	/*
		filename :="foo.db"
		//设置邮件
		boundary :="http://dojava.cn"
		mime.WriteString("From: 数据库备份文件<"+USER+">\r\nTo: "+to+"\r\nSubject: sqlite数据库备份文件\r\nMIME-Version: 1.0\r\n")
		mime.WriteString("Content-Type: multipart/mixed; boundary="+boundary+"\r\n\r\n")
		mime.WriteString("--"+boundary+"\r\n")    //自定义邮件内容分隔符
		//邮件正文
		mime.WriteString("\r\n\r\n\r\n")
		html :="备份数据已通过邮件发送到您的邮箱,请下载后用打开"  //邮件正文
		mime.WriteString("Content-Type: text/html; charset=utf-8\r\n\r\n")  //text/html html text/plain 纯文本
		mime.WriteString(html)
		//附件
		mime.WriteString("--"+boundary+"\r\n")
		mime.WriteString("Content-Type: application/vnd.ms-excel\r\n")   //application/octet-stream
		mime.WriteString("Content-Transfer-Encoding: base64\r\n")
		mime.WriteString("Content-Disposition: attachment; filename=\""+"C:/user/My Go/Task02sqlite/"+filename+"\"")
		mime.WriteString("\r\n\r\n")
		//将文件转为base64
		//读取并编码文件内容
		//attaData, err := ioutil.ReadFile("../bapi/main.go")
		fileName :="./Workbook.xls"
		attaData, err := ioutil.ReadFile(fileName)
		if err!= nil {
			fmt.Print(err)
		}
		b :=make([]byte, base64.StdEncoding.EncodedLen(len(attaData)))
		base64.StdEncoding.Encode(b, attaData)
		mime.Write(b)
		mime.WriteString("\r\n")
		mime.WriteString("--"+boundary+"--")

		str3 := mime.String()
		auth:= smtp.PlainAuth("", USER, PASSWORD, HOST)
		errs := smtp.SendMail(SERVER_ADDR,auth,USER,[]string{to}, []byte(str3))
		if errs!= nil {
			fmt.Println(errs)
		}else{
			fmt.Println("邮件发送成功!")
		}
	*/
}
