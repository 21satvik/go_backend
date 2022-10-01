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

func CreateResponseUser(userModel models.User) User {
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
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)
	var responseUsers []User
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}
	return c.Status(200).JSON(responseUsers)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	database.Database.Db.First(&user, id)
	if user.Id == 0 {
		return c.Status(404).JSON("User not found!")
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	database.Database.Db.First(&user, id)
	if user.Id == 0 {
		return c.Status(404).JSON("User not found!")
	}
	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser
	err = c.BodyParser(&updateData)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	database.Database.Db.First(&user, id)
	if user.Id == 0 {
		return c.Status(404).JSON("User not found!")
	}
	database.Database.Db.Delete(&user)
	return c.Status(200).JSON("User deleted!")
}
