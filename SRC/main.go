package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Network struct {
	Name     string
	Password string
}

func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func createNetwork() Network {
	rand.Seed(time.Now().UnixNano())
	name := generateRandomString(10)
	password := generateRandomString(10)

	network := Network{
		Name:     name,
		Password: password,
	}

	fmt.Printf("Network created! Name: %s, Password: %s\n", network.Name, network.Password)

	return network
}

func connectToNetwork() {
	var command string
	fmt.Println("Enter command to connect to a network:")
	fmt.Scanln(&command)

	// Here you should implement the logic to connect to the network using the command
	// This is a placeholder as the actual implementation would depend on the networking library you are using
	fmt.Printf("Connecting to network with command: %s\n", command)
}

func main() {
	createNetwork()
	connectToNetwork()
}
