package mapper

import (
	eventModel "events-api/models"
)

func EventsMapper(event *eventModel.Event) *eventModel.Event {
	return &eventModel.Event{
		ID:          event.ID,
		Name:        event.Name,
		DateTime:    event.DateTime,
		Description: event.Description,
		Location:    event.Location,
		UserId:      event.UserId,
	}
}
