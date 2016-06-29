package smsexpress

import (
	"errors"
	"net/http"
)

var providers = []Provider{}

// Register adds a provider
func Register(p Provider) {
	providers = append(providers, p)
}

// Send SMS messages via registered providers.
// Executes in order of registration.
// Returns on first successful send.
func Send(client *http.Client, from, to, message string) (bool, error) {
	if len(providers) == 0 {
		panic(errors.New("smsexpress: No provider. Register provider(s)"))
	}
	if client == nil {
		client = http.DefaultClient
	}
	errs := Errors{}
	for _, p := range providers {
		err := p.Send(client, from, to, message)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if len(errs) > 0 {
			err = errs
		}
		return true, err
	}
	return false, errs
}
