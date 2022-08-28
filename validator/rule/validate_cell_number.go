package rule

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func ValidateCellNumber() validation.RuleFunc {
	return func(value interface{}) error {

		cellNumber := value.(string)

		isMatch, _ := regexp.MatchString("^09(0[1-5]|[1 3]\\d|2[0-2]|98)\\d{7}$", cellNumber)

		if !isMatch {
			return errors.New("cell Number formant incorrect")
		}

		return nil
	}
}
