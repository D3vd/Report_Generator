package main

import (
	"encoding/csv"
	"os"
)

// WriteFlightsToCSV : Writes the array of flights to a CSV file
func WriteFlightsToCSV(flights []Flight) (ok bool) {

	csvFile, err := os.Create("./output/report.csv")

	if err != nil {
		return false
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	defer writer.Flush()

	headers := []string{
		"TimeStamp",
		"Carrier",
		"OriginCityName",
		"OriginCountry",
		"DestCityName",
		"DestCountry",
		"FlightTimeMin",
		"AvgTicketPrice",
		"Cancelled",
		"FlightDelayType",
		"FlightDelay",
		"DistanceKilometers",
	}

	writeErr := writer.Write(headers)

	if writeErr != nil {
		return false
	}

	for _, flight := range flights {
		value := []string{
			flight.TimeStamp,
			flight.Carrier,
			flight.OriginCityName,
			flight.OriginCountry,
			flight.DestCityName,
			flight.DestCountry,
			flight.FlightTimeMin,
			flight.AvgTicketPrice,
			flight.Cancelled,
			flight.FlightDelayType,
			flight.FlightDelay,
			flight.DistanceKilometers,
		}

		err := writer.Write(value)

		if err != nil {
			return false
		}
	}

	return true
}
