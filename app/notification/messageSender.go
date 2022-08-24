package notification

type ContractMessageSender interface {
	Send(message Message) error
}

type messageSender struct {
}


func (n *messageSender) Send(message Message) error {
	var err error
	for _, channel := range message.Channels() {
		err = channel.Send()
	}

	return err
}

func New() *messageSender {
	return &messageSender{}
}
