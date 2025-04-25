package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// value receivers can be invoked on both values and ptrs
func (r Runtime) MarshalJSON() ([]byte, error) {
	// generate a str containing the movie runtime in the required format
	jsonValue := fmt.Sprintf("%d mins", r)

	// use Quote to wrap it in double quotes so its valid JSON
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
