package routes

import (
	"time"

	"github.com/21satvik/go_fiber_tut/database"
	"github.com/21satvik/go_fiber_tut/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	Id        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		Id:        order.Id,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err := database.Database.Db.First(&user, order.UserRefer).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product

	if err := database.Database.Db.First(&product, order.ProductRefer).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Create(&order).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&order)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	database.Database.Db.Find(&orders)
	var responseOrders []Order
	for _, order := range orders {
		var user models.User
		database.Database.Db.First(&user, order.UserRefer)
		var product models.Product
		database.Database.Db.First(&product, order.ProductRefer)
		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)
		responseOrders = append(responseOrders, CreateResponseOrder(order, responseUser, responseProduct))
	}
	return c.Status(200).JSON(responseOrders)
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var order models.Order
	database.Database.Db.First(&order, id)
	if order.Id == 0 {
		return c.Status(404).JSON("Order not found!")
	}
	var user models.User
	database.Database.Db.First(&user, order.UserRefer)
	var product models.Product
	database.Database.Db.First(&product, order.ProductRefer)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}

func UpdateOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var order models.Order
	database.Database.Db.First(&order, id)
	if order.Id == 0 {
		return c.Status(404).JSON("Order not found!")
	}
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Save(&order)
	return c.Status(200).JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var order models.Order
	database.Database.Db.First(&order, id)
	if order.Id == 0 {
		return c.Status(404).JSON("Order not found!")
	}
	database.Database.Db.Delete(&order)
	return c.Status(200).JSON("Order deleted!")
}
