module spiremock

go 1.25.3

// Removed unused require directives (they were reported as unused by `go mod tidy`).
// Add specific requires again only if your code imports those modules.

require (
	github.com/shanmugara/spireauthlib v0.0.0-20251121190911-144e942fe6ff
	github.com/sirupsen/logrus v1.9.3
	github.com/spiffe/go-spiffe/v2 v2.6.0
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/go-jose/go-jose/v4 v4.1.2 // indirect
	github.com/shanmugara/spiremock v0.0.0-20251121191855-452f124d5028 // indirect
	github.com/spiffe/spire-api-sdk v1.13.3 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250811230008-5f3141c8851a // indirect
	google.golang.org/grpc v1.75.0 // indirect
	google.golang.org/protobuf v1.36.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
