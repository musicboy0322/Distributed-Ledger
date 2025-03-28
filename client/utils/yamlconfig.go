package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetEnterPorts() []int {
	// set viper read config
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
    enterPorts := vp.GetIntSlice("server_ports")
	return enterPorts
}

func GetEnterServer() []string {
	// set viper read config
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	enterServers := vp.GetStringSlice("server_address")
	return enterServers
}