package smsexpress

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	Register(Smsgh{
		ID:     os.Getenv("SMSGH_ClIENT_ID"),
		Secret: os.Getenv("SMSGH_CLIENT_SECRET"),
	})

	Register(NasaraMobile{
		APIKey: os.Getenv("NASARA_MOBILE_API_KEY"),
	})

	ok, _ := Send(nil, "smsexpress", "0246662003", "This is not a test!")
	assert.True(t, ok)
	// assert.NoError(t, err)
}
