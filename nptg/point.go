package nptg

import (
	"time"
)

// StopPoint describes an access point to public transport
type StopPoint struct {
	// Timestamp of StopPoint creation date and time
	Creation time.Time `xml:"CreationDateTime,attr"`

	// Unique NaPTAN system identifier of StopPoint. Codes are unique
	// within the NaPTAN database for Great Britain. AtcoCode instances
	// normally have the form a0b, where 'a' is the three digit
	// AtcoAreaCode (Note that some additional values are used, for example
	// ‘910 Network Rail’), 0 is fixed, and b is an arbitrary unique
	// alphanumeric code of up to eight characters.
	AtcoCode string `xml:"AtcoCode"`

	// Groups together alternative unique identifiers of a StopPoint.

	PrivateCode string `xml:"PrivateCode"`

	CommonName string `xml:"Descriptor>CommonName"`

	Place *Place `xml:"Place"`

	Classification *StopClassification `xml:"StopClassification"`

	AdministrativeAreaReference string `xml:"AdministrativeAreaRef"`

	Notes string `xml:"Notes"`
}

type Place struct {
	LocalityReference string    `xml:"NptgLocalityRef"`
	Suburb            string    `xml:"Suburb"`
	Location          *GeoPoint `xml:"Location"`
}

type GeoPoint struct {
	Longitude float32 `xml:"Longitude"`
	Latitude  float32 `xml:"Latitude"`
}

type StopClassification struct {
	Type string `xml:"StopType"`
	// TODO: Offstreet etc
}
