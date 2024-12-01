package main

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "list-keeper-backend/docs" // Если вы добавили пакет со сгенерированными swagger-доками
	"list-keeper-backend/handler"
	"list-keeper-backend/router"
	"list-keeper-backend/store"

	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// @title Shopping List API
// @version 1.0
// @description API сервиса для управления списками покупок.
// @host localhost:4000
// @BasePath /

func main() {
	// Получаем переменные окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // значение по умолчанию, если переменная не установлена
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Инициализируем базу данных
	db, err := store.InitDatabase(databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	h := &handler.Handler{DB: db}

	e := echo.New()

	// Настроим маршрутизацию
	router.SetupRoutes(e, h)

	// Добавление роутов для Swagger

	// Настраиваем маршрутизацию
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Запускаем сервер
	e.Logger.Fatal(e.Start(":" + port))
}
