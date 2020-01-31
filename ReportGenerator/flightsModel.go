package main

// Flights : Collection of Flight
type Flights struct {
	Flights []Flight `json:"Flights"`
}

// Flight : Format that needs to be written to CSV
type Flight struct {
	TimeStamp          string `json:"timestamp"`
	Carrier            string `json:"Carrier"`
	OriginCityName     string `json:"OriginCityName"`
	OriginCountry      string `json:"OriginCountry"`
	DestCityName       string `json:"DestCityName"`
	DestCountry        string `json:"DestCountry"`
	FlightTimeMin      string `json:"FlightTimeMin"`
	AvgTicketPrice     string `json:"AvgTicketPrice"`
	Cancelled          string `json:"Cancelled"`
	FlightDelayType    string `json:"FlightDelayType"`
	FlightDelay        string `json:"FlightDelay"`
	DistanceKilometers string `json:"DistanceKilometers"`
}
