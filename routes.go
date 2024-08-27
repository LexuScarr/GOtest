package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/edit/:Id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("Id"))
		if err != nil {
			log.Printf("Ошибка: неверный ID %v", c.Params("Id"))
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}

		log.Printf("POST /edit/%d - изменение новости", id)

		var input struct {
			Title      string `json:"Title"`
			Content    string `json:"Content"`
			Categories []uint `json:"Categories"`
		}
		if err := c.BodyParser(&input); err != nil {
			log.Printf("Ошибка парсинга JSON: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		var news News
		if err := db.Preload("Categories").First(&news, id).Error; err != nil {
			log.Printf("Новость с ID %d не найдена", id)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
		}

		if input.Title != "" {
			news.Title = input.Title
		}
		if input.Content != "" {
			news.Content = input.Content
		}

		if len(input.Categories) > 0 {
			var categories []Category
			if err := db.Find(&categories, input.Categories).Error; err != nil {
				log.Printf("Ошибка при поиске категорий: %v", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Categories not found"})
			}
			news.Categories = categories
		}

		if result := db.Save(&news); result.Error != nil {
			log.Printf("Ошибка при сохранении новости с ID %d: %v", id, result.Error)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
		}

		log.Printf("Новость с ID %d успешно обновлена", id)
		return c.JSON(news)
	})

	app.Get("/list", func(c *fiber.Ctx) error {
		log.Println("GET /list - получение списка новостей")

		var news []News
		if err := db.Preload("Categories").Find(&news).Error; err != nil {
			log.Printf("Ошибка при получении списка новостей: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при получении новостей"})
		}

		var response []map[string]interface{}
		for _, n := range news {
			item := map[string]interface{}{
				"Id":         n.ID,
				"Title":      n.Title,
				"Content":    n.Content,
				"Categories": []uint{},
			}

			for _, cat := range n.Categories {
				item["Categories"] = append(item["Categories"].([]uint), cat.ID)
			}

			response = append(response, item)
		}

		log.Println("Список новостей успешно получен")
		return c.JSON(fiber.Map{"Success": true, "News": response})
	})
}
