package budgetalertsms

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func init() {
	functions.CloudEvent("BudgetAlertSms", budgetAlertSms)
}

type MessagePublishedData struct {
	Message PubSubMessage
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func budgetAlertSms(ctx context.Context, e event.Event) error {
	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}

	data := string(msg.Message.Data)
	if data == "" {
		log.Printf("No data found in Pub/Sub message")
	} else {
		log.Printf("Pub/Sub message: %s", data)
	}

	resp, err := sendSMS(data)
	if err != nil {
		log.Printf("Error sending SMS: %s", err.Error())
	} else {
		log.Printf("SMS sent successfully: %d", resp.Sid)
	}

	return nil
}

func sendSMS(message string) (*openapi.ApiV2010Message, error) {

	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
