package models

// DTO für die Registrierung mit Validierung
type RegisterUserDTO struct{
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// DTO für Login, ohne Username
type LoginUserDTO struct{
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type User struct {
	UserID   int 
	Username string
	Email    string
	Password string
}

type UserRepository interface{
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}