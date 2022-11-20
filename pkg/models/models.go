package models

import "time"

type CineRow struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	AuthorTitle string    `json:"author_title"`
	Album       string    `json:"album"`
	Duration    string    `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RandomTable struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
