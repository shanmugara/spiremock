package main

import (
	"fmt"
	"spiremock/spiremock"
	"time"
)

func main() {
	fmt.Println("Calling NewMockClient...")
	if true {
		c, err := spiremock.NewTlsMockClient()
		if err != nil {
			fmt.Println("Error creating TLS mock client:", err)
		} else {
			fmt.Println("TLS mock client created:", c)
		}

		j, err := spiremock.NewJWTMockClient()
		if err != nil {
			fmt.Println("Error creating JWT mock client:", err)
		} else {
			fmt.Println("JWT mock client created successfully", j)
		}

		time.Sleep(5 * time.Second)
	}

}
