package cmd

import (
  "os"
  "fmt"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "log-forwarder",
  Short: "Test log-forwarder for fluentd in go",
  Long: `An attempt in a single, simple binary for forwarding logs and metrics to a fluentd or fluent-bit endpoint`,

  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}
