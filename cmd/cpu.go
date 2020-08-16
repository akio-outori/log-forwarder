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

func init() {
  cpuinfo.Flags().BoolVarP(&info,    "info",    "i", false, "whether to include cpuinfo in the final result")
  cpuinfo.Flags().BoolVarP(&percent, "percent", "p", false, "whether to include cpu usage percent in final result")
  cpuinfo.Flags().BoolVarP(&cputime, "cputime", "t", false, "whether to include cpu time in the final result")
  cpuinfo.Flags().BoolVarP(&percpu,  "percpu",  "c", false, "whether to split metrics per cpu core")
  rootCmd.AddCommand(cpuinfo)
}

func cpuInfo() []byte {
  info, _ := cpu.Info()
  return format(info)
}

func cpuPercent() []byte {
  percent, _ := cpu.Percent(0, percpu)
  return format(percent)
}

func cpuTime() []byte {
  time, _ := cpu.Times(percpu)
  return format(time)
}

func format(data interface{}) []byte {
  json, _ := helpers.ConvertToJson(data)
  return json
}

var cpuinfo = &cobra.Command {

  Use:   "cpuinfo",
  Short: "Return CPU metrics as a json string",
  Long:  `Command should get linux / unix cpu and return it
  in a JSON formatted string that can be consumed by fluentd`,

  Run: func(cmd *cobra.Command, args []string) {

    var output []byte

    var info_out    []cpu.InfoStat
    var cputime_out []cpu.TimesStat
    var percent_out []float64

    if info == true {
      info_out = cpuInfo()
    }

    if percent == true {
      percent_out = cpuPercent()
    }

    if cputime == true {
      cputime_out = cpuTime()
    }

    output = SerializeJson(info_out, percent_out, cputime_out)
    fmt.Print(string(output))

  },
}
