package db

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ForeignKeyViolation = "23503"  // 违反外键约束
	UniqueViolation     = "23505"  // 违反唯一约束
)

var ErrRecordNotFound = pgx.ErrNoRows  

var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}


// error -> code 
func ErrorCode(err error) string {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
