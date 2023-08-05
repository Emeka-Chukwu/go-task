package auths

import (
	"errors"

	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}
func IsContained(value string, list []string) error {
	err := fmt.Errorf("%s not supported", value)
	for _, element := range list {
		if strings.EqualFold(value, element) {
			err = nil
		}
	}
	return err
}

func ValidateNotEmptyString(value string) error {
	n := len(value)
	if n < 1 {
		return fmt.Errorf("cannot be empty characters")
	}

	return nil
}
func ValidateGender(value string) error {
	if strings.EqualFold(value, "male") {
		return nil
	} else if strings.EqualFold(value, "female") {
		return nil
	} else if strings.EqualFold(value, "others") {
		return nil
	}
	return errors.New("invalid gender")
}
func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
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
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}
func ValidateNumber(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}
func ValidateNotLessThanZeroNumber(value int64) error {
	if value < 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}

func ValidateTimeStamp(value time.Time) error {
	const layout = "2006-01-02 03:04:05.999"
	_, error := time.Parse(layout, value.String())
	return error
}

func ValidateTimeStampRange(start, end time.Time) error {
	const layout = "2006-01-02 03:04:05.999"
	_, err := time.Parse(layout, start.String())
	if err != nil {
		return err
	}
	_, err = time.Parse(layout, end.String())
	if err != nil {
		return err
	}

	if end.After(start) {
		return nil
	}
	return errors.New("start must be before end")
}

func ValidateStartNilEndNotNil(start, end *time.Time) error {
	if start == nil && end != nil {
		return errors.New("start_date must not be nil when end_date is not nil")
	}
	return nil
}

func ValidateUUID(id string) error {
	if id == "00000000-0000-0000-0000-000000000000" {
		return errors.New("field cannot be empty")
	}
	vale, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	fmt.Println(vale, id)
	return nil
}
