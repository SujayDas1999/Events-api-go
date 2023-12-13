package routes

import (
	"events-api/helpers/db"
	"events-api/models/dtos/eventDto"
	eventModel "events-api/models/event"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateEvent(context *gin.Context) {

	var event eventModel.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	event.UserId = 1

	savedEvent, err := eventModel.Event.Save(event, db.DB)

	if err != nil {
		panic(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   savedEvent,
	})
}

func GetEventById(context *gin.Context) {
	//var id int

	id, boole := context.Params.Get("id")

	if boole == false {
		panic("Something went wrong")
		return
	}

	event, err := eventModel.GetEventById(uuid.MustParse(id), db.DB)

	if err != nil {
		panic(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})

}

func GetEvents(context *gin.Context) {
	var events []eventModel.Event

	eventModel.GetAllEvents(&events, db.DB)
	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func UpdateEvent(context *gin.Context) {

	id, boole := context.Params.Get("id")

	if boole == false {
		panic("Something went wrong")
		return
	}

	event, err := eventModel.GetEventById(uuid.MustParse(id), db.DB)

	if err != nil {
		panic(err)
		return
	}

	var updatedEvent eventDto.EventDto

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		panic(err)
	}

	event.Description = updatedEvent.Description
	event.Location = updatedEvent.Location
	event.DateTime = updatedEvent.DateTime
	event.Name = updatedEvent.Name

	event, err = eventModel.UpdateEvent(db.DB, event)

	if err != nil {
		panic(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
