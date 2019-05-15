package model

import (
	"context"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

// GaeMail store the mail infos
type GaeMail struct {
	Ctx      context.Context
	To       string
	CC       string
	BCC      string
	Subject  string
	Body     string
	HTMLBody string
}

// Send will send mail
func (gaeMail *GaeMail) Send() (err error) {
	ctx := gaeMail.Ctx
	msg := &mail.Message{
		Sender:  "noreply@" + appengine.AppID(ctx) + ".appspotmail.com",
		To:      strings.Split(gaeMail.To, ","),
		Cc:      strings.Split(gaeMail.CC, ","),
		Bcc:     strings.Split(gaeMail.BCC, ","),
		Subject: gaeMail.Subject,
		Body:    gaeMail.Body,
	}

	if err = mail.Send(ctx, msg); err != nil {
		log.Errorf(ctx, "Couldn't send email: %v", err)
	} else {
		log.Infof(ctx, "Mail send to %s", gaeMail.To)
	}
	return
}
