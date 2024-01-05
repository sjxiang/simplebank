package val

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateEmailOtherWay(value string) error {
	email := value 
	if err := validator.New().Var(email, "required,email"); err != nil {
		return fmt.Errorf("不是有效的 email：%w", err)
	}

	allowed := false
	for _, domain := range emailDomainWhitelist {
		if strings.HasSuffix(email, "@"+domain) {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Errorf("邮箱地址的域名不在白名单中")
	}

	return nil 
}

// 白名单
var emailDomainWhitelist = []string{
	"gmail.com",
	"163.com",
	"126.com",
	"qq.com",
	"outlook.com",
	"hotmail.com",
	"yahoo.com",
	"foxmail.com",
}

