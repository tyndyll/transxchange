package nptg

type JourneyPatternSection struct {
	TimingLinks []*JouneyTimingLink `xml:"JourneyTimingLink"`
}

type JourneyTimingLink struct {
	ID                 string           `xml:"id,attr"`
	From               *JourneySequence `xml:"From"`
	To                 *JourneySequence `xml:"To"`
	RouteLineReference string           `xml:"RL_90-1-_-r12-1-O-1-1"`
	RunTime            string           `xml:"RunTime"`
}

type JourneySequence struct {
	Number             int    `xml:"SequenceNumber,attr"`
	Activity           string `xml:"Activity"`
	StopPointReference string `xml:"StopPointRef"`
	TimingStatus       string `xml:"TimingStatus"`
}
