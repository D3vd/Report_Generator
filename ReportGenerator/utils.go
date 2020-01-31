package main

import "time"
import "encoding/json"
import "gopkg.in/olivere/elastic.v3"

// ConvertTimeLayoutToISO : Convert UI form Layout to ISO Default
func ConvertTimeLayoutToISO(date string) (ISO time.Time, ok bool) {
	const UIFormat string = "01/02/2006"

	t, err := time.Parse(UIFormat, date)

	if err != nil {
		return time.Time{}, false
	}

	return t, true
}

// ParseESResultToModel : Convert the ES Hits Result to Flights Model
func ParseESResultToModel(hits []*elastic.SearchHit) (fs []Flight, ok bool) {
	var flights []Flight

	for _, hit := range hits {
		var flight Flight

		err := json.Unmarshal(*hit.Source, &flight)

		if err != nil {
			return flights, false
		}

		flights = append(flights, flight)
	}

	return flights, true
}
