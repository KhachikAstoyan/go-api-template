// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID           int32
	Username     string
	Email        string
	PasswordHash string
	IsActive     pgtype.Bool
	Role         pgtype.Text
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
	LastLogin    pgtype.Timestamptz
}
