package configs

import (
	"hiyoko-fiber/pkg/mail/smtp"
	"hiyoko-fiber/utils"
)

func NewSmtpConf() (conf smtp.Config) {
	conf.Host = utils.Env("MAIL_HOST").GetString()
	conf.Port = utils.Env("MAIL_PORT").GetString()
	conf.User = utils.Env("MAIL_USERNAME").GetString()
	conf.Password = utils.Env("MAIL_PASSWORD").GetString()
	conf.TLSEnabled = utils.Env("MAIL_TLS").GetBool()
	conf.AuthMethod = "PLAIN"
	return
}

func NewEmailConf() (*smtp.Email, error) {
	email := &smtp.Email{
		From:     utils.Env("MAIL_FROM_ADDRESS").GetString(),
		FromName: utils.Env("MAIL_FROM_NAME").GetString(),
	}
	err := email.SetAssetsPngImages("./assets/images")

	return email, err
}
