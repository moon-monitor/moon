package validate

import (
	"regexp"

	"github.com/moon-monitor/moon/pkg/merr"
)

// TextIsNull 判断字符串是否为空
func TextIsNull(text string) bool {
	return len(text) == 0
}

func CheckEmail(email string) error {
	match := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	if TextIsNull(email) || !match.MatchString(email) {
		return merr.ErrorParamsError("The email format is incorrect.").WithMetadata(map[string]string{
			"email": email,
		})
	}
	return nil
}
