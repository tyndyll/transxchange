package nptg

type Route struct {
	ID                    string `xml:"id,attr"`
	PrivateCode           string `xml:"PrivateCode"`
	Description           string `xml:"Description"`
	RouteSectionReference string `xml:"RouteSectionRef"`
}
