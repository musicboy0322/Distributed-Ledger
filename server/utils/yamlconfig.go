package utils

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func GetServerPort() string {
	// set viper read config
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
    serverPorts := vp.GetIntSlice("server_ports")
	serverPort := strconv.Itoa(serverPorts[0])
	return serverPort
}