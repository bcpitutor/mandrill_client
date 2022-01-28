package mandrill

func (m *Message) AddReceipient(email string, name string, sendType string) {
	to := &To{email, name, sendType}
	m.To = append(m.To, to)
}
