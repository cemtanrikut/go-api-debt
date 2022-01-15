package helper

import (
	"encoding/json"

	"github.com/cemtanrikut/go-api-debt/api"
)

func JsonMarshal(result api.Response) []byte {
	jsonResult, jsonError := json.Marshal(result)
	if jsonError != nil {
		return []byte(jsonError.Error())
	}
	return jsonResult
}
