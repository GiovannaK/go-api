package request

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%^&*"`
}