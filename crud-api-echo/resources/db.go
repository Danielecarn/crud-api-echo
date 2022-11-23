package resources

import (
	"math/rand"
	"time"

	"github.com/Danielecarn/crud-api-echo/models"
)

func AddUser(users []models.User, user models.User) models.Users {
	users = append(users, user)
	return users
}

func GenerateID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10000)
}

func ChangeUser(user *models.User, updateUser models.User) {
	user.Id = updateUser.Id
	user.Username = updateUser.Username
	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName
	user.Email = updateUser.Email
	user.Password = updateUser.Password
	user.Phone = updateUser.Phone
	user.UserStatus = updateUser.UserStatus

}
