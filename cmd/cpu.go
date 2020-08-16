package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/shirou/gopsutil/cpu"
  "github.com/akio-outori/log-forwarder/helpers"
)

var info    bool
var percent bool
var cputime bool
var percpu  bool

type document struct {
  Hostname string
  Role     string
  Info     []cpu.InfoStat
  Percent  []float64
  Time     []cpu.TimesStat
}

func init() {
  cpuinfo.Flags().BoolVarP(&info,    "info",    "i", false, "whether to include cpuinfo in the final result")
  cpuinfo.Flags().BoolVarP(&percent, "percent", "p", false, "whether to include cpu usage percent in final result")
  cpuinfo.Flags().BoolVarP(&cputime, "cputime", "t", false, "whether to include cpu time in the final result")
  cpuinfo.Flags().BoolVarP(&percpu,  "percpu",  "c", false, "whether to split metrics per cpu core")
  rootCmd.AddCommand(cpuinfo)
}

func cpuInfo() []cpu.InfoStat {
  info, _ := cpu.Info()
  return info
}

func cpuPercent() []float64 {
  percent, _ := cpu.Percent(0, percpu)
  return percent
}

func cpuTime() []cpu.TimesStat {
  time, _ := cpu.Times(percpu)
  return time
}

func format(data interface{}) []byte {
  json, _ := helpers.ConvertToJson(data)
  return json
}

var cpuinfo = &cobra.Command {

  Use:   "cpuinfo",
  Short: "Return CPU metrics as a json string",
  Long:  `Command should get linux / unix cpu and return it in a JSON formatted string that can be consumed by fluentd`,

  Run: func(cmd *cobra.Command, args []string) {

    var data   document
    var output []byte

    if info == true {
      data.Info = cpuInfo()
    }

    if percent == true {
      data.Percent = cpuPercent()
    }

    if cputime == true {
      data.Time = cpuTime()
    }

    output, _ = helpers.ConvertToJson(data)
    fmt.Print(string(output))

  },
}
