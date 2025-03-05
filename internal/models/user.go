package auth_models

import "context"

type User struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	RoleID      uint   `json:"roleId"`
	Balance     int    `json:"balance"`
	CreatedAt   string `json:"createdAt"`
}

type UserProfile struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Balance int    `json:"balance"`
}

type UserRequest struct {
	Email       string   `json:"email"`
	Password    Password `json:"password"`
	PhoneNumber string   `json:"phoneNumber"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type AddUser struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, userID int) (User, error)
	GetProfile(c context.Context, userID int) (UserProfile, error)
	CreateUser(c context.Context, user UserRequest) (int, error)
	EditPersonalData(c context.Context, userID int, name string, email string, phoneNumber string, Birthday string) error
	AddUser(c context.Context, name string, balance int) (int, error) 
}
