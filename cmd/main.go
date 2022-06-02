package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/llulioscesar/mercadolibre-x-men/internal/mutant"
	"log"
)

func main() {
	apiRouter := fiber.New()
	apiRouter.Use(requestid.New())
	apiRouter.Use(logger.New(logger.Config{
		Format:   "${pid} ${locals:requestid} [${ip}]:${port}  ${status} - ${method} ${path}\n",
		TimeZone: "America/Bogota",
	}))
	apiRouter.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-Token",
		AllowCredentials: true,
		ExposeHeaders:    "Links",
		MaxAge:           300,
	}))
	apiRouter.Use(recover.New())

	mutant.NewMutantRouter(apiRouter)

	log.Println("Starting HTTP server")

	err := apiRouter.Listen(":8000")
	if err != nil {
		panic("Unable to start HTTP server")
	}
}
