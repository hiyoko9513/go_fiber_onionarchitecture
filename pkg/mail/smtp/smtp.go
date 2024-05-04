package smtp

import (
	"crypto/tls"
	"errors"
	"net/smtp"
)

const (
	TemplateDir    = "./pkg/mail/templates"
	SummaryDirName = "common"
)

type Config struct {
	Host       string
	Port       string
	User       string
	Password   string
	TLSEnabled bool
	AuthMethod string
}

func (c *Config) Send(email Email) error {
	auth, err := c.buildAuth()
	if err != nil {
		return err
	}

	email.buildData()
	msg, err := email.buildMsgFromTemplate()
	if err != nil {
		return err
	}

	if c.TLSEnabled {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         c.Host,
		}

		conn, err := tls.Dial("tcp", c.Host+":"+c.Port, tlsConfig)
		if err != nil {
			return err
		}

		client, err := smtp.NewClient(conn, c.Host)
		if err != nil {
			return err
		}

		defer client.Close()
		if err := client.Auth(auth); err != nil {
			return err
		}

		if err := sendEmail(client, email, msg); err != nil {
			return err
		}
	} else {
		err := smtp.SendMail(c.Host+":"+c.Port, auth, email.From, email.createRecipientList(), []byte(msg))
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) buildAuth() (smtp.Auth, error) {
	if c.AuthMethod == "PLAIN" {
		return smtp.PlainAuth("", c.User, c.Password, c.Host), nil
	}
	return nil, errors.New("unsupported auth method")
}

func sendEmail(client *smtp.Client, email Email, msg string) error {
	if err := client.Mail(email.From); err != nil {
		return err
	}

	for _, addr := range email.createRecipientList() {
		if err := client.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return client.Quit()
}
