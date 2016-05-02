package smsexpress

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Smsgh provider implements smsexpress.Provider interface
type Smsgh struct {
	ID, Secret string
}

// Send SMS
func (s Smsgh) Send(client *http.Client, from, to, message string) error {
	v := make(url.Values)
	v.Set("From", from)
	v.Set("To", to)
	v.Set("Content", message)
	v.Set("ClientId", s.ID)
	v.Set("ClientSecret", s.Secret)
	response, err := client.Get("https://api.smsgh.com/v3/messages/send?" + v.Encode())
	if err != nil {
		return err
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("smsexpress: Smsgh failed with HTTP status %s\n%s", response.Status, string(contents))
}
