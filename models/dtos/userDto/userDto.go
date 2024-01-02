package userDto

import "github.com/google/uuid"

type UserDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginDto struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}
