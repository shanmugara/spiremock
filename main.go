package main

import (
	"log"
	"os"
	"spiremock/spiremock"
	"time"
)

func main() {
	// Ensure logs go to stdout so they are captured by container/pod logs
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	log.Println("Calling NewMockClient...")
	if true {
		c, err := spiremock.NewTlsMockClient()
		if err != nil {
			log.Println("Error creating TLS mock client:", err)
		} else {
			log.Println("TLS mock client created:", c)
		}

		b, j, err := spiremock.NewJWTMockClient()
		if err != nil {
			log.Println("Error creating JWT mock client:", err)
		} else {
			if j != nil {
				log.Println("JWT mock client created successfully:", j.Marshal())
			} else {
				log.Println("JWT mock client created successfully: <nil SVID>")
			}

			if b == nil {
				log.Println("JWT bundle set is nil")
			} else {
				for _, bu := range b.Bundles() {
					log.Printf("Bundle for trust domain: %s (keys: %d)", bu.TrustDomain(), len(bu.JWTAuthorities()))
				}
			}
		}

		time.Sleep(5 * time.Second)
	}

}
