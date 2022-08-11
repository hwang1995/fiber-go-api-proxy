# Go Parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_LAMBDA_ENVIRONMENT=GOOS=linux GOARCH=amd64
BINARY_NAME=main

# AWS Parameters
AWSCMD=aws
DEPLOYCMD=$(AWSCMD) cloudformation
DEPLOY_BUCKET_NAME=lambda-zip-golang
STACK_NAME=aws-lambda-fiber-go-api-proxy
all: clean build package
build:
	$(BUILD_LAMBDA_ENVIRONMENT) $(GOBUILD) ./$(BINARY_NAME).go
package:
	zip main.zip $(BINARY_NAME)
clean:
	rm -f ./main
	rm -f ./main.zip
deploy:
	$(DEPLOYCMD) package --template-file sam.yaml --output-template output-sam.yaml --s3-bucket $(DEPLOY_BUCKET_NAME)
	$(DEPLOYCMD) deploy --template-file output-sam.yaml --stack-name=$(STACK_NAME) --capabilities CAPABILITY_IAM