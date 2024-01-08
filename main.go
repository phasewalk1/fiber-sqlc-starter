package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASS")

	conn, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s",
			dbHost, dbUser, dbPass, dbName,
		),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close(ctx)

	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(logger.New())
	app.Listen(":3000")
}
