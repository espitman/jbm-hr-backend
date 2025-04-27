package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

const (
	smtpHost     = "smtp.c1.liara.email"
	smtpPort     = 465
	smtpUsername = "sharp_aryabhata_4kbvvi"
	smtpPassword = "f38a7571-046b-4ab3-8d7e-da3c2d3aa3c2"
	senderEmail  = "life@jabama.org"
)

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
	subject := "Your OTP Code"
	body := fmt.Sprintf(`
		<html>
			<body>
				<h2>Your OTP Code</h2>
				<p>Your OTP code is: <strong>%s</strong></p>
				<p>This code will expire in 5 minutes.</p>
				<p>If you didn't request this code, please ignore this email.</p>
			</body>
		</html>
	`, otp)

	return SendEmail(to, subject, body)
}
