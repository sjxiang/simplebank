package db

import (
	"errors"
	"fmt"

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


// 抽离出 err code
func ErrorCode(err error) string {
	var pgErr *pgconn.PgError

	// 第二个参数是：错误类型（管他咋实现的 error 接口）变量的指针
	if errors.As(err, &pgErr) {
		fmt.Println(">>> ", pgErr.ConstraintName)
		return pgErr.Code
	}
	return ""
}
