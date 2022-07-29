package channels

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"global.p2p.api/gp2p"
	"global.p2p.api/gp2p/notification"
	"time"
)

type mailGun struct {
	domain string
	privatKey string
	message   notification.MailMessage
}

func NewMailGun(env *gp2p.Env) *mailGun {
	return &mailGun{
		domain: env.Get("MAILGUN_DOMAIN"),
		privatKey: env.Get("MAILGUN_PRIVATE_KEY"),
	}
}

func (mg mailGun) SendMail(notifiable notification.MailNotifiable, message notification.MailMessage) {
	mg.Send()
}

func (mg mailGun) Send() {
	mailMessage, notifiable := mg.message.ToMail()
	client := mailgun.NewMailgun(mg.domain, mg.privatKey)

	to := mailMessage.To
	if to == "" {
		to = notifiable.RouteForMail()
	}

	clientMessage := client.NewMessage(mailMessage.From, mailMessage.Subject, mailMessage.Body, to)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, _ = client.Send(ctx, clientMessage)
}

func (mg mailGun) SetMessage(message notification.MailMessage) notification.MailChannel {
	mg.message = message
	return mg
}