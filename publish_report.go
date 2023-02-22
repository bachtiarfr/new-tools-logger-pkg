package publish_report

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func PublishReport(ctx context.Context, data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshal :", err)
	}

	jsonFile, errJsonFile := os.Open("bachtiar-development-73ca13e5c16e.json")
    if errJsonFile != nil {
        fmt.Println("error when open json file :", errJsonFile)
    }
	
	byteJsonFile, _ := ioutil.ReadAll(jsonFile)
	projectID := "bachtiar-development"
	topicID := "dev-logger-topic"

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsJSON(byteJsonFile))
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