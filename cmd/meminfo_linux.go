package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/shirou/gopsutil/mem"
)

var swap bool
var vm   bool
var vmex bool

type memdata struct {

  // Import global data 
  data

  Swap            *mem.SwapMemoryStat
  VirtualMemory   *mem.VirtualMemoryStat
  VirtualMemoryEx *mem.VirtualMemoryExStat
}

func init() {
  meminfo.Flags().BoolVarP(&swap, "swap", "s", false, "whether to collect swap information")
  meminfo.Flags().BoolVarP(&vm,   "vm",   "m", false, "whether to collect memory information")
  meminfo.Flags().BoolVarP(&vmex, "ex",   "e", false, "whether to collect extended memory information")
  rootCmd.AddCommand(meminfo)
}

func swapMemory() (*mem.SwapMemoryStat, error) {
  return mem.SwapMemory()
}

func virtualMemory() (*mem.VirtualMemoryStat, error) {
  return mem.VirtualMemory()
}

func virtualMemoryEx() (*mem.VirtualMemoryExStat, error) {
  return mem.VirtualMemoryEx()
}

var meminfo = &cobra.Command {
  Use:   "meminfo",
  Short: "Return memory metrics as a json string",
  Long:  `Command should get linux / unix memory and return it in a JSON formatted string that can be consumed by fluentd`,
  Run:   func(cmd *cobra.Command, args []string) {

    var response memdata
    var json     []byte

    config            := initConfig()
    response.Hostname  = config.GetString("Hostname")
    response.Role      = config.GetString("Role")
    response.Config    = config.GetString("Config")

    if swap == true {
      response.Swap, _ = swapMemory()
    }

    if vm == true {
      response.VirtualMemory, _ = virtualMemory()
    }

    if vmex == true {
      if mem.(VirtualMemoryExStat) {
        response.VirtualMemoryEx, _ = virtualMemoryEx()
      }
    }

    json, _ = convertToJson(response)
    fmt.Println(string(json))
  },
}
