package model

import "time"

type NotificationEvent struct {
	UserID    int32     `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type EmailEvent struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}
