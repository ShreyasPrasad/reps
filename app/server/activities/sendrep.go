package activities

import models "github.com/shreyasprasad/reps/app/server/models"

type SendRep interface {
	Send(rep models.Rep) error
}
