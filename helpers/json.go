package helpers

import "encoding/json"

func ConvertToJson(response interface{}) ([]byte, error) {
  json, err := json.Marshal(response)
  return json, err
}
