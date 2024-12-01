package handler

import (
	"gorm.io/gorm"
	"list-keeper-backend/model"
	"list-keeper-backend/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateShoppingList Создать списки покупок
// @Summary Создать список покупок
// @Description Создает новый список покупок
// @Tags shopping-lists
// @Accept json
// @Produce json
// @Param shoppingList body model.ApiCreateShoppingList true "Список покупок"
// @Success 201 {object} model.ShoppingList
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /shopping-lists [post]
func (h *Handler) CreateShoppingList(c echo.Context) error {
	list := new(model.ShoppingList)
	if err := c.Bind(list); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "Invalid data"})
	}
	if err := h.DB.Create(&list).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Failed to create shopping list"})
	}
	return c.JSON(http.StatusCreated, list)
}

// GetShoppingLists Получить списки покупок
// @Summary Возвращает все списки покупок
// @Description Загружает все списки покупок вместе с их элементами
// @Tags shopping-lists
// @Produce json
// @Success 200 {array} model.ShoppingList "Списки покупок успешно получены"
// @Failure 500 {object} responses.ErrorResponse "Не удалось получить списки покупок"
// @Router /shopping-lists [get]
func (h *Handler) GetShoppingLists(c echo.Context) error {
	var lists []model.ShoppingList
	if err := h.DB.Preload("Items").Order("created_at desc").Find(&lists).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Failed to fetch shopping lists"})
	}
	return c.JSON(http.StatusOK, lists)
}

// GetShoppingList Получить список покупок по ID
// @Summary Возвращает список покупок по указанному ID
// @Description Загружает список покупок вместе с его элементами, используя предоставленный ID
// @Tags shopping-lists
// @Param id path int true "ID списка покупок"
// @Produce json
// @Success 200 {object} model.ShoppingList "Список покупок успешно получен"
// @Failure 404 {object} responses.ErrorResponse "Список покупок не найден"
// @Router /shopping-lists/{id} [get]
func (h *Handler) GetShoppingList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	list := new(model.ShoppingList)
	if err := h.DB.Preload("Items").First(&list, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "Shopping list not found"})
	}
	return c.JSON(http.StatusOK, list)
}

// UpdateShoppingList Обновить список покупок
// @Summary Обновляет существующий список покупок по ID
// @Description Обновляет данные списка покупок, используя предоставленный ID
// @Tags shopping-lists
// @Param id path int true "ID списка покупок"
// @Param list body model.ShoppingList true "Данные списка покупок для обновления"
// @Produce json
// @Success 200 {object} model.ShoppingList "Список покупок успешно обновлен"
// @Failure 400 {object} responses.ErrorResponse "Недопустимые данные"
// @Failure 404 {object} responses.ErrorResponse "Список покупок не найден"
// @Failure 500 {object} responses.ErrorResponse "Ошибка обновления списка покупок"
// @Router /shopping-lists/{id} [put]
func (h *Handler) UpdateShoppingList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	list := new(model.ShoppingList)
	if err := h.DB.First(&list, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "Shopping list not found"})
	}

	if err := c.Bind(list); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "Invalid data"})
	}
	if err := h.DB.Save(&list).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Failed to update shopping list"})
	}
	return c.JSON(http.StatusOK, list)
}

// DeleteShoppingList Удалить список покупок
// @Summary Удаляет существующий список покупок по ID
// @Description Удаляет список покупок, используя предоставленный ID
// @Tags shopping-lists
// @Param id path int true "ID списка покупок"
// @Produce json
// @Success 200 {object} responses.ErrorResponse "Список покупок успешно удален"
// @Failure 500 {object} responses.ErrorResponse "Ошибка удаления списка покупок"
// @Router /shopping-lists/{id} [delete]
func (h *Handler) DeleteShoppingList(c echo.Context) error {
	// Пытаемся преобразовать параметр в целое число и проверяем на ошибку
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "Invalid shopping list ID"})
	}

	// Проводим транзакцию
	err = h.DB.Transaction(func(tx *gorm.DB) error {
		// Удаляем связанные элементы
		if err := tx.Delete(&model.ShoppingItem{}, "shopping_list_id = ?", id).Error; err != nil {
			return err
		}

		// Удаляем саму запись списка покупок
		if err := tx.Delete(&model.ShoppingList{}, id).Error; err != nil {
			return err
		}

		return nil
	})

	// Проверяем результат транзакции
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Failed to delete shopping list"})
	}

	// Возвращаем успешный ответ
	return c.JSON(http.StatusOK, responses.ErrorResponse{Message: "success"})
}
