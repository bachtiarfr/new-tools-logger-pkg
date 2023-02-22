package publish_report

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func PublishReport(ctx context.Context, data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshal :", err)
	}

	projectID := "bachtiar-development"
	credentialFilePath := "bachtiar-development-73ca13e5c16e.json"
	topicID := "dev-logger-topic"

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(d),
	})
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %v", err)
	}

	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}