# jp6rt-web-message
Service for creating message record and forwarding it to SMS and email

# Build

## create message
> GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/create handlers/create/main.go

## notify sns subscribers
> GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/notify handlers/notify/main.go

# Deployment

Cloudformation outputs are referenced from the service so we deploy the cloudformation templates before deploying the microservices.

Templates can be found on the [infra folder](./infra)

## deploy templates

### Deploy dev stack
> aws cloudformation deploy --stack-name dev-jp6rt-web --template-file ./infra/resources-template.yml

### Deploy production stack
> aws cloudformation deploy --stack-name prod-jp6rt-web --template-file ./infra/resources-template.yml
