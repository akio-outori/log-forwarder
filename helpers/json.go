package helpers

import "encoding/json"

func ConvertToJson(data interface{}) ([]byte, error) {
  json, err := json.Marshal(data)
  return json, err
}
