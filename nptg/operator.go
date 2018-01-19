package nptg

type Operator struct {
	ID          string `xml:"id,attr"`
	Code        string `xml:"OperatorCode"`
	Name        string `xml:"OperatorNameOnLicence"`
	TradingName string `xml:"TradingName"`
}
