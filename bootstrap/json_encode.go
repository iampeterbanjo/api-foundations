package bootstrap

import (
	"encoding/json"
)

// JSONEncode struct
func JSONEncode(r interface{}) string {
	jsonString, _ := json.MarshalIndent(r, "", "\t")
	return string(jsonString[:])
}
