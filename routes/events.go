package routes

import (
	"events-api/helpers/ReturnHelper"
	"events-api/helpers/db"
	eventModel "events-api/models"
	"events-api/models/dtos/eventDto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateEvent(context *gin.Context) {

	var eventDTO eventDto.EventDto
	userid, _ := context.Get("user_id")
	err := context.ShouldBindJSON(&eventDTO)

	if err != nil {
		fmt.Println(err)

		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var event eventModel.Event

	event.ID = uuid.New()
	event.Name = eventDTO.Name
	event.Location = eventDTO.Location
	event.DateTime = eventDTO.DateTime
	event.Description = eventDTO.Description

	userId := uuid.MustParse(userid.(string))

	event.UserId = userId

	savedEvent, err := eventModel.Event.Save(event, db.DB)

	if err != nil {
		panic(err)
		return
	}

	var eventResponseDto eventDto.EventResponseDto

	eventResponseDto.Name = savedEvent.Name
	eventResponseDto.UserId = savedEvent.UserId
	eventResponseDto.DateTime = savedEvent.DateTime
	eventResponseDto.Description = savedEvent.Description
	eventResponseDto.Location = savedEvent.Location

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   eventResponseDto,
	})
}

func GetEventById(context *gin.Context) {

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

	userid, _ := context.Get("user_id")

	userId := uuid.MustParse(userid.(string))

	if boole == false {
		panic("Something went wrong")
		return
	}

	event, err := eventModel.GetEventById(uuid.MustParse(id), db.DB)

	if event.UserId != userId {
		panic("not authenticated")
		return
	}

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

func DeleteEvent(context *gin.Context) {
	id, success := context.Params.Get("id")

	userid, _ := context.Get("user_id")

	userId := uuid.MustParse(userid.(string))
	if success == false {
		ReturnHelper.Return(context, http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	event, err := eventModel.GetEventById(uuid.MustParse(id), db.DB)

	if event.UserId != userId {
		panic("not authenticated")
		return
	}

	var eventid = event.ID

	if err != nil {
		ReturnHelper.Return(context, http.StatusBadRequest, gin.H{
			"error": "Event not found",
		})
		return
	}

	result, err := eventModel.DeleteEvent(event, db.DB)

	if err != nil && result != true {
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": "Deleted Event:" + eventid.String(),
	})

}
