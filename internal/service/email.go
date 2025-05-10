package service

import (
	"fmt"
	"net/smtp"

	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	emailTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Email Verification</title>
</head>
<body>
	<p>Dear User,</p>
	<p>Thank you for registering at <strong>axpz.org</strong>. To confirm that it was you who performed this action, please click the link below to verify your email address:</p>
	<p>
		<a href="%s" target="_blank" style="color: #1a73e8;">Click here to verify your email</a>
	</p>
	<p>If the link is not clickable, please copy and paste the following URL into your browser:</p>
	<p style="word-break: break-all;">%s</p>
	<p>This link is valid for 24 hours. After that, it will expire. If you did not initiate this request, please ignore this email.</p>
	<br>
	<p>Sincerely,</p>
	<p><strong>axpz.org Team</strong></p>
	<p style="font-size: 12px; color: #888;">This is an automated message. Please do not reply directly.</p>

	<hr>

	<p>尊敬的用户，您好：</p>
	<p>感谢您注册 <strong>axpz.org</strong>。为确保是您本人操作，请点击以下链接完成邮箱验证：</p>
	<p>
		<a href="%s" target="_blank" style="color: #1a73e8;">点击此处验证邮箱</a>
	</p>
	<p>如果链接无法点击，请将以下地址复制到浏览器中打开：</p>
	<p style="word-break: break-all;">%s</p>
	<p>该链接有效期为 24 小时，过期将无法使用。如非本人操作，请忽略此邮件。</p>
	<br>
	<p>此致</p>
	<p><strong>axpz.org 团队</strong></p>
	<p style="font-size: 12px; color: #888;">本邮件由系统自动发送，请勿直接回复。</p>
</body>
</html>`
)

// EmailService 邮件服务
type EmailService struct {
	host     string
	port     int
	username string
	password string
	from     string
}

// NewEmailService 创建邮件服务
func NewEmailService(host string, port int, username, password, from string) *EmailService {
	return &EmailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *EmailService) SendVerificationEmail(c *gin.Context, verificationLink, userEmail string) error {

	subject := "【axpz.org】Email Verification / 邮箱验证通知"

	body := fmt.Sprintf(emailTemplate, verificationLink, verificationLink, verificationLink, verificationLink)

	msg := fmt.Sprintf("To: %s\r\n"+
		"From: axpz.org <%s>\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
		"%s", userEmail, s.from, subject, body)

	logger := utils.LoggerFromContext(c.Request.Context())
	logger.Info("SendVerificationEmail", zap.String("email", userEmail), zap.String("link", verificationLink))

	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	logger.Info("SendVerificationEmail", zap.String("email", userEmail), zap.String("link", verificationLink))

	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	logger.Info("SendVerificationEmail", zap.String("email", userEmail), zap.String("link", verificationLink))

	return smtp.SendMail(addr, auth, s.from, []string{userEmail}, []byte(msg))
}
