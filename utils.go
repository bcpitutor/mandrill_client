package main

func (m *Message) AddReceipient(email string, name string, sendType string) {
	to := &To{email, name, sendType}
	m.To = append(m.To, to)
}

// ConvertMapToVariables converts a regular string/string map into the Variable struct
// e.g. `vars := ConvertMapToVariables(map[string]interface{}{"bob": "cheese"})`
func ConvertMapToVariables(i interface{}) []*Variable {
	imap := map[string]interface{}{}

	switch i.(type) {
	// Handle older API for passing just map[string]string
	case map[string]string:
		for k, v := range i.(map[string]string) {
			imap[k] = v
		}
	case map[string]interface{}:
		imap, _ = i.(map[string]interface{})
	default:
		return []*Variable{}
	}

	variables := make([]*Variable, 0, len(imap))
	for k, v := range imap {
		variables = append(variables, &Variable{k, v})
	}
	return variables
}
