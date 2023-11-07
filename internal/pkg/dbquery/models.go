// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dbquery

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Campaign struct {
	ID          *uuid.UUID
	Name        string
	Description string
	AdserverID  *uuid.UUID
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type GooseDbVersion struct {
	ID        int32
	VersionID int64
	IsApplied bool
	Tstamp    pgtype.Timestamp
}

type GueJob struct {
	JobID      string
	Priority   int16
	RunAt      pgtype.Timestamptz
	JobType    string
	Args       []byte
	ErrorCount int32
	LastError  pgtype.Text
	Queue      string
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}
