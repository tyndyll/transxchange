package nptg

import (
	"time"
)

type NationalPublicTransportGazetteer struct {
	// Timestamp of document creation date and time
	Created time.Time `xml:"CreationDateTime,attr"`

	// Timestamp of document last modification date and time
	Modified time.Time `xml:"ModificationDateTime,attr"`

	// Name of file containing the document.
	FileName string `xml:"FileName,attr"`

	// Details the nature of the change: new or revision. For NPTG documents this will
	// always be ’revision’. Individual elements within the document may be ’new’.
	Modification string `xml:"Modification,attr"`

	// Optional sequence number for versioning overall document content. Each subsequent
	// issue of the NPTG data should have a higher number than the previous one.
	RevisionNumber string `xml:"RevisionNumber,attr"`

	// Schema version identifier used for the document content model.
	SchemaVersion string `xml:"SchemaVersion,attr"`
}
