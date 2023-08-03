package entity

type User struct {
	Id           uint   `json:"id"`
	UserName     string `json:"user_name"`
	ChatId       int64  `json:"chat_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}
