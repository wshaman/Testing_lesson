package email

type Email struct {
}

func (e Email) NotifyOnRegister(recipient string) (string, error) {
	//sending email logic
	return "emlId:123", nil
}
