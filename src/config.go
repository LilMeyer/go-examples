package examples

import (
  "fmt"
  "github.com/spf13/viper"
  "gopkg.in/pg.v5"
)

var options = pg.Options{}

func InitConfig() {

  viper.AddConfigPath("..")
  viper.AddConfigPath("../config")
  viper.SetConfigName("env")

  err := viper.ReadInConfig()

  if err != nil {
    fmt.Println("No configuration file loaded - using defaults")
  }

  // If no config is found, use the default(s)
  viper.SetDefault("msg", "Hello World (default)")

  options.User = viper.GetString("postgres.user")
  options.Password = viper.GetString("postgres.password")

  theMessage := viper.GetString("msg")
  fmt.Println(theMessage)
}
