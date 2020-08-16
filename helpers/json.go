package helpers

import "encoding/json"

func SerializeJson(data [][]byte) []string {
  var output []string

  for _, doc := range data {
    output = append(output, string(doc))
  }
  return output
}

func ConvertToJson(data interface{}) ([]byte, error) {
  json, err := json.Marshal(data)
  return json, err
}
