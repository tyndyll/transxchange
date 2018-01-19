package nptg

type Service struct {
	Code            string  `xml:"ServiceCode"`
	PrivateCode     string  `xml:"PrivateCode"`
	Lines           []*Line `xml:"Lines"`
	OperatingPeriod *Period `xml:"OperatingPeriod"`
	// TODO: OperatingProfile
	OperatorReference string `xml:"RegisteredOperatorRef"`
	// TODO: StopRequirements
	// TODO: Are these a String type?
	Mode string `xml:"Mode"`
	// Better name?
	StandardService *ServiceProfile `xml:"StandardService"`
}

type Line struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"LineName"`
}

type Period struct {
	StartDate time.Date `xml:"StartDate"`
	EndDate   time.Date `xml:"EndDate"`
}

type ServiceProfile struct {
	Origin          string            `xml:"Origin"`
	Destination     string            `xml:"Destination"`
	JourneyPatterns []*JourneyPattern `xml:"JourneyPattern"`
}
