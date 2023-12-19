package db

import (
	"context"
)


const createUser = `
INSERT INTO user 
	(username, hashed_password, full_name, email, is_email_verified, role) 
VALUES 
	(?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	Username        string `json:"username"`
	HashedPassword  string `json:"hashed_password"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"is_email_verified"`
	Role            string `json:"role"`
}

func (data *Data) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	result, err := data.db.ExecContext(ctx, createUser,
		arg.Username, 
		arg.HashedPassword,
		arg.FullName, 
		arg.Email,
		arg.IsEmailVerified,
		arg.Role,
	)
	if err != nil {
		return 0, err
	}
	
	// 最后插入的 id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return id, nil
}


const getUser = `
SELECT 
	username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified, role
FROM 
	user 
WHERE 
	username = ?
LIMIT 1
`

func (data *Data) GetUser(ctx context.Context, username string) (User, error) {
	
	row := data.db.QueryRowContext(ctx, getUser, username)
	if row.Err() != nil {
		return User{}, row.Err()
	}
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
		&i.Role,
	)

	return i, err
}


const listUser = `
SELECT 
	username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified, role
FROM 
	user 
ORDER BY 
	id
LIMIT 
	?
OFFSET 
	?
`


type ListUsersParams struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (data *Data) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := data.db.QueryContext(ctx, listUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err 
	}
	defer rows.Close()

	items := []User{}
	for rows.Next() {
		var i User
		
		// 常见错误：1、漏了指针 2. 结果集顺序颠倒（和 select 字段 顺序对上）
		if err := rows.Scan(
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.IsEmailVerified,
			&i.Role,
		); err != nil {
			return nil, err 
		}
		items = append(items, i)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
