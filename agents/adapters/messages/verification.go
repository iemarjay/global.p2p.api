package messages

import (
	"global.p2p.api/agents/helpers"
	"global.p2p.api/app"
	"global.p2p.api/app/notification"
	"global.p2p.api/app/notification/channels"
	"global.p2p.api/app/notification/messages"
)

type VerificationMessageNotifiable interface {
	notification.MailNotifiable
	OtpKeyForEmail() string
	OtpKeyForPhone() string
}

type Verification struct {
	mailChannel notification.MailChannel
	notifiable   VerificationMessageNotifiable
	sender      notification.ContractMessageSender
	via         map[string][]notification.Channel
	channels    []notification.Channel
	otp         *helpers.OtpGenerator
}

func (m *Verification) SetTo(notifiable VerificationMessageNotifiable) *Verification {
	m.notifiable = notifiable
	return m
}

func (m *Verification) Channels() []notification.Channel {
	return m.channels
}

func (m *Verification) ToMail() (*messages.MailMessageData, notification.MailNotifiable) {
	token, err := m.otp.TokenFor(m.notifiable.OtpKeyForEmail())
	if err != nil {
		return nil, nil
	}

	mm := &messages.MailMessageData{
		Subject: "[GP2P] Verification Code",
		Body:    "Your verification code: " + token,
	}

	return mm, m.notifiable
}

func (m *Verification) Send() {
	m.channels = m.via["both"]
	_ = m.sender.Send(m)
}

func (m *Verification) SendEmail() {
	m.channels = m.via["email"]
	_ = m.sender.Send(m)
}

func (m *Verification) SendPhone() {
	m.channels = m.via["phone"]
	_ = m.sender.Send(m)
}

func (m *Verification) SendVia(route string) {
	if "phone" == route {
		m.SendPhone()
	} else if "email" == route {
		m.SendEmail()
	} else if "both" == route {
		m.Send()
	}
}

func NewVerificationMessage(env *app.Env, sender notification.ContractMessageSender, otp *helpers.OtpGenerator) *Verification {
	mail := channels.NewMailGun(env)
	v := &Verification{
		mailChannel: mail,
		sender:      sender,
		otp:         otp,
	}
	v.via = map[string][]notification.Channel{
		"email": {mail.SetMessage(v)},
		"phone": {},
		"both":  {mail.SetMessage(v)},
	}
	return v
}
