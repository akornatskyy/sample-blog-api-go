package iojson

import (
	"encoding/json"
	"io/ioutil"
)

// ReadFile reads the file named by filename as JSON-encoded
// value and stores it in the value pointed to by v.
func ReadFile(filename string, v interface{}) error {
	value, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(value, v)
}
