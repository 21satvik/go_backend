package routes

import (
	"github.com/21satvik/go_fiber_tut/database"
	"github.com/21satvik/go_fiber_tut/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	SerialNo string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{
		Id:       productModel.Id,
		Name:     productModel.Name,
		SerialNo: productModel.SerialNo,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.Database.Db.Find(&products)
	var responseProducts []Product
	for _, product := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(product))
	}
	return c.Status(200).JSON(responseProducts)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	database.Database.Db.First(&product, id)
	if product.Id == 0 {
		return c.Status(404).JSON("Product not found!")
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	database.Database.Db.First(&product, id)
	if product.Id == 0 {
		return c.Status(404).JSON("Product not found!")
	}
	err = c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Save(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	database.Database.Db.First(&product, id)
	if product.Id == 0 {
		return c.Status(404).JSON("Product not found!")
	}
	database.Database.Db.Delete(&product)
	return c.Status(200).JSON("Product deleted successfully!")
}
