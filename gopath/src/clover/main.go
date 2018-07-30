package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
 {
    "merchantId": "DNT6PRJ6HWCR8",
    "body": [
        "{\"appId\":\"Q0WJXNDEMCDMG\",\"merchants\":{\"DNT6PRJ6HWCR8\":[{\"objectId\":\"A:Q0WJXNDEMCDMG\",\"type\":\"CREATE\",\"ts\":1532899088803}]}}"
    ]
 }
*/
type Message struct {
	MerchantID string   `json:"merchantId"`
	Body       []string `json:"body"`
}

type Body struct {
	AppID     string                      `json:"appId"`
	Merchants map[string]*json.RawMessage `json:"merchants"`
}

type Event struct {
	ObjectID  string `json:"objectId"`
	Type      string `json:"type"`
	Timestamp uint64 `json:"ts"`
}

func handler(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		fmt.Printf("Subject = %s \n", snsRecord.Subject)
		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		var msg Message
		err := json.Unmarshal([]byte(snsRecord.Message), &msg)
		if err != nil {
			fmt.Println("error:", err)
			return err
		}
		fmt.Printf("%+v \n", msg)

		var body Body
		err1 := json.Unmarshal([]byte(msg.Body[0]), &body)
		if err1 != nil {
			fmt.Println("error1", err1)
			return err1
		}
		fmt.Printf("body %+v \n", body)
		var events []Event
		err2 := json.Unmarshal(*body.Merchants[msg.MerchantID], &events)
		if err2 != nil {
			fmt.Println("error2", err2)
			return err2
		}
		fmt.Printf("events %+v \n", events[0])
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
