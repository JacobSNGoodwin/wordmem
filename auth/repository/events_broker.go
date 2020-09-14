package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jacobsngoodwin/wordmem/auth/rerrors"

	"cloud.google.com/go/pubsub"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// EventsBroker is implementation for plusing events that will be
// published to pubsub
type EventsBroker struct {
	PubSub *pubsub.Client
}

const topic = "user-updates"

// PublishUserUpdated sends user to "user-updated" topic of pubsub
func (e *EventsBroker) PublishUserUpdated(u *model.User, isNewUser bool) error {
	t := e.PubSub.Topic(topic)

	serializedUser, err := json.Marshal(u)

	if err != nil {
		log.Printf("Problem serializing user in PublishUserUpdated: %v\n", err)
		return rerrors.NewInternal()
	}

	ctx := context.Background()

	var eventType string
	if isNewUser {
		eventType = "user-created"
	} else {
		eventType = "user-updated"
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: serializedUser,
		Attributes: map[string]string{
			"type": eventType,
		},
	})

	_, err = result.Get(ctx)

	if err != nil {
		log.Printf("Failure getting id of published result: %v\n", err)
	}

	return nil
}
