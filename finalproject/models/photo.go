package models

type Photo struct {
	ID         uint   `gorm:"primary key" json:"id"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Photo_url  string `json:"photo_url"`
	UserID     uint   `json:"user_id"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	User       *User  `json:"User,omitempty"`
}
