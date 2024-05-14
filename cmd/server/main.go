package main

import (
	"github.com/mrangelba/go-exp-temperature/internal/infra/http/rest"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	rest.Initialize()
}
