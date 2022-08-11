package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

var fiberLambda *fiberadapter.FiberLambda
var app *fiber.App

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func init() {
	app = fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello, World")
	})

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	if inLambda() {
		log.Println("AWS Lambda Fiber Start")
		lambda.Start(Handler)
	} else {
		log.Println("Local PC Start")
		app.Listen(":8080")
	}

}
