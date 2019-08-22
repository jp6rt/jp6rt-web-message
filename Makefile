build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/create handlers/create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/notify handlers/notify/main.go