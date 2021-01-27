package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var (
	port     string
	products []map[int]string
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func init() {
	port = os.Getenv("APP_PORT")
	if port == "" {
		port = "9090"
	}

	products = []map[int]string{
		{1: "Mobile"},
		{2: "TV"},
		{3: "Laptop"},
	}
}

func main() {

	e := echo.New()
	v := validator.New()

	// Endpoint GET : /
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi Primz... ")
	})

	// Endpoint GET : /products/:id
	e.GET("/products/:id", func(c echo.Context) error {

		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		for _, p := range products {
			for k := range p {
				if pID == k {
					product = p
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, product)
	})

	// Endpoint GET : /products
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	// Endpoint POST : /products
	e.POST("/products", func(c echo.Context) error {

		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		reqBody := body{}
		e.Validator = &ProductValidator{
			validator: v,
		}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}

		if err := c.Validate(reqBody); err != nil {
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)

		return c.JSON(http.StatusOK, product)
	})

	// End point PUT : /products/:id
	e.PUT("/products/:id", func(c echo.Context) error {

		var product map[int]string

		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		for _, p := range products {
			for k := range p {
				if k == pID {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}

		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		reqBody := body{}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}

		e.Validator = &ProductValidator{validator: v}
		if err := c.Validate(reqBody); err != nil {
			return err
		}

		product[pID] = reqBody.Name
		return c.JSON(http.StatusOK, product)
	})

	// End point DELETE : /products/:id
	e.DELETE("/products/:id", func(c echo.Context) error {

		var product map[int]string
		var index int

		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		for i, p := range products {
			for k := range p {
				if k == pID {
					product = p
					index = i
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}

		splice := func(m []map[int]string, index int) []map[int]string {
			return append(m[:index], m[index+1:]...)
		}

		products = splice(products, index)

		return c.JSON(http.StatusOK, product)

	})

	// End point GET : /query
	e.GET("/query", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.QueryParam("q"))
	})

	e.Logger.Printf("Listening on port %s", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))

}
