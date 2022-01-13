package models

type UpdateUser struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"email,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
	GenderID int16  `json:"gender_id" validate:"required"`
}
