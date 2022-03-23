package forms

import (
	"net/http"
	"net/url"
)

// Creates a custom form struct, embeds a url.Value object
type Form struct {
	url.Values
	Errors errors
}

// Initialize a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Check form has the field and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}

	return true
}
