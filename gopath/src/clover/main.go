package main

import (
        "fmt"
        "context"
        "encoding/json"
        "github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) {
        for _, record := range snsEvent.Records {
                  snsRecord := record.SNS
                  fmt.Printf("Subject = %s \n", snsRecord.Subject)
                  fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
        }
}

func main() {
        lambda.Start(handler)
}