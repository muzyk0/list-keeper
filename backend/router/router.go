package router

import (
	"list-keeper-backend/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, h *handler.Handler) {
	// Настроим маршруты для списка покупок
	e.POST("/shopping-lists", h.CreateShoppingList)
	e.GET("/shopping-lists", h.GetShoppingLists)
	e.GET("/shopping-lists/:id", h.GetShoppingList)
	e.PUT("/shopping-lists/:id", h.UpdateShoppingList)
	e.DELETE("/shopping-lists/:id", h.DeleteShoppingList)
	return
}
