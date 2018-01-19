package nptg

type RouteSection struct {
	ID    string       `xml:"id,attr"`
	Links []*RouteLink `xml:"RouteLink"`
}

type RouteLink struct {
	ID        string `xml:"id,attr"`
	From      string `xml:"From>StopPointRef"`
	To        string `xml:"To>StopPointRef"`
	Direction string `xml:"Direction"`
}
