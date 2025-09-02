package utilities

import "strings"

func PhoneFomat(value string) string {
	str := value
	// Check if the first character is '0' and remove it if it is
	if len(str) > 0 && str[0] == '0' {
		str = str[1:]
	}

	// Remove spaces from the string
	result := strings.ReplaceAll(str, " ", "")

	return result
}
