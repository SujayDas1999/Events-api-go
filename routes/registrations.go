package routes

import (
	"events-api/helpers/ReturnHelper"
	"events-api/helpers/db"
	"events-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateRegistration(context *gin.Context) {
	id, _ := context.Params.Get("id")
	eventId := uuid.MustParse(id)

	//var event models.Event
	_, err := models.GetEventById(eventId, db.DB)

	if err != nil {
		panic(err)
		return
	}

	userid, _ := context.Get("user_id")

	userId := uuid.MustParse(userid.(string))

	var registration models.Registration

	registration.ID = uuid.New()
	registration.IsCancelled = false
	registration.Event_Id = eventId
	registration.User_Id = userId

	registration.Save(db.DB)

	var event models.Event

	db.DB.Preload("Registrations").First(&event, eventId)

	ReturnHelper.ReturnSuccessWithData(context, event)

}

func CancelRegistration(context *gin.Context) {
	id, _ := context.Params.Get("id")
	registrationId := uuid.MustParse(id)

	userid, _ := context.Get("user_id")

	userId := uuid.MustParse(userid.(string))

	err := models.Cancel(registrationId, userId, db.DB)

	if err != nil {
		panic("Not authorised")
		return
	}

	ReturnHelper.ReturnSuccess(context, "success")
	return

}
