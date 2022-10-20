package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userstatus"`
}

type Users []User

var users = Users{
	User{generateID(), "superman", "dani", "carnauba", "dani@dani.com", "secret", "88997773137", 0},
	User{generateID(), "mulhergato", "dani", "carnauba", "dani@dani.com", "secret", "88997773137", 0},
}

func generateID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10000)
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func postUser(c echo.Context) error {
	user := User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, "request não possui body")
	}

	user.Id = generateID()
	users = append(users, user)
	return c.JSON(201, "usuário criado")

}

func postUserArray(c echo.Context) error {
	arrayUsers := []User{}
	if err := c.Bind(&arrayUsers); err != nil {
		return c.JSON(400, err)
	}
	for _, user := range arrayUsers {
		users = append(users, user)
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

	updateUser := User{}
	err := c.Bind(&updateUser)
	if err != nil {
		return c.JSON(400, "request não possui body")
	}

	for i, _ := range users {
		if users[i].Username == username {
			users[i].Id = updateUser.Id
			users[i].Username = updateUser.Username
			users[i].FirstName = updateUser.FirstName
			users[i].LastName = updateUser.LastName
			users[i].Email = updateUser.Email
			users[i].Password = updateUser.Password
			users[i].Phone = updateUser.Phone
			users[i].UserStatus = updateUser.UserStatus

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
	for i, _ := range users {
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
