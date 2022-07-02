package models

type Comment struct {
	ID         uint   `gorm:"primary key" json:"id"`
	UserID     uint   `json:"user_id"`
	PhotoID    uint   `json:"photo_id"`
	Message    string `json:"message"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	User       *User
	Photo      *Photo
}
