package activities

import (
	"os"

	models "github.com/shreyasprasad/reps/app/server/models"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSSendRep struct {
	Number string
}

func (s SMSSendRep) Send(rep models.Rep) error {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(s.Number)
	params.SetFrom(from)
	params.SetBody(rep.Text)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
