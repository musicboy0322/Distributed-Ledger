package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomPort(ports []int) int {
	// Check if there are any ports to choose from
	if len(ports) == 0 {
		fmt.Println("No ports available")
		return -1  // Return an invalid port or handle the case as needed
	}
	
	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())
	
	// Generate a random index based on the length of the ports slice
	randomIndex := rand.Intn(len(ports))
	fmt.Println(randomIndex)
	// Return the randomly selected port
	return ports[randomIndex]
}