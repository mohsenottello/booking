package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

func (f Form) Valid() bool {
	if len(f.Errors) == 0 {
		return true
	} else {
		return false
	}
}

// Form create custom form struct and embeds url.Values objects
type Form struct {
	url.Values
	Errors errors
}

// New initials
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	return x != ""
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	value := r.Form.Get(field)

	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d charachters long", length))
		return false
	}

	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
