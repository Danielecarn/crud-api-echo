package models

import "github.com/go-playground/validator/v10"

type (
	User struct {
		Id         int    `json:"id" `
		Username   string `json:"username" validate:"required"`
		FirstName  string `json:"firstname" validate:"required"`
		LastName   string `json:"lastname" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required"`
		Phone      string `json:"phone" validate:"required"`
		UserStatus int    `json:"userstatus" validate:"gte=0,lte=2"`
	}
)

type Users []User

func Validacao(user User) bool {
	validate := validator.New()
	errV := validate.Struct(user)
	if errV != nil {
		return false
	}
	return true
}
