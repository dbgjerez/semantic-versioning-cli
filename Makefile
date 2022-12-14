go test -v ./... -coverprofile=coverage.out -covermode count
go tool cover -func coverage.out 
go build -ldflags "-X main.version="$(semver info v)
