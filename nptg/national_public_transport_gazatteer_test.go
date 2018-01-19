package nptg

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	"golang.org/x/text/encoding/charmap"
)

func TestNationalPublicTransportGazetteer_Unmarshal(t *testing.T) {
	testXML := bytes.NewBuffer([]byte(`
	<?xml version="1.0" encoding="Windows-1252"?>
	<TransXChange CreationDateTime="2017-12-15T09:27:14.0355862Z" ModificationDateTime="2017-12-15T09:27:14.0511865Z" Modification="new" FileName="TXC_2017-12-15T09-26-59_SHARON.xml" RevisionNumber="1" SchemaVersion="2.4" xml:lang="en" xsi:schemaLocation="http://www.transxchange.org.uk/ http://www.transxchange.org.uk/schema/2.4/TransXChange_general.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.transxchange.org.uk/">
	</TransXChange>
	`))

	gazetteer := &NationalPublicTransportGazetteer{}

	decoder := xml.NewDecoder(testXML)
	decoder.CharsetReader = makeCharsetReader
	if err := decoder.Decode(&gazetteer); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if gazetteer.Created.String() != "2017-12-15 09:27:14.0355862 +0000 UTC" {
		t.Error(fmt.Errorf("Created attribute incorrect (was %s)", gazetteer.Created.String()))
		t.Fail()
	}

	if gazetteer.Modified.String() != "2017-12-15 09:27:14.0511865 +0000 UTC" {
		t.Error(fmt.Errorf("Modified attribute incorrect (was %s)", gazetteer.Created.String()))
		t.Fail()
	}

	if gazetteer.FileName != "TXC_2017-12-15T09-26-59_SHARON.xml" {
		t.Error(fmt.Errorf("FileName attribute incorrect (was %s)", gazetteer.FileName))
		t.Fail()
	}

	if gazetteer.Modification != "new" {
		t.Error(fmt.Errorf("Modification attribute incorrect (was %s)", gazetteer.Modification))
		t.Fail()
	}

	if gazetteer.RevisionNumber != "1" {
		t.Error(fmt.Errorf("RevisionNumber attribute incorrect (was %s)", gazetteer.RevisionNumber))
		t.Fail()
	}

	if gazetteer.SchemaVersion != "2.4" {
		t.Error(fmt.Errorf("SchemaVersion attribute incorrect (was %s)", gazetteer.SchemaVersion))
		t.Fail()
	}
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch charset {
	case "ISO-8859-1":
	case "Windows-1252":
		fmt.Println("In reader")
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("Unknown charset: %s", charset)
}
