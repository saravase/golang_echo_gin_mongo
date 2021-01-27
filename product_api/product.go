package product_api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var (
	products = []map[int]string{
		{1: "Mobile"},
		{2: "TV"},
		{3: "Laptop"},
	}
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func GetProductByID(c echo.Context) error {

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
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {

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
}

func UpdateProductByID(c echo.Context) error {

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
}

func DeleteProductByID(c echo.Context) error {

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

}
