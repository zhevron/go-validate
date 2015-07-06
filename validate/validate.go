// Package validate provides object field validation through reflection.
package validate

import (
	"fmt"
	"reflect"

	"github.com/zhevron/go-validate/validate/validators"
)

// validator defines the function signature for a validator.
type validator func(string, string, reflect.Value) error

// Maps each validator function to its tag name.
var validatorFuncs = map[string]validator{
	"nil":    validators.Nil,
	"min":    validators.Min,
	"max":    validators.Max,
	"minLen": validators.MinLen,
	"maxLen": validators.MaxLen,
}

// Validate attempts to validate a struct object, looping each field in the
// object and checking for validation tags.
func Validate(o interface{}) (bool, []error) {
	var err []error
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println(v.Kind().String())
		err = append(err, fmt.Errorf("Type %#q is not a struct", v.Type().Name()))
		return false, err
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		fv := v.FieldByName(f.Name)

		switch fv.Kind() {
		case reflect.Array:
		case reflect.Interface:
		case reflect.Slice:
		case reflect.Struct:
			if reflect.DeepEqual(v.Interface(), fv.Interface()) {
				continue
			}
		}

		if len(f.Tag) > 0 {
			err = append(err, validateField(f, v.FieldByName(f.Name))...)
		}
	}

	return len(err) == 0, err
}

func validateField(f reflect.StructField, v reflect.Value) []error {
	var err []error

	if f.Type.Kind() == reflect.Struct {
		_, errs := Validate(v.Interface())
		return errs
	}

	for tag, validate := range validatorFuncs {
		if len(f.Tag.Get(tag)) > 0 {
			if e := validate(f.Name, f.Tag.Get(tag), v); e != nil {
				err = append(err, e)
			}
		}
	}

	return err
}
