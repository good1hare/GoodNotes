package entity

type Note struct {
	Id          int    `json:"id"`
	ChatId      int    `json:"chat_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int    `json:"category_id"`
	Status      int    `json:"status"`
}
