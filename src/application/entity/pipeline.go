package entity

import (
	"github.com/leryn1122/kreutzer/v2/infra/id"
	wh "github.com/leryn1122/kreutzer/v2/infra/webhook"
)

type Pipeline struct {
	ID      id.Identifier
	Name    string
	Enabled bool
}

// Webhook Received from adapter
type Webhook struct {
	ID        id.BigInt
	HookID    string
	Channel   string
	ChannelID int
	Enabled   bool
	URL       string
	Secret    string
	Payload   wh.Webhook
}

func (wh *Webhook) Verify() error {
	return wh.Payload.Verify([]byte(wh.Secret))
}
