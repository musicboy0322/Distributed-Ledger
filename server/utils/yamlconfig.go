package utils

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func GetServerPort(serverNumber int) string {
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
	serverPort := strconv.Itoa(serverPorts[serverNumber])
	return serverPort
}

func GetOtherPorts(currentPort string) []string {
	// initialize variable
	var otherPorts []string

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
	for _, port := range serverPorts {
		serverPort := strconv.Itoa(port)
		if serverPort == currentPort {
			continue
		}
		otherPorts = append(otherPorts, serverPort)
	}
	return otherPorts
}