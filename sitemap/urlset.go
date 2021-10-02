package sitemap

import (
	"encoding/xml"
)

type URLSet struct {
	XMLName	xml.Name	`xml:"urlset"`
	URLs	[]URL		`xml:"url"`
}

func (set *URLSet) add(url URL) error {
	if err := url.validate(); err != nil {
		return err
	}
	set.URLs = append(set.URLs, url)

	return nil
}

func (set *URLSet) xml() ([]byte, error) {
	return xml.MarshalIndent(set, "  ", "  ")
}
