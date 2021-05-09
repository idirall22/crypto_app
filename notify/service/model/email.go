package model

type RegisterUserConfirmationEmailParams struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}
