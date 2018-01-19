package nptg

type VehicleJourney struct {
	PrivateCode string `xml:"PrivateCode"`
	// TODO OperatingProfile

	Code                    string `xml:"VehicleJourneyCode"`
	ServiceReference        string `xml:"ServiceRef"`
	LineReference           string `xml:"LineRef"`
	JourneyPatternReference string `xml:"JourneyPatternRef"`
	DepartureTime           string `xml:"DepartureTime"`
}
