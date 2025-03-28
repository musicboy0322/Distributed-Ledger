package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomServers(servers []string) string {
	// Check if there are any ports to choose from
	if len(servers) == 0 {
		fmt.Println("No ports available")
		return "-1"  // Return an invalid port or handle the case as needed
	}
	
	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())
	
	// Generate a random index based on the length of the ports slice
	randomIndex := rand.Intn(len(servers))
	// Return the randomly selected port
	return servers[randomIndex]
}