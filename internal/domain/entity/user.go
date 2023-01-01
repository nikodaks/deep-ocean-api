package entity

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	AuthCode int    `json:"authCode" db:"auth_code"`
}
