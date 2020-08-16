package cmd

import (
  "fmt"
  "encoding/json"
  //"strings"
  "github.com/spf13/cobra"
  "github.com/shirou/gopsutil/cpu"
)

func init() {
  rootCmd.AddCommand(cpuinfo)
}

func convert(data interface{}) (string, error) {
  json, err := json.Marshal(data)
  return string(json), err
}

func cpuInfo() string {
  info, _ := cpu.Info()
  json, _ := convert(info)
  return json
}

var cpuinfo = &cobra.Command {
  Use:   "cpuinfo",
  Short: "Return CPU metrics as a json string",
  Long:  `Command should get linux / unix cpu and return it
  in a JSON formatted string that can be consumed by fluentd`,

  Run: func(cmd *cobra.Command, args []string) {
    //usage, _ := gops.Percent(0, true)
    //fmt.Print(usage)

    fmt.Print(cpuInfo())
    //info, _ := cpuInfo()
    //fmt.Print(info)
  },
}
