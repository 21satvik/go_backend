package routes

import (
	"github.com/21satvik/go_fiber_tut/database"
	"github.com/21satvik/go_fiber_tut/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	//not model User, but a serializer for User
	Id       uint   `json:"id" gorm:"primaryKey"`
	FirsName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseuser(userModel models.User) User {
	return User{
		Id:       userModel.Id,
		FirsName: userModel.FirstName,
		LastName: userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseuser(user)
	return c.Status(200).JSON(responseUser)
}
