package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/deepakgonda/go-serverless-test/src/pkg/logger"
	"github.com/deepakgonda/go-serverless-test/src/pkg/middleware"
	"github.com/deepakgonda/go-serverless-test/src/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {

	// Initiating logger

	logger.InitLogger()

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Fiber cold start")
	app := fiber.New()

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	routes.PublicRoutes(app)        // Register a public routes for app.
	routes.PrivateRoutes(app)       // Register a private routes for app.
	routes.NotFoundRoute(app)       // Register route for 404 Error.

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! from Serverless Go Fiber App...")
	})

	fiberLambda = fiberadapter.New(app)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {

	var isRunningWithoutServerless = os.Getenv("IS_RUNNING_WO_SERVERLESS")

	if isRunningWithoutServerless == "true" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		logger.InitLogger()

		// stdout and stderr are sent to AWS CloudWatch Logs
		logger.Info.Printf("Starting Fiber without Serverless...")
		app := fiber.New()

		// Middlewares.
		middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
		routes.PublicRoutes(app)        // Register a public routes for app.
		routes.PrivateRoutes(app)       // Register a private routes for app.
		routes.NotFoundRoute(app)       // Register route for 404 Error.

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World! from Serverless Go Fiber App...")
		})

		log.Fatal(app.Listen(":8000"))

	} else {

		lambda.Start(Handler)

	}

}
