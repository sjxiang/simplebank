package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	// 昵称（小写字母、数字、下划线）
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	// 姓名（大小写字母、）
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

// 辅助
func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("字符长度必须在 %d-%d 之间", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	// 先验证长度
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	// 再校验是否合规
	if !isValidUsername(value) {
		return fmt.Errorf("只能有小写字母、数字或下划线组成")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("只能有字母或空格组成")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}

	// 验证邮箱地址
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("不是有效的 email")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("只能是正整数")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
