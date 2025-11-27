package main

import (
	"flag"
	"log"
	"os"
	"spiremock/spiremock"
	"time"

	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

func main() {
	// Ensure logs go to stdout so they are captured by container/pod logs
	var dlgns string
	var dlgsa string
	var audience string
	var podlabel string

	flag.StringVar(&podlabel, "pod-label", "spiffe.io/cluster:omega-admin.omegaworld.net", "Pod labels for DLG JWT SVID selectors")
	flag.StringVar(&dlgns, "dlg-namespace", "myapp", "Namespace for DLG JWT SVID")
	flag.StringVar(&dlgsa, "dlg-service-account", "mydlgsa", "Service account for DLG JWT SVID")
	flag.StringVar(&audience, "audience", "omega", "Audience for DLG JWT SVID")
	flag.Parse()
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	log.Println("Starting mock client refresh loop (every 15 seconds)...")

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	// Function to create and test clients
	testClients := func() {
		log.Println("Creating/refreshing mock clients...")

		// Test TLS mock client
		c, err := spiremock.NewTlsMockClient()
		if err != nil {
			log.Println("Error creating TLS mock client:", err)
		} else {
			log.Println("TLS mock client created:", c)
		}

		// Test JWT mock client
		b, j, err := spiremock.NewJWTMockClient(audience)
		if err != nil {
			log.Println("Error creating JWT mock client:", err)
		} else {
			if j != nil {
				// jwtsvid.SVID does not expose a Marshal method; print Subject and SPIFFE ID instead
				log.Printf("JWT mock client created successfully: %s", j.Marshal())
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
		// Test DLG JWT mock client
		selectors := []*types.Selector{
			{Type: "k8s", Value: "pod-label:" + podlabel},
			{Type: "k8s", Value: "sa:" + dlgsa},
			{Type: "k8s", Value: "ns:" + dlgns},
		}
		err = spiremock.NewDLGJWTMockClient(selectors, audience)
		if err != nil {
			log.Println("Error creating DLG JWT mock client:", err)
		} else {
			log.Println("DLG JWT mock client created successfully")
		}

	}

	// Run immediately on startup
	testClients()

	// Then repeat every 15 seconds
	for range ticker.C {
		testClients()
	}

}
