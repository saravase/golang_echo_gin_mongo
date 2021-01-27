package product_api

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var (
	e *echo.Echo
	v *validator.Validate
)

func init() {

	e = echo.New()
	v = validator.New()

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		e.Logger.Fatal("Unable to read configuration")
	}
	e.Logger.Printf("%+v\n", cfg)
}

// Middleware
func MiddlewareMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		e.Logger.Print("Message middleware triggered")
		return next(c)
	}
}

func Start() {

	// Initialize Middleware
	e.Use(MiddlewareMessage)

	// Endpoint GET : /products/:id
	e.GET("/products/:id", GetProductByID)

	// Endpoint GET : /products
	e.GET("/products", GetProducts)

	// Endpoint POST : /products
	e.POST("/products", CreateProduct)

	// End point PUT : /products/:id
	e.PUT("/products/:id", UpdateProductByID)

	// End point DELETE : /products/:id
	e.DELETE("/products/:id", DeleteProductByID)

	e.Logger.Printf("Listening on port %s", cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
