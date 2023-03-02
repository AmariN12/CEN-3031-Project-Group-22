package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisg0416/CEN-3031-Project-Group-22/Models"
)

var flashCards []Models.Card

func CreateFlashCard(c *fiber.Ctx) error {
	flashCard := &Models.Card{}
		
	if err := c.BodyParser(flashCard); err != nil {
		return err
	}

	flashCard.Id = len(flashCards) + 1

	flashCards = append(flashCards, *flashCard)

	//fmt.Println(flashCard)

	return c.JSON(flashCards)
}

func GetFlashCards(c *fiber.Ctx) error {
	return c.JSON(flashCards)
}

func GetFlashCardsID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		var card Models.Card
		for _, card := range flashCards {
			if card.Id == id {
				return c.JSON(card)
			}
		}

		return c.JSON(card)
} 