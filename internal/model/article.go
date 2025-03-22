package model

import (
	"time"

	"github.com/lib/pq"
)

type Article struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Author    string         `json:"author"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`
	CreatedAt time.Time
}
