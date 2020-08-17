package cmd

import (
  "encoding/json"
  "github.com/spf13/viper"
)

type data struct {
  Hostname string
  Role     string
  Config   string
}

func initConfig() *viper.Viper {

  config := viper.New()
  config.SetDefault("hostname", "undefined")
  config.SetDefault("role",     "undefined")
  config.SetDefault("config",   "undefined")

  config.SetConfigFile("config.yaml")
  config.SetConfigType("yaml")
  config.AddConfigPath(".")

  if err := config.ReadInConfig(); err != nil {
    panic(err)
  }

  return config
}

func checkError(err error) {
  if err != nil {
    panic(err)
  }
}

func convertToJson(response interface{}) ([]byte, error) {
  json, err := json.Marshal(response)
  return json, err
}
