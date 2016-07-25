package smsexpress

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// NasaraMobile implements Provider
type NasaraMobile struct {
	APIKey string
}

func (n NasaraMobile) Send(client *http.Client, from, to, message string) error {
	v := make(url.Values)
	v.Set("api_key", n.APIKey)
	v.Set("sender_id", from)
	v.Set("phone", to)
	v.Set("message", message)
	response, err := client.Get("http://sms.nasaramobile.com/api?" + v.Encode())
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("smsexpress: NasaraMobile failed with HTTP status %s\n%s", response.Status, string(contents))
}
