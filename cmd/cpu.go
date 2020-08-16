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

func cpuInfo() ([]cpu.InfoStat, error) {
  return cpu.Info()
}

func cpuPercent() ([]float64, error) {
  return cpu.Percent(0, percpu)
}

func cpuTime() ([]cpu.TimesStat, error) {
  return cpu.Times(percpu)
}

func format(data interface{}) ([]byte, error) {
  return helpers.ConvertToJson(data)
}

var cpuinfo = &cobra.Command {

  Use:   "cpuinfo",
  Short: "Return CPU metrics as a json string",
  Long:  `Command should get linux / unix cpu and return it in a JSON formatted string that can be consumed by fluentd`,

  Run: func(cmd *cobra.Command, args []string) {

    var data   document
    var output []byte

    if info == true {
      data.Info, _ = cpuInfo()
    }

    if percent == true {
      data.Percent, _ = cpuPercent()
    }

    if cputime == true {
      data.Time, _ = cpuTime()
    }

    output, _ = format(data)
    fmt.Print(string(output))

  },
}
