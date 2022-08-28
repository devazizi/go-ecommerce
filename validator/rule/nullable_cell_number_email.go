package rule

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

func RequiredEmailOrPassword(anotherField any) validation.RuleFunc {
	return func(value interface{}) error {

		val := value.(string)
		anotherFieldString := anotherField.(string)

		if val == "" && anotherFieldString == "" {
			return errors.New("can not null email or cell number with together")
		}

		return nil
	}
}
