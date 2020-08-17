package cmd

import "encoding/json"

func convertToJson(response interface{}) ([]byte, error) {
  json, err := json.Marshal(response)
  return json, err
}
