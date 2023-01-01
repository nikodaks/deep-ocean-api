package dto

type SignUpInput struct {
	Username string `json:"username" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
}
