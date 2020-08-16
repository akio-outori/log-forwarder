package helpers

import "encoding/json"

func SerializeJson(data... []byte) [][]byte {
  var output [][]byte
  output = append(output, data...)
  return output
}

func ConvertToJson(data interface{}) ([]byte, error) {
  json, err := json.Marshal(data)
  return json, err
}
