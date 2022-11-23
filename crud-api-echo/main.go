package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/Danielecarn/crud-api-echo/models"
	"github.com/Danielecarn/crud-api-echo/resources"

	"github.com/labstack/echo/v4"
)

var users models.Users

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func postUser(c echo.Context) error {
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if !models.Validacao(user) {
		return c.JSON(400, "request não possui body")
	}

	user.Id = resources.GenerateID()
	users = resources.AddUser(users, user)
	return c.JSON(201, "usuário criado")

}

func postUserArray(c echo.Context) error {
	arrayUsers := []models.User{}
	if err := c.Bind(&arrayUsers); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if len(arrayUsers) == 0 {
		return c.JSON(400, "request não possui body")
	}

	for _, user := range arrayUsers {

		if !models.Validacao(user) {
			return c.JSON(400, "user inválido")
		}
		users = resources.AddUser(users, user)
	}
	return c.JSON(201, "Todos os usuários criados")

}

func getUser(c echo.Context) error {
	username := c.Param("username")

	matched, _ := regexp.MatchString(`\W.*`, username)
	if matched {
		return c.JSON(400, "username inválido")
	}

	for _, user := range users {
		if user.Username == username {
			return c.JSON(200, user)
		}
	}
	return c.JSON(404, "Usuário não encontrado")
}

func putUser(c echo.Context) error {
	username := c.Param("username")

	matched, _ := regexp.MatchString(`\W.*`, username)
	if matched {
		return c.JSON(400, "username inválido")
	}

	updateUser := models.User{}
	err := c.Bind(&updateUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if !models.Validacao(updateUser) {
		return c.JSON(400, "request não possui body")
	}
	for i := range users {
		if users[i].Username == username {
			resources.ChangeUser(&users[i], updateUser)
			return c.JSON(200, "usuário atualizado")
		}
	}
	return c.JSON(404, "usuário não encontrado")
}

func deleteUser(c echo.Context) error {
	username := c.Param("username")
	matched, _ := regexp.MatchString(`\W.*`, username)

	if matched {
		return c.JSON(400, "username inválido")
	}
	for i := range users {
		if users[i].Username == username {
			users = append(users[:i], users[i+1:]...)

			return c.JSON(200, "usuário apagado")
		}
	}
	return c.JSON(404, "usuário não encontrado")
}

func main() {

	fmt.Println("Server is running...")
	e := echo.New()
	e.GET("/users", getUsers)
	e.POST("/users", postUser)
	e.POST("/users/createWithArray", postUserArray)
	e.GET("/users/:username", getUser)
	e.PUT("/users/:username", putUser)
	e.DELETE("/users/:username", deleteUser)

	e.Start(":8080")
}
