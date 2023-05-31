package validate

import "errors"

func IsNumeric(s string) error {
	for _, v := range s {
		if v < '0' || v > '9' {
			return errors.New("string is not numeric")
		}
	}
	return nil
}
