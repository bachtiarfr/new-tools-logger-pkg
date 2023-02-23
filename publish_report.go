package publish_report

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type dataConfig struct {
	ProjectID string
	TopicID	  string
	CredentialFileJson []byte
}

func PublishReport(ctx context.Context, data map[string]interface{}, conf []byte) error {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshal :", err)
	}

	fmt.Printf("configurasi byte :", conf)

	dataConfig := json.RawMessage{}

	json.Unmarshal(conf, dataConfig)

	fmt.Printf("unmarshal :", dataConfig)

	projectID := "bachtiar-development"
	topicID := "dev-logger-topic"

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsJSON(conf))
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