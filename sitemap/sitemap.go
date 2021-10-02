package sitemap

import (
	"errors"
	"os"
)

const header = `<?xml version="1.0" encoding="UTF-8"?>`

type Sitemap struct {
	xmlFile	string
	urlSet	URLSet
}

func (sitemap *Sitemap) Add(url URL) error {
	return sitemap.urlSet.add(url)
}

func (sitemap *Sitemap) Write() error {
	xml, err := sitemap.urlSet.xml()
	if err != nil {
		return err
	}

	file, _ := os.OpenFile(sitemap.xmlFile, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0666)
	status, err := file.Stat()
	if err != nil {
		return errors.New("file not exists")
	} else if !status.Mode().IsRegular() {
		return errors.New("is not a filename")
	}
	defer file.Close()

	_, err = file.Write([]byte(header))
	if err != nil {
		return err
	}
	_, err = file.Write(xml)

	return err
}

func New(xmlFile string) *Sitemap {
	return &Sitemap{xmlFile: xmlFile}
}
