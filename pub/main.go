package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

const messageChannel = "send-user-data"

var ctx = context.Background()

func main() {

	e.GET("/", articleIndex)

	e.POST("/publish", func(c echo.Context) error {
		var user User

		if err := c.Bind(&user); err != nil {
			return err
		}

		payload, err := json.Marshal(user)
		if err != nil {
			return err
		}

		if err := rdb.Publish(ctx, messageChannel, payload).Err(); err != nil {
			return nil
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Pub!")
}
