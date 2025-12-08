package model

import (
	"time"

	"github.com/google/uuid"
)

type Flow struct {
	ID        uuid.UUID `json:"id"; gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `json:"name"; gorm:"type:varchar(255)"`
	Status    string    `json:"status"; gorm:"type:varchar(50);default:'pending'"`
	CreatedAt time.Time `json:"created_at";`
}
