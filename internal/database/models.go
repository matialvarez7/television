// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Codetag   string
}
