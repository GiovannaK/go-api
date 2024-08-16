package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50,containsany=!@#$%^&*()_+"`
	Age      int8   `json:"age" binding:"required,min=1,max=120"`
}
