package smsexpress

// Errors after sending SMS
type Errors []error

func (e Errors) Error() string {
	msg := ""
	for _, err := range e {
		msg += err.Error() + "\n"
	}
	return msg
}
