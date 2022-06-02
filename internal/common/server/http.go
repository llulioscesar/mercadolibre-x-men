package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
)

func RunHTTPServer(createHandler func(router fiber.Router) fiber.Router) {
	//RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler)
	RunHTTPServerOnAddr(":8000", createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router fiber.Router) fiber.Router) {
	apiRouter := fiber.New()
	setMiddlewares(apiRouter)
	//rootRouter := apiRouter.Group("/api")
	createHandler(apiRouter)
	log.Println("Starting HTTP server")

	err := apiRouter.Listen(addr)
	if err != nil {
		panic("Unable to start HTTP server")
	}
}

func setMiddlewares(router *fiber.App) {
	router.Use(requestid.New())
	router.Use(logger.New(logger.Config{
		Format:   "${pid} ${locals:requestid} [${ip}]:${port}  ${status} - ${method} ${path}\n",
		TimeZone: "America/Bogota",
	}))
	router.Use(recover.New())
	addCorsMiddleware(router)
}

func addCorsMiddleware(router *fiber.App) {
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if len(allowedOrigins) == 0 {
		return
	}

	router.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     allowedOrigins,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-Token",
		AllowCredentials: true,
		ExposeHeaders:    "Links",
		MaxAge:           300,
	}))
}
