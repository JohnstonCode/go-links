package main

import (
	"github.com/JohnstonCode/go-links/model"
	"github.com/JohnstonCode/go-links/utils"
	"github.com/gofiber/fiber/v2"
)

func getAllLinks(ctx *fiber.Ctx) error {
	links, err := model.GetAllLinks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting all links",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(links)
}

func createLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var err error
	var link model.Link
	err = ctx.BodyParser(&link)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"messahe": "error parsing JSON" + err.Error(),
		})
	}

	link.Hash = utils.RandomString(5)

	err = model.CreateLink(link)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to create link",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(link)
}

func setupAndListen() {
	router := fiber.New()

	router.Get("/links", getAllLinks)
	router.Post("/links", createLink)

	router.Listen(":3000")
}

func main() {
	model.Setup()

	setupAndListen()
}
