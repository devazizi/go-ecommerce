package rule

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
)

func PasswordAndPasswordConfirmationEqual(passwordConfirmation string) validation.RuleFunc {
	return func(value interface{}) error {
		password := value.(string)

		if passwordConfirmation != password {
			return fmt.Errorf("password and confirimation not equal")
		}

		return nil
	}
}
