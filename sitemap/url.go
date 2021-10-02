package sitemap

import (
	"encoding/xml"
	"errors"
	"time"

	"github.com/go-playground/validator"
)

const RFC3339 = "2006-01-02T15:04:05Z07:00"

var validate *validator.Validate

type URL struct {
	XMLName			xml.Name	`xml:"url"`
	Location		string		`xml:"loc" validate:"required,url"`
	LastModified		string		`xml:"lastmod,omitempty" validate:"RFC3339"`
	ChangeFrequently	string		`xml:"changefreq,omitempty" validate:"changefreq"`
	Priority		float64		`xml:"priority,omitempty" validate:"gte=0.0,lte=1.0"`
}

func (url *URL) validate() error {
	err := validate.Struct(url)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		errs := err.(validator.ValidationErrors)[0]

		return errors.New(errs.StructField() + ": " + errs.ActualTag())
	}

	return nil
}

func validateChangeFrequently(fieldLevel validator.FieldLevel) bool {
	validValues := []string{"", "always", "hourly", "daily", "weekly", "monthly", "yearly", "never"}
	value := fieldLevel.Field().String()
	for _, validValue := range validValues {
		if value == validValue {
			return true
		}
	}

	return false
}

func validateRFC3339(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()
	if value == "" {
		return true
	}
	_, err := time.Parse(RFC3339, value)
	if err != nil {
		return false
	}

	return true
}

func init() {
	validate = validator.New()
	validate.RegisterValidation("changefreq", validateChangeFrequently)
	validate.RegisterValidation("RFC3339", validateRFC3339)
}
