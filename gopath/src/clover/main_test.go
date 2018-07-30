package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	msgMock := `{
		"merchantId": "DNT6PRJ6HWCR8",
		"body": [
			"{\"appId\":\"Q0WJXNDEMCDMG\",\"merchants\":{\"DNT6PRJ6HWCR8\":[{\"objectId\":\"A:Q0WJXNDEMCDMG\",\"type\":\"CREATE\",\"ts\":1532899088803}]}}"
		]
	 }`
	snsEventMock := events.SNSEvent{
		Records: []events.SNSEventRecord{
			{
				SNS: events.SNSEntity{
					Message: msgMock,
				},
			},
		},
	}
	ctx := context.Background()
	// Add keys to your liking, then:
	lc := new(lambdacontext.LambdaContext)
	ctx = lambdacontext.NewContext(ctx, lc)
	err := handler(ctx, snsEventMock)
	assert.NoError(t, err)
}
