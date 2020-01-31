package main

import "time"

// ConvertTimeLayoutToISO : Convert UI form Layout to ISO Default
func ConvertTimeLayoutToISO(date string) (ISO time.Time, ok bool) {
	const UIFormat string = "01/02/2006"

	t, err := time.Parse(UIFormat, date)

	if err != nil {
		return time.Time{}, false
	}

	return t, true
}
