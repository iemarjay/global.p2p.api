package notification

import "global.p2p.api/app/notification/messages"

type MailMessage interface {
	Message
	ToMail() (*messages.MailMessageData, MailNotifiable)
}

type Message interface {
	Channels() []Channel
}


