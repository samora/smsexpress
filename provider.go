package smsexpress

import (
	"net/http"
)

// Provider is SMS provider
type Provider interface {
	Send(client *http.Client, from, to, message string) error
}
