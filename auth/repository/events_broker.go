package repository

import (
	"cloud.google.com/go/pubsub"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// EventsBroker is implementation for plusing events that will be
// published to pubsub
type EventsBroker struct {
	PubSub *pubsub.Client
}

// PublishUserUpdated sends user to "user-updated" topic of pubsub
func (e *EventsBroker) PublishUserUpdated(u *model.User) error {
	panic("Not implemented")
}
