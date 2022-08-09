package notification

type Notifiable interface {
}

type MailNotifiable interface {
	Notifiable
	RouteForMail() string
}
