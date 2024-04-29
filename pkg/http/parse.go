package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

// Unmarshal and validates the input.
func Parse(r *http.Request, out interface{}) error {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return err
	}
	return Check(out)
}

// Checks constraints on `validate` tags.
func Check(in interface{}) error {
	value := reflect.ValueOf(in)
	switch value.Kind() {
	case reflect.Ptr:
		return Check(value.Elem().Interface())
	case reflect.Struct:
		return validate.Struct(in)
	}
	return nil
}
