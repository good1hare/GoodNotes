package entity

type NoteCategory struct {
	Id       int    `json:"id"`
	ChatId   int    `json:"chat_id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
	Status   int    `json:"status"`
}
