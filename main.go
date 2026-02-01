package main

import (
	"database/sql"
	"encoding/json"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

func main() {
	_ = godotenv.Load()
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)
	defer db.Close()

	e := echo.New()

	e.GET("/", func(c *echo.Context) error {
		return c.String(200, "v1.0")
	})

	e.GET("/health", func(c *echo.Context) error {
		return c.String(200, "health ok")
	})

	e.POST("/api/login", func(c *echo.Context) error {
		var req LoginRequest
		err := json.NewDecoder(c.Request().Body).Decode(&req)
		if err != nil {
			return err
		}

		token, err := Login(db, req)
		if err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, map[string]any{
			"code":  200,
			"message": "Login successful",
			"data": map[string]string{
				"token": token,
			},
		})
	})

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
