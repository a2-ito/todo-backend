package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Owner       uuid.UUID
	Title       string
	Status      string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
