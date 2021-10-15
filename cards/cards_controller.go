package card

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/v1/cards", AddCard)
	e.GET("/v1/cards", FindAllCards)
}

func AddCard(c echo.Context) error {
	card := new(Card)
	if err := c.Bind(card); err != nil {
		return err
	}
	CardsList.Add(*card)
	return c.JSON(http.StatusCreated, card)
}

func FindAllCards(c echo.Context) error {
	return c.JSON(http.StatusOK, CardsList)
}
