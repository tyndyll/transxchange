package nptg

type JourneyPattern struct {
	ID                             string `xml:"id,attr"`
	RouteReference                 string `xml:"RouteRef"`
	JourneyPatternSectionReference string `xml:"JourneyPatternSectionRefs"`
}
