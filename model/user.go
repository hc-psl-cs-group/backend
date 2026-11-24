package model

type User struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	Uid        string      `json:"-"` // UserID as specified by OIDC OP
	Email      string      `json:"email"`
	Username   string      `json:"username"`
}
