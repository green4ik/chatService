package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts api key from the headers
//of http request
//example:
//	Authorization: ApiKey : {insert api key here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication information found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part authentication header")
	}
	return vals[1], nil
}
