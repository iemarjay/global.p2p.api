package notification

type Channel interface {
	Send() error
}

type MailChannel interface {
	Channel
	SetMessage(message MailMessage) MailChannel
}

