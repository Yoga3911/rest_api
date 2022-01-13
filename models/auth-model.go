package models

type Login struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" validate:"required,email,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type Register struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
	GenderID int16  `json:"gender_id" validate:"required"`
}
