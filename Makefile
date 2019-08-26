build:
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/create handlers/create/main.go
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/notify handlers/notify/main.go
