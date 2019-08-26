# jp6rt-web-message
Service for creating message record and forwarding it to SMS and email

# build

## create message
> GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/create handlers/create/main.go

## notify sns subscribers
> GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/notify handlers/notify/main.go
