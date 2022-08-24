package messages

import (
	"global.p2p.api/app"
	"global.p2p.api/app/notification"
	"global.p2p.api/app/notification/channels"
	"global.p2p.api/app/notification/messages"
)

type WelcomeMessageNotifiable interface {
	notification.MailNotifiable
}

type Welcome struct {
	mailChannel notification.MailChannel
	notifiable  WelcomeMessageNotifiable
}

func (m *Welcome) SetTo(notifiable WelcomeMessageNotifiable) {
	m.notifiable = notifiable
}

func (m *Welcome) Channels() []notification.Channel {
	return []notification.Channel{ m.mailChannel.SetMessage(m) }
}

func (m *Welcome) ToMail() (*messages.MailMessageData, notification.MailNotifiable) {
	mm := &messages.MailMessageData{
		Subject: "Welcome to piggyfi",
		Body:    "Welcome to Piggyfi",
	}
	return mm, m.notifiable
}

func NewWelcomeMessage(env *app.Env) *Welcome {
	return &Welcome{
		mailChannel: channels.NewMailGun(env),
	}
}