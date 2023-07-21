package entity

type NoteCategory struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
	Status   int    `json:"status"`
}
