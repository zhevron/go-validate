package validators

import (
	"fmt"
	"reflect"
	"strconv"
)

// Min validates a numeric reflect.Value to be at least the value of tag.
func Min(name string, tag string, v reflect.Value) error {
	min, err := strconv.ParseFloat(tag, 64)
	if err != nil {
		return fmt.Errorf(
			"The tag %#q could not be converted to a numeric value",
			tag,
		)
	}

	var n float64
	switch v.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		n = float64(v.Int())

	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		n = float64(v.Uint())

	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		n = v.Float()

	default:
		return fmt.Errorf("Value %#q is not a valid numeric type", name)
	}

	if n < min {
		return fmt.Errorf("Value %#q (%f) is lower than %f", name, n, min)
	}
	return nil
}

// Max validates a numeric reflect.Value to be no higher than the value of tag.
func Max(name string, tag string, v reflect.Value) error {
	max, err := strconv.ParseFloat(tag, 64)
	if err != nil {
		return fmt.Errorf(
			"The tag %#q could not be converted to a numeric value",
			tag,
		)
	}

	var n float64
	switch v.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		n = float64(v.Int())

	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		n = float64(v.Uint())

	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		n = v.Float()

	default:
		return fmt.Errorf("Value %#q is not a valid numeric type", name)
	}

	if n > max {
		return fmt.Errorf("Value %#q (%f) is higher than %f", name, n, max)
	}
	return nil
}
