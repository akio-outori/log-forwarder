package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/shirou/gopsutil/cpu"
)

var info     bool
var percent  bool
var cputime  bool
var percpu   bool

type cpudata struct {

  // Import global data 
  data

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

var cpuinfo = &cobra.Command {
  Use:   "cpuinfo",
  Short: "Return CPU metrics as a json string",
  Long:  `Command should get linux / unix cpu and return it in a JSON formatted string that can be consumed by fluentd`,
  Run:   func(cmd *cobra.Command, args []string) {

    var response cpudata
    var json     []byte

    config            := initConfig()
    response.Hostname  = config.GetString("Hostname")
    response.Role      = config.GetString("Role")
    response.Config    = config.GetString("Config")

    if info == true {
      response.Info, _ = cpuInfo()
    }

    if percent == true {
      response.Percent, _ = cpuPercent()
    }

    if cputime == true {
      response.Time, _ = cpuTime()
    }

    json, _ = convertToJson(response)
    fmt.Println(string(json))
  },
}
