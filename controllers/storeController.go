package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"storeApiRest/models"
	"storeApiRest/services"
	"strconv"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	s := e.Group("/api")
	s.GET("/users", getUsers)
	s.POST("/users", postUser)
	s.PUT("/users", putUser)
	s.GET("/products", getProducts)
	s.POST("/products", postProduct)
	s.PUT("/products", putProduct)
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}

func getUsers(c echo.Context) error {
	users, err := services.ReadUsersService()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func postUser(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	if err := services.CreateUserService(user); err != nil {
		return err
	}
	return c.String(http.StatusCreated, "User created successfully")
}

func putUser(c echo.Context) error {
	param := c.QueryParam("id")
	userID, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := services.UpdateUserService(user, userID); err != nil {
		return err
	}
	return c.String(http.StatusOK, "User updated successfully")
}

func getProducts(c echo.Context) error {
	products, err := services.ReadProductsService()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func postProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := services.CreateProductService(product); err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Product created successfully")
}

func putProduct(c echo.Context) error {
	var product models.Product
	param := c.QueryParam("id")
	productID, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := services.UpdateProductService(product, productID); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Product updated successfully")
}