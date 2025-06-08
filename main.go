package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users []user

func main() {
	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})

	//e.GET("/users/:user_id", func(c echo.Context) error {
	//	return c.JSON(http.StatusOK, users)
	//})

	//e.PUT("/users/:usr_id", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})

	e.POST("/users", func(c echo.Context) error {
		request := new(user)
		if err := c.Bind(request); err != nil {
			return err
		}

		for _, value := range users {
			if value.ID == request.ID {
				err := errors.New("id already exists")
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": err.Error(),
				})
			}
		}

		users = append(users, *request)
		return c.JSON(http.StatusCreated, request)
	})

	e.DELETE("/users/:user_id", func(c echo.Context) error {
		idStr := c.Param("user_id")
		if idStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "id must no be null",
			})
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		for index, value := range users {
			if value.ID == id {
				users = append(users[:index], users[index+1:]...)
				return c.JSON(http.StatusOK, "user removed")
			}
		}

		return c.String(http.StatusBadRequest, "Not Found")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
