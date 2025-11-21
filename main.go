package main

import (
	"fmt"
	"spiremock/spiremock"
)

func main() {
	fmt.Println("Calling NewMockClient...")
	spiremock.NewMockClient()
	fmt.Println("Done NewMockClient...")
}
