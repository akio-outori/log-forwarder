package cmd

import (
  "encoding/json"
  "github.com/spf13/cobra"
)

func ConvertToJson(data interface{}) (string, error) {
  json, err := json.Marshal(data)
  return string(json), err
}

func Execute() error {
  return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
  Use:   "log-forwarder",
  Short: "Test log-forwarder for fluentd in go",
  Long: `An attempt in a single, simple binary for forwarding logs and metrics to a fluentd or fluent-bit endpoint`,

  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}
