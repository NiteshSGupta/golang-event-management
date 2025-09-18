package main

import (
	"crypto/tls"
	"reast-api-in-gin/internal/env"

	"gopkg.in/gomail.v2"
)

func otp() {
	d := gomail.NewDialer("smtp.example.com", 587, env.GetEnvString("MAIL_USERNAME", "user"), env.GetEnvString("MAIL_PASSWORD", "abcdefghy"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
}
