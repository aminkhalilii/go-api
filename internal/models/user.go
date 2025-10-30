package models

type User struct {
	ID       int    `json:"id" `
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
}
