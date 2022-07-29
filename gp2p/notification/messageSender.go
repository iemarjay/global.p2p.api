package notification

type ContractMessageSender interface {
	Send(message Message)
}

type messageSender struct {
}


func (n *messageSender) Send(message Message) {
	for _, channel := range message.Channels() {
		channel.Send()
	}
}

func New() *messageSender {
	return &messageSender{}
}
