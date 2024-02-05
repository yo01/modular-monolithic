package helper

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"modular-monolithic/module/v1/transaction/dto"

	"go.uber.org/zap"
)

func GenerateInvoiceHTML(data dto.Email) (string, error) {
	const invoiceTemplate = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Invoice</title>
	</head>
	<body>
		<h1>Invoice</h1
		<p>Dear Yohanes</p>
		<p>Amount due: Rp. 10.000</p>
		<p>Thank you for your business!</p>
	</body>
	</html>
	`

	tmpl, err := template.New("invoice").Parse(invoiceTemplate)
	if err != nil {
		zap.S().Error(err)
		return "", err
	}

	var tplBuffer bytes.Buffer
	if err = tmpl.Execute(&tplBuffer, data); err != nil {
		zap.S().Error(err)
		return "", err
	}

	return tplBuffer.String(), nil
}

func SendEmail(config dto.Email) error {
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPServer)

	message := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-version: 1.0;\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\";\r\n"+
		"\r\n"+
		"%s", config.RecipientEmail, config.SubjectEmail, config.BodyEmail)

	if err := smtp.SendMail(config.SMTPServer+":"+fmt.Sprintf("%v", config.SMTPPort), auth, config.SenderEmail, []string{config.RecipientEmail}, []byte(message)); err != nil {
		zap.S().Error(err)
		return err
	}

	return nil
}
