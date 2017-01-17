// Package http provides parameter validation for http request (for now?).
package http

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// ValidateNumber validates paramValue as integer.
// It takes string parameter (paramValue) and returns its integer value if successful.
func ValidateNumber(paramValue string) (int64, error) {
	intValue, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil {
		message := fmt.Sprintf("Invalid parameter value: %s", paramValue)
		return 0, errors.New(message)
	}
	return intValue, nil
}

// ValidatePositiveNumber validates paramValue as positive integer (zero allowed if allowZero true).
// It takes string parameter (paramValue) and returns its integer value if successful.
func ValidatePositiveNumber(paramValue string, allowZero bool) (int64, error) {
	intValue, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil || (intValue == 0 && !allowZero) || intValue < 0 {
		message := fmt.Sprintf("Invalid parameter value: %s", paramValue)
		return 0, errors.New(message)
	}
	return intValue, nil
}

// ValidateNegativeNumber validates paramValue as negative integer.
// It takes string parameter (paramValue) and returns its interger value if successful.
func ValidateNegativeNumber(paramValue string) (int64, error) {
	intValue, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil || intValue >= 0 {
		message := fmt.Sprintf("Invalid parameter value: %s", paramValue)
		return 0, errors.New(message)
	}
	return intValue, nil
}

// ValidateDateFormat validates parameter of type date based on layout parameter.
// The layout parameter is based on standard library layout reference "2 January 2006".
// It takes string parameter and returns time.Time object (with hour, minute, second set to zero)
// when no error occurs.
func ValidateDateFormat(paramValue string, layout string) (time.Time, error) {
	t, err := time.Parse(layout, paramValue)
	if err != nil {
		message := fmt.Sprintf("Invalid date format: input[%s] layout[%s]", paramValue, layout)
		return time.Time{}, errors.New(message)
	}
	return t, nil
}
