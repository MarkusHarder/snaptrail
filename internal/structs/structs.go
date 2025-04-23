package structs

import (
	"time"
)

type Hello struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Text      string    `gorm:"type:text;not null" json:"text"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
