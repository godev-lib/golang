package pusher_custom

import (
	"github.com/godev-lib/golang/config"
	"github.com/pusher/pusher-http-go/v5"
)

func NewPusherClient(config *config.Config) *pusher.Client {
	client := &pusher.Client{
		AppID:  config.Pusher.AppID,
		Key:    config.Pusher.Key,
		Secret: config.Pusher.Secret,
		Host:   config.Pusher.Host,
		Secure: config.Pusher.Secure,
	}
	return client
}
