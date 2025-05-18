package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/espitman/jbm-hr-backend/utils/config"
)

var (
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
	senderEmail  string
)

func init() {
	// Load SMTP configuration from environment variables using config package
	smtpHost = config.GetConfig("SMTP_HOST", "smtp.c1.liara.email")

	portStr := config.GetConfig("SMTP_PORT", "465")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 465 // default port
	}
	smtpPort = port

	smtpUsername = config.GetConfig("SMTP_USERNAME", "sharp_aryabhata_4kbvvi")
	smtpPassword = config.GetConfig("SMTP_PASSWORD", "f38a7571-046b-4ab3-8d7e-da3c2d3aa3c2")
	senderEmail = config.GetConfig("SMTP_SENDER_EMAIL", "life@jabama.org")
}

// SendEmail sends an email using SMTP
func SendEmail(to, subject, body string) error {
	// Create auth
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// Connect to the SMTP server
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", smtpHost, smtpPort), tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()

	// Create SMTP client
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()

	// Authenticate
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	// Set sender
	if err := client.Mail(senderEmail); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	// Set recipient
	if err := client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	// Create message
	message := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", senderEmail, to, subject, body)

	// Send message
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to create data writer: %w", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}

	return nil
}

// SendOTPEmail sends an OTP code to a user's email
func SendOTPEmail(to, otp string) error {
	subject := "فهرست آشنایی - کد ورود یکبار مصرف"
	body := fmt.Sprintf(`
		<html>
			<head>
				<meta charset="UTF-8">
				<link rel="preconnect" href="https://fonts.googleapis.com">
				<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
				<link href="https://fonts.googleapis.com/css2?family=Vazirmatn:wght@100..900&display=swap" rel="stylesheet">
				<style>
					* {
						direction: rtl;
					}
					body {
						font-family: "Vazirmatn", sans-serif;
						font-optical-sizing: auto;
						font-weight: 400;
						font-style: normal;
						font-size: 16px;
						line-height: 1.6;
						color: #333;
						direction: rtl;
						margin: 0;
						padding: 0;
						background-color: #f8f9fa;
					}
					.container {
						max-width: 600px;
						margin: 0 auto;
						padding: 20px;
						background-color: #ffffff;
						border-radius: 8px;
						box-shadow: 0 2px 4px rgba(0,0,0,0.1);
					}
					.header {
						text-align: center;
						margin-bottom: 30px;
					}
					.header img {
						width: 100%%;
						max-width: 600px;
						height: auto;
						border-radius: 8px;
					}
					.otp-code {
						background-color: #f8f9fa;
						padding: 15px;
						border-radius: 4px;
						text-align: center;
						font-size: 24px;
						font-weight: 700;
						letter-spacing: 4px;
						margin: 20px 0;
						direction: ltr;
					}
					.footer {
						margin-top: 30px;
						padding-top: 20px;
						border-top: 1px solid #e9ecef;
						font-size: 14px;
						color: #6c757d;
					}
					h2 {
						font-weight: 600;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<div class="header">
						<img src="https://jabama-files.storage.c2.liara.space/images/email/otp-banner.jpg" alt="OTP Banner">
					</div>
					<h2 style="text-align: center; color: #212529;">کد ورود یکبار مصرف شما</h2>
					<div class="otp-code">%s</div>
					<p style="text-align: center; color: #6c757d;">این کد در سه دقیقه دیگر منقضی می شود.</p>
					<div class="footer">
						<p style="text-align: center;">اگر این کد را درخواست نکرده اید، لطفا این ایمیل را نادیده بگیرید.</p>
					</div>
				</div>
			</body>
		</html>
	`, otp)

	return SendEmail(to, subject, body)
}
