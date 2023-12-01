package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	// 昵称
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	// 姓名
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
		return fmt.Errorf("必须只包含小写字母、数字或下划线")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("必须只包含字母或空格")
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
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("不是有效的电子邮件地址")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("必须是正整数")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
