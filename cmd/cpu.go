package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/shirou/gopsutil/cpu"
)

var info    bool
var percent bool
var cputime bool
var percpu  bool

func init() {
  cpuinfo.Flags().BoolVarP(&info,    "info",    "i", false, "whether to include cpuinfo in the final result")
  cpuinfo.Flags().BoolVarP(&percent, "percent", "p", false, "whether to include cpu usage percent in final result")
  cpuinfo.Flags().BoolVarP(&cputime, "cputime", "t", false, "whether to include cpu time in the final result")
  cpuinfo.Flags().BoolVarP(&percpu,  "percpu",  "c", false, "whether to split metrics per cpu core")
  rootCmd.AddCommand(cpuinfo)
}

func cpuInfo() string {
  info, _ := cpu.Info()
  return format(info)
}

func cpuPercent() string {
  percent, _ := cpu.Percent(0, percpu)
  return format(percent)
}

func cpuTime() string {
  time, _ := cpu.Times(percpu)
  return format(time)
}

func format(data interface{}) string {
  json, _ := ConvertToJson(data)
  return json
}

var cpuinfo = &cobra.Command {

  Use:   "cpuinfo",
  Short: "Return CPU metrics as a json string",
  Long:  `Command should get linux / unix cpu and return it
  in a JSON formatted string that can be consumed by fluentd`,

  Run: func(cmd *cobra.Command, args []string) {

    var output []string

    if info == true {
      output = append(output, cpuInfo())
    }

    if percent == true {
      output = append(output, cpuPercent())
    }

    if cputime == true {
      output = append(output, cpuTime())
    }

    fmt.Print(output)

  },
}
