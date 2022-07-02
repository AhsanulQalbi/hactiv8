package models

type SocialMedia struct {
	ID               uint   `gorm:"primary key" json:"id,omitempty" `
	Name             string `json:"name,omitempty"`
	Social_Media_Url string `json:"social_media_url,omitempty"`
	UserID           uint   `json:"user_id,omitempty"`
	Created_at       string `json:"created_at,omitempty"`
	Updated_at       string `json:"updated_at"`
	User             *User
}
