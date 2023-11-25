package db

import (
	"database/sql/driver"
	"encoding/json"
	"time"
	"fmt"
)

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	IsEmailVerified   int       `json:"is_email_verified"`
	Role              string    `json:"role"`
}

func (u *User) TableName() string {
	return "users"
}


// T 是可以被 json 处理的数据
type JsonColumn [T any] struct {
	Val T 
	// 主要解决 NULL 之类的问题
	Valid bool
}

func (j *JsonColumn[T]) Value() (driver.Value, error) {
	// NULL，我没有数据
	if !j.Valid {
		return nil, nil
	}
	return json.Marshal(j.Val)
}


func (j *JsonColumn[T]) Scan(src any) error {

//	int64，不能被 JSON 处理
//	float64
//	bool
//	time.Time

//	[]byte
//	string
//  nil

	var bs []byte
	switch data := src.(type) {
	case []byte:
		// 考虑额外处理 []byte{}
		bs = data
	case string:
		// 考虑额外处理空字符串，或者提倡用户不应该这么做
		bs = []byte(data)
	case nil:
		// 说明数据库里面存的就是 NULL
		return nil
	default:
		return fmt.Errorf("JsonColumn.Scan 不支持 src 类型 %v", src)
	}
	
	err := json.Unmarshal(bs, &j.Val)
	if err == nil {
		// 代表有数据
		j.Valid = true
	}
	return err
}
