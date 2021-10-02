package sitemap

import (
	"encoding/xml"
	"testing"
)

func TestAddURL(t *testing.T) {
	url := URL{}
	set := URLSet{}

	if err := set.add(url); err == nil {
		t.Errorf("URLSet.Add: adding invalid URL")
	}

	url.Location = "http://example.com"
	if err := set.add(url); err != nil {
		t.Fatal(err)
	}
}

func TestXML(t *testing.T) {
	url := URL{Location: "http://example.com"}
	set := URLSet{}
	xmlEncoded, err := set.xml()
	if err != nil {
		t.Fatal(err)
	}

	result, _ := xml.MarshalIndent(&set, "  ", "  ")
	if string(xmlEncoded) != string(result) {
		t.Errorf("URLSet.xml: bad encoding XML -  '%s'",string(xmlEncoded))
	}

	err = set.add(url)
	xmlEncoded, err = set.xml()
	if err != nil {
		t.Fatal(err)
	}

	result, _ = xml.MarshalIndent(&set, "  ", "  ")
	if string(xmlEncoded) != string(result) {
		t.Errorf("URLSet.xml: bad encoding XML -  '%s'",string(xmlEncoded))
	}
}
