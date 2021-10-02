package sitemap

import (
	"testing"
)

func TestAdd(t *testing.T) {
	url := URL{}
	sitemap := &Sitemap{}

	if err := sitemap.Add(url); err == nil {
		t.Errorf("Sitemap.Add: adding invalid URL")
	}

	url.Location = "http://example.com"
	if err := sitemap.Add(url); err != nil {
		t.Fatal(err)
	}
}
