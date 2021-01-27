package product_api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var (
	port string
	e    *echo.Echo
	v    *validator.Validate
)

func Start() {

	port = os.Getenv("APP_PORT")
	if port == "" {
		port = "9090"
	}

	e = echo.New()
	v = validator.New()

	// Endpoint GET : /
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi Primz... ")
	})

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

	// End point GET : /query
	e.GET("/query", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.QueryParam("q"))
	})

	e.Logger.Printf("Listening on port %s", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
