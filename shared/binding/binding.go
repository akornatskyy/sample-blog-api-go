// Package binding updates struct from a values map.
package binding

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
)

// Bind parses values and sets corresponding fields per binding tag of not
// nil struct pointer.
//
// Example:
//
//	package main
//
// 	var m struct {
//		Test string `binding:"test"`
// 	}
//
// 	values := map[string][]string{"test": {"hello"}}
// 	err := binding.Bind(&m, values)
//
// It supports the following types: string, int, uint, bool, time.Duration,
// time.Time, slice.
func Bind(m interface{}, values map[string][]string) error {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("expects not nil ptr")
	}
	v = v.Elem()
	var e *errorstate.ErrorState
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag
		name, ok := tag.Lookup("binding")
		if !ok {
			continue
		}
		value := values[name]
		if len(value) == 0 {
			continue
		}
		if err := setValues(v.Field(i), value, tag); err != nil {
			if e == nil {
				e = &errorstate.ErrorState{}
			}
			e.Errors = append(e.Errors, &errorstate.Detail{
				Domain:   "binding",
				Type:     "field",
				Location: name,
				Reason:   err.Error(),
				Message:  "The input is not in supported format.",
			})
		}
	}
	if e == nil {
		return nil
	}
	return e
}

func setValues(v reflect.Value, values []string, tag reflect.StructTag) error {
	switch v.Kind() {
	case reflect.Slice:
		return setSlice(v, values, tag)
	default:
		return setValue(v, values[0], tag)
	}
}

func setValue(v reflect.Value, s string, tag reflect.StructTag) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(s)
	case reflect.Int:
		return setInt(v, s)
	case reflect.Int64:
		switch v.Interface().(type) {
		case time.Duration:
			value, err := time.ParseDuration(s)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(value))
		}
	case reflect.Uint:
		return setUint(v, s)
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err.(*strconv.NumError).Err
		}
		v.SetBool(b)
	case reflect.Struct:
		switch v.Interface().(type) {
		case time.Time:
			return setTime(v, s, tag)
		}
	}
	return nil
}

func setInt(v reflect.Value, s string) error {
	value, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return err.(*strconv.NumError).Err
	}
	v.SetInt(value)
	return nil
}

func setUint(v reflect.Value, s string) error {
	value, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return err.(*strconv.NumError).Err
	}
	v.SetUint(value)
	return nil
}

func setTime(v reflect.Value, s string, tag reflect.StructTag) error {
	layout, ok := tag.Lookup("layout")
	if !ok {
		layout = time.RFC3339
	}
	loc := time.UTC
	if l, ok := tag.Lookup("loc"); ok {
		l, err := time.LoadLocation(l)
		if err != nil {
			return err
		}
		loc = l
	}
	tm, err := time.ParseInLocation(layout, s, loc)
	if err != nil {
		err := err.(*time.ParseError)
		return errors.New("invalid time" + err.Message)
	}
	v.Set(reflect.ValueOf(tm))
	return nil
}

func setSlice(v reflect.Value, values []string, tag reflect.StructTag) error {
	l := len(values)
	slice := reflect.MakeSlice(v.Type(), l, l)
	for i, s := range values {
		err := setValue(slice.Index(i), s, tag)
		if err != nil {
			return err
		}
	}
	v.Set(slice)
	return nil
}
