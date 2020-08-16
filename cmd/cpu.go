package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/shirou/gopsutil/cpu"
)

var info bool
var percent bool

func init() {
  cpuinfo.Flags().BoolVarP(&info, "info", "i", false, "whether to include cpuinfo in the final result")
  cpuinfo.Flags().BoolVarP(&percent, "percent", "p", false, "whether to include cpu usage percent in final result")
  rootCmd.AddCommand(cpuinfo)
}

func cpuInfo() string {
  info, _ := cpu.Info()
  json, _ := ConvertToJson(info)
  return json
}

func cpuPercent() string {
  info, _ := cpu.Percent(0, true)
  json, _ := ConvertToJson(info)
  return json
}

var cpuinfo = &cobra.Command {

  Use:   "cpuinfo",
  Short: "Return CPU metrics as a json string",
  Long:  `Command should get linux / unix cpu and return it
  in a JSON formatted string that can be consumed by fluentd`,

  Run: func(cmd *cobra.Command, args []string) {

    if info == true {
      fmt.Print(cpuInfo())
    }

    if percent == true {
      fmt.Print(cpuPercent())
    }

  },
}
