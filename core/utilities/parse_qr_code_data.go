package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseQRCodeData(input string) (contractID, customerTypeID, dateTime string, err error) {
	// Split the string using "-" as a delimiter
	parts := strings.Split(input, "-")
	if len(parts) < 5 {
		return "", "", "", fmt.Errorf("invalid input format")
	}

	// Assign the components
	contractID = parts[0]
	customerTypeID = parts[1]
	dateTime = strings.Join(parts[2:], "-") // Join the rest as date-time

	return contractID, customerTypeID, dateTime, nil
}

func ParseQRCodeDataUint(input string) (contractID uint, customerTypeID uint, dateTime string, err error) {
	// Split the string using "-" as a delimiter
	parts := strings.Split(input, "-")
	if len(parts) < 5 {
		return 0, 0, "", fmt.Errorf("invalid input format")
	}

	// Convert contractID to uint
	contractIDInt, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid contract ID: %w", err)
	}
	contractID = uint(contractIDInt)

	// Convert customerTypeID to uint
	customerTypeIDInt, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid customer type ID: %w", err)
	}
	customerTypeID = uint(customerTypeIDInt)

	// Assign the date-time part
	dateTime = strings.Join(parts[2:], "-") // Join the rest as date-time

	return contractID, customerTypeID, dateTime, nil
}
