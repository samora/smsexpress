package smsexpress

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Hubtel provider implements smsexpress.Provider interface
type Hubtel struct {
	ID, Secret string
}

// Send SMS
func (h Hubtel) Send(client *http.Client, from, to, message string) error {
	v := make(url.Values)
	v.Set("From", from)
	v.Set("To", to)
	v.Set("Content", message)
	v.Set("ClientId", h.ID)
	v.Set("ClientSecret", h.Secret)
	response, err := client.Get("https://api.hubtel.com/v1/messages/send?" + v.Encode())
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("smsexpress: Hubtel failed with HTTP status %s\n%s", response.Status, string(contents))
}
