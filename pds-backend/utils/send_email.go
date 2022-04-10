package utils

import "net/smtp"

func SendEmail(subject string, emailBody string, to string, from string, password string) error {
	body := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		emailBody

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	toReceiver := []string{to}

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, toReceiver, []byte(body))
	if err != nil {
		return err
	}
	return nil
}
