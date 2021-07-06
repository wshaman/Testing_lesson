package notification

type Notifier interface {
	NotifyOnRegister(recipient string) (string, error)
}

type Client struct {
	Address string
}

type Messender struct {
	notifier Notifier
}

func NewMessenger(n Notifier) Messender {
	return Messender{
		notifier: n,
	}
}

func (m *Messender) Register(c Client) error {
	_, err := m.notifier.NotifyOnRegister(c.Address)
	return err
}
