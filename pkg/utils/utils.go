package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody reads the body of the request and unmarshals the JSON into a Go object
func ParseBody(r *http.Request, v interface{}) error {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), v); err != nil {	
			return err
		}
	}
	return nil
}