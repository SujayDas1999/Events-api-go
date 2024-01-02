package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Registration struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Event_Id    uuid.UUID `gorm:"type:uuid" json:"event_Id"`
	User_Id     uuid.UUID `gorm:"type:uuid" json:"user_Id"`
	Event       Event     `json:"-"`
	User        User      `json:"-"`
	IsCancelled bool      `json:"is_cancelled"`
}

func (r Registration) Save(db *gorm.DB) {
	db.Create(&r)
}

func Cancel(id uuid.UUID, userId uuid.UUID, db *gorm.DB) error {

	var registration Registration
	db.Where("id=?", id).First(&registration)

	if userId != registration.User_Id {
		return errors.New("not authorised")
	}

	registration.IsCancelled = true

	db.Save(&registration)

	return nil
}
