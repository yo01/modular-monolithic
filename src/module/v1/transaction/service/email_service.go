package service

import (
	"net/http"

	"modular-monolithic/config"
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/module/v1/transaction/helper"

	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"
)

func SendEmail() (merr merror.Error) {
	// GET DATA CONFIG
	config := config.Get()

	// dummy data
	data := dto.Email{
		SMTPServer:     config.SMTPServer,
		SMTPPort:       config.SMTPPort,
		SMTPUsername:   config.SMTPUsername,
		SMTPPassword:   config.SMTPPassword,
		SenderEmail:    "yohaneslie0140@gmail.com",
		RecipientEmail: "yohaneslie0140@gmail.com",
		SubjectEmail:   "testing lagi",
	}

	// Create HTML email body using the invoice template
	emailBody, err := helper.GenerateInvoiceHTML(data)
	if err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	data.BodyEmail = emailBody

	// Send email
	if err = helper.SendEmail(data); err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	return
}
