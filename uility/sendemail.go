package uility

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"io"
	"log"
	"net/http"
	"net/smtp"
)

func SendErrorEmail(code, ip string, errormessage ErrorMessage) {
	e := email.NewEmail()
	//发送者
	e.From = "Cloud-k服务器 <2673893724@qq.com>"
	//接收者
	e.To = []string{code}
	//主题
	e.Subject = "Cloud-k服务器报警信息，请尽快处理"
	//html
	e.HTML = []byte("<!DOCTYPE html>\n<html>\n<head>\n" +
		"<meta charset=\"UTF-8\">\n<title>服务器报警信息请查收</title>\n" +
		"<style>\nbody {\nfont-family: Arial, sans-serif;\nbackground-color: #f4f4f4;\nmargin: 0;\npadding: 0;\n}\n\n   " +
		" .container {\n        max-width: 600px;\n        margin: 20px auto;\n        background-color: #fff;\n        padding: 20px;\n        border-radius: 5px;\n        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);\n    }\n\n    h1 {\n        font-size: 24px;\n        margin-top: 0;\n        color: #333;\n    }\n\n    p {\n        font-size: 16px;\n        margin-bottom: 20px;\n        color: #555;\n    }\n\n    .error-details {\n        font-size: 14px;\n        margin-top: 30px;\n        padding: 10px;\n        background-color: #f8f8f8;\n        border-radius: 5px;\n        border: 1px solid #eee;\n        color: #777;\n    }\n\n    .footer {\n        margin-top: 30px;\n        text-align: center;\n        color: #999;\n    }\n</style>\n</head>\n<body>\n<div class=\"container\">\n<h1>服务器报错报警</h1>\n<p>尊敬的管理员：</p>\n<p>您的服务器发生了一个错误，请立即处理以下问题：</p>\n<div class=\"error-details\">\n" +
		"<p>错误类型：" + errormessage.ErrorType + "</p>\n<p>错误描述：" + errormessage.ErrorDescription + "</p>\n" +
		"<p>错误时间：" + errormessage.ErrorTime.Format("2006 01 02 15:04:05") + "</p>\n" +
		"<p>错误详情：" + errormessage.ErrorDetails + "</p>\n" +
		"<p>服务器IP：" + ip + "</p>\n" +
		"</div>\n<p class=\"footer\">此邮件为系统自动发送，请勿回复。</p>\n</div>\n</body>\n</html>")

	err := e.SendWithStartTLS("smtp.qq.com:587", smtp.PlainAuth("", "邮箱", "授权码", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.gmail.com:465"})

	if err != nil {
		log.Println("stmp:", err)

	}
	log.Println("发送成功！")
}

func SendEmail(emails, code string) {
	e := email.NewEmail()
	//发送者
	e.From = "小学生 <2673893724@qq.com>"
	//接收者
	e.To = []string{emails}
	//主题
	e.Subject = "登录验证码"
	//文本
	e.Text = []byte("[小学生]您的登录验证码为：" + code)

	err := e.SendWithStartTLS("smtp.qq.com:587", smtp.PlainAuth("", "邮箱", "授权码", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.gmail.com:465"})
	if err != nil {
		log.Println("stmp:", err)

	}
	log.Println("发送成功！")
}

func GetServerIP() string {
	res, err := http.Get("https://ifconfig.me/")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	ip, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(ip))
	return string(ip)
}
