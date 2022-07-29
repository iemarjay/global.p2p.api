package notification

type Channel interface {
	Send()
}

type MailChannel interface {
	Channel
	SetMessage(message MailMessage) MailChannel
}

