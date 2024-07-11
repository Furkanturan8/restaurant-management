package models

import "time"

type Note struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Text      string    `json:"text"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	NoteID    string    `json:"note_id" gorm:"unique"`
}
