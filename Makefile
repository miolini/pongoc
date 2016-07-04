build:
	go get -v
	CGO_ENABLED=0 go build -v -o bin/osx_amd64/pongoc main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/linux_amd64/pongoc main.go
