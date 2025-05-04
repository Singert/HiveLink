package email

import (
	_ "log"

	"github.com/jordan-wright/email"
)

// SMTPClient 封装SMTP客户端功能
type SMTPClient struct {
	host     string
	port     string
	username string
	password string
}

// NewSMTPClient 创建新的SMTP客户端
func NewSMTPClient(host, port, username, password string) *SMTPClient {
	return &SMTPClient{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

// SendEmail 发送邮件
func (c *SMTPClient) SendEmail(to, subject, body string, attachments []string) error {
	e := email.NewEmail()
	e.From = c.username
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	// 添加附件
	for _, attachment := range attachments {
		_, err := e.AttachFile(attachment)
		if err != nil {
			return err
		}
	}

	// // 使用OAuth2认证
	// token := &oauth2.Token{
	// 	AccessToken: c.getAccessToken(),
	// }
	// auth := smtp.PlainAuth("", c.username, c.password, c.host)

	// 发送邮件
	addr := c.host + ":" + c.port
	return e.Send(addr, nil)
}

// // getAccessToken 获取OAuth2访问令牌
// func (c *SMTPClient) getAccessToken() string {
// 	// TODO: 实现OAuth2令牌获取逻辑
// 	return ""
// }
