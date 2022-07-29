package messages

import (
	"global.p2p.api/gp2p"
	"global.p2p.api/gp2p/notification"
	"global.p2p.api/gp2p/notification/channels"
	"global.p2p.api/gp2p/notification/messages"
)

type WelcomeMessageNotifiable interface {
	notification.MailNotifiable
}

type WelcomeMessage struct {
	mailChannel notification.MailChannel
	notifiable  WelcomeMessageNotifiable
}

func (m *WelcomeMessage) SetTo(notifiable WelcomeMessageNotifiable) {
	m.notifiable = notifiable
}

func (m *WelcomeMessage) Channels() []notification.Channel {
	return []notification.Channel{ m.mailChannel.SetMessage(m) }
}

func (m *WelcomeMessage) ToMail() (*messages.MailMessageData, notification.MailNotifiable) {
	mm := &messages.MailMessageData{
		Subject: "Welcome to piggyfi",
		Body:    "Welcome to Piggyfi",
	}
	return mm, m.notifiable
}

func NewWelcomeMessage(env *gp2p.Env) *WelcomeMessage {
	return &WelcomeMessage{
		mailChannel: channels.NewMailGun(env),
	}
}