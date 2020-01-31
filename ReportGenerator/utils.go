package main

import "time"

// UIFormat : Format that is generated from UI form
const UIFormat string = "02/01/2006"

// ISODefault : Format that elasticsearch accepts
const ISODefault string = "2006-01-02"

// ConvertTimeLayoutToISO : Convert UI form Layout to ISO Default
func ConvertTimeLayoutToISO(date string) (ISO string) {

	t, err := time.Parse(UIFormat, date)

	if err != nil {
		return "error"
	}

	return t.Format(ISODefault)
}
