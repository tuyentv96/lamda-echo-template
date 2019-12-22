package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tuyentv96/lamda-echo-template/handler"
	"log"
)

var echoLambda *echoadapter.EchoLambda

func initHttp() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "hello")
	})

	e.GET("/ping", func(c echo.Context) error {
		log.Println("Handler!!")
		return c.String(200, "pong")
	})

	// Init article handler
	handler.NewArticleHandler(e)

	echoLambda = echoadapter.New(e)
	// Start server
	//e.Logger.Fatal(e.Start(":1323"))
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	initHttp()
	lambda.Start(Handler)
}
