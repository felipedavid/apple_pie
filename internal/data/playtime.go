package data

import (
	"fmt"
	"strconv"
)

type playtime int32

func (p playtime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("%d min", p)

	// To be a valid json object, it needs to be wrapped around double quotes
	quotedJson := strconv.Quote(json)

	return []byte(quotedJson), nil
}
