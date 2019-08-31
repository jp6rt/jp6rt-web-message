# jp6rt-web-message
Service for creating message record and forwarding it to SMS and email

## Technology stack
* Golang
* Serverless Framework
* Circle CI for continous deployment
* AWS Cloudformation
* AWS DynamoDB
* AWS SNS

## Build

#### Create message
> GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/create handlers/create/main.go

#### Notify SNS subscribers
> GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/notify handlers/notify/main.go

## Deployment

### Using CircleCI
Release tags will trigger a deploy to production environment

### Manual

Cloudformation outputs are referenced from the service so we deploy the cloudformation templates before deploying the microservices.
Templates can be found on the [infra folder](./infra)

#### Deploy templates

##### Deploy dev stack
> aws cloudformation deploy --stack-name dev-jp6rt-web --template-file ./infra/resources-template.yml

##### Deploy production stack
> aws cloudformation deploy --stack-name prod-jp6rt-web --template-file ./infra/resources-template.yml
