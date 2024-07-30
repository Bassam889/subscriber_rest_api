package controllers

import (
	"app/initializers"
	"app/models"

	"github.com/gofiber/fiber/v2"
)

func GetSubscriber(c *fiber.Ctx) error {
	var subscribers []models.Subscriber
	initializers.DB.Find(&subscribers)
	if len(subscribers) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error",
			"message": "Users not found", "data": nil})
	}
	if subscribers != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": nil})
	}
	return c.Status(fiber.StatusOK).JSON(subscribers)
}

func CreateSubscriber(c *fiber.Ctx) error {
	subscriber := new(models.Subscriber)
	err := c.BodyParser(subscriber)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": nil})
	}
	initializers.DB.Create(&subscriber)
	return c.Status(fiber.StatusCreated).JSON(subscriber)
}

func GetSubscriberById(c *fiber.Ctx) error {
	id := c.Params("id")

	var subscriber models.Subscriber

	result := initializers.DB.Find(&subscriber, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error",
			"message": "User not found", "data": nil})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"stauts": "success",
		"message": "User found", "data": subscriber})
}

func UpdateSubscriber(c *fiber.Ctx) error {
	type UpdateSubscriber struct {
		Name                string `json:"name"`
		SubscribedToChannel string `json:"subscribedToChannel"`
	}
	id := c.Params("id")
	var subscriber models.Subscriber
	result := initializers.DB.Find(&subscriber, id)

	var updateSubscriber UpdateSubscriber
	err := c.BodyParser(&updateSubscriber)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
			"message": "Something's wrong with your input", "data": err})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error",
			"message": "User not found", "data": nil})
	}

	subscriber.Name = updateSubscriber.Name
	subscriber.SubscribedToChannel = updateSubscriber.SubscribedToChannel

	initializers.DB.Save(&subscriber)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success",
		"message": "User's found", "data": subscriber})
}

func DeleteSubscriber(c *fiber.Ctx) error {
	id := c.Params("id")
	var subscriber models.Subscriber
	result := initializers.DB.Find(&subscriber, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error",
			"message": "User not found", "data": nil})
	}
	err := initializers.DB.Delete(&subscriber, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success",
		"message": "User deleted", "data": subscriber})
}
