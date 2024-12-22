package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

func IsValidPhoneNumber(phoneNumber string) (string, error) {
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")

	num, err := phonenumbers.Parse(phoneNumber, "TR")
	if err != nil {
		return "", fmt.Errorf("Error parsing phone number: %w", err)
	}

	if len(phoneNumber) < 12 || len(phoneNumber) > 13 {
		return "", fmt.Errorf("Phone number is less than 12 or greater than 13 digits: %s", phoneNumber)
	}

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !re.MatchString(phoneNumber) {
		return "", fmt.Errorf("Phone number is not valid: %s", phoneNumber)
	}

	return phonenumbers.Format(num, phonenumbers.NATIONAL), nil
}
