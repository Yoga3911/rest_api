package models

import "time"

type Login struct {
	Email    string `json:"email" validate:"required,email,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type Register struct {
	Name     string    `json:"name" validate:"required,min=3,max=50"`
	Email    string    `json:"email" validate:"required,email,min=6,max=50"`
	Password string    `json:"password" validate:"required,email,min=6,max=50"`
	GenderID int16     `json:"gender_id" validate:"required"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
