package model

type User struct {
	ID    string `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name" validate:"required"`
	Email string `json:"email" xml:"email" validate:"required,email"`
}
