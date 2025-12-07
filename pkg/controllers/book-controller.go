package controllers

import (
    "bookmanagementmysql/config"
    "bookmanagementmysql/models"
    "github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
    var books []models.Book
    config.DB.Find(&books)
    return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
    var book models.Book
    if err := c.BodyParser(&book); err != nil {
        return err
    }
    config.DB.Create(&book)
    return c.JSON(book)
}

func GetBookByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := config.DB.First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
    }
    return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := config.DB.First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
    }

    var data models.Book
    if err := c.BodyParser(&data); err != nil {
        return err
    }

    book.Title = data.Title
    book.Author = data.Author

    config.DB.Save(&book)
    return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
    id := c.Params("id")
    config.DB.Delete(&models.Book{}, id)
    return c.JSON(fiber.Map{"message": "Book deleted"})
}
