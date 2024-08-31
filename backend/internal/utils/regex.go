package utils

import "regexp"

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString
	hasSpecial := regexp.MustCompile(`[^A-Za-z0-9]`).MatchString

	return hasUpper(password) && hasLower(password) && hasDigit(password) && hasSpecial(password)
}
