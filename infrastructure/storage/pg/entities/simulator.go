package entities

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)


type Simulator struct {
	ID            uuid.UUID      `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Date          string         `json:"date"`
	Landing       int            `json:"landing"`
	Role          string         `json:"role"`
	Qualification string         `json:"qualification"`
	Session       float32        `json:"session"`
	CreatedAt     time.Time      `json:"created"`
	UpdatedAt     time.Time      `json:"updated"`
	DeletedAt     gorm.DeletedAt `json:"deleted"`
}


func (s *Simulator) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()

	if &s.ID != nil {
		return
	}
	err = errors.New("failure to generate uuid...")
	return err
}
