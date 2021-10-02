package sitemap

import (
	"testing"
)

func TestRequiredLocation(t *testing.T) {
	url := URL{}
	if err := url.validate(); err == nil {
		t.Errorf("Required URL.Location")
	}
}

func TestInvalidLocation(t *testing.T) {
	url := URL{Location: "invalid"}
	if err := url.validate(); err == nil {
		t.Errorf("Invalid URL.Location")
	}
}

func TestValidLocation(t *testing.T) {
	url := URL{Location: "http://example.com"}
	if err := url.validate(); err != nil {
		t.Fatal(err)
	}
}

func TestCheckChangeFrequently(t *testing.T) {
	values := []string{"", "always", "hourly", "daily", "weekly", "monthly", "yearly", "never"}
	url := URL{Location: "http://example.com"}
	for _, value := range values {
		url.ChangeFrequently = value
		if err := url.validate(); err != nil {
			t.Errorf("URL.ChangeFrequently: invalid value '%s'", value)
		}
	}
}

func TestCheckLastModified(t *testing.T) {
	url := URL{Location: "http://example.com", LastModified: "2021-02-08T18:19:00+01:00"}
	if err := url.validate(); err != nil {
		t.Fatal(err)
	}
	url.LastModified = ""
	if err := url.validate(); err != nil {
		t.Fatal(err)
	}
	url.LastModified = "2020-06-15"
	if err := url.validate(); err == nil {
		t.Errorf("URL.LastModified: invalid time standard")
	}
}

func TestCheckPriority(t *testing.T) {
	url := URL{Location: "http://example.com", Priority: 1.0}
	if err := url.validate(); err != nil {
		t.Errorf("URL.Priority: invalid value '1.0'")
	}
	url.Priority = 0.0
	if err := url.validate(); err != nil {
		t.Errorf("URL.Priority: invalid value '0.0'")
	}
	url.Priority = -0.9
	if err := url.validate(); err == nil {
		t.Errorf("URL.Priority: invalid value '-0.9'")
	}
	url.Priority = 1.1
	if err := url.validate(); err == nil {
		t.Errorf("URL.Priority: invalid value '1.1'")
	}
}
