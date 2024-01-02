package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	DateTime      time.Time `json:"date_time" binding:"required"`
	Location      string    `json:"location" binding:"required"`
	UserId        uuid.UUID `gorm:"type:uuid;" json:"user_id"`
	User          User      `json:"-"`
	Registrations []Registration
}

func (e Event) Save(db *gorm.DB) (Event, error) {
	//later: add it to a database
	result := db.Create(&e)
	if result.Error != nil {
		return Event{}, result.Error
	}

	return e, nil
}

func UpdateEvent(db *gorm.DB, event Event) (Event, error) {
	result := db.Save(&event)
	if result.Error != nil {
		return Event{}, result.Error
	}

	return event, nil
}

func GetEventById(id uuid.UUID, db *gorm.DB) (Event, error) {
	var event Event
	result := db.Preload("Registrations").First(&event, id)
	if result.Error != nil {
		return Event{}, result.Error
	}

	return event, nil
}

func DeleteEvent(event Event, db *gorm.DB) (bool, error) {
	result := db.Delete(&event, event.ID)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func CheckForValidEvent(eventId uuid.UUID, db *gorm.DB) error {
	var event Event
	result := db.Where("id=?", eventId).First(&event)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllEvents(events *[]Event, db *gorm.DB) *gorm.DB {
	return db.Find(events)
}
