package smtp

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
	"strings"
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

	msg, err := buildMsgFromTemplate(email)
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
		err := smtp.SendMail(c.Host+":"+c.Port, auth, email.From, createRecipientList(email), []byte(msg))
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

func buildMsgFromTemplate(email Email) (string, error) {
	commonDir := filepath.Join(TemplateDir, SummaryDirName)
	templateFiles, err := filepath.Glob(commonDir + "/*.tmpl")
	if err != nil {
		return "", err
	}

	t, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return "", err
	}

	t, err = t.ParseFiles(filepath.Join(TemplateDir, email.TemplateFileName))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.ExecuteTemplate(buf, email.TemplateFileName, email.Data); err != nil {
		return "", err
	}

	header := make(map[string]string)
	header["From"] = fmt.Sprintf("%s <%s>", email.FromName, email.From)
	header["To"] = strings.Join(email.To, ",")
	header["Cc"] = strings.Join(email.Cc, ",")
	header["Subject"] = email.Subject
	header["Content-Type"] = `text/html; charset="UTF-8"`

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	msg += "\r\n" + buf.String()
	return msg, nil
}

func createRecipientList(email Email) []string {
	recipients := append(email.To, email.Cc...)
	recipients = append(recipients, email.Bcc...)
	return recipients
}

func sendEmail(client *smtp.Client, email Email, msg string) error {
	if err := client.Mail(email.From); err != nil {
		return err
	}

	for _, addr := range createRecipientList(email) {
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
