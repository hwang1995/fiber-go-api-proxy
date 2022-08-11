package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/hwang1995/fiber-go-api-proxy/util"
	"go.uber.org/zap"
	"os"
)

var fiberLambda *fiberadapter.FiberLambda
var app *fiber.App
var log *zap.SugaredLogger

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

type TestData struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth-date"`
}

func init() {
	log, _ = util.GetInitializeLogger()
	app = fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello, World")
	})

	app.Get("/test", func(context *fiber.Ctx) error {
		return context.JSON(TestData{
			Name:      "HELLO WORLD",
			BirthDate: "9999-12-31",
		})
	})

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	if inLambda() {
		log.Info("AWS Lambda Fiber Start!")
		lambda.Start(Handler)
	} else {
		log.Info("Local PC Start!")
		app.Listen(":8080")
	}

}
