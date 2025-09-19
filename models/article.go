package models

import "time"

type Article struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"size:255; not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Author    string    `gorm:"size:255" json:"author"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
