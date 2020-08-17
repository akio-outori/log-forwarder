package helpers

import (
  "github.com/spf13/viper"
)

type response struct {
  hostname string
  role     string
  config   string
}

func InitConfig() *viper.Viper {

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
