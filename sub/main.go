package main

import (
	"context"
	"encoding/json"
	"fmt"
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

type SubscribeParam struct {
	Channels []string `json:"channels"`
}

var rdb = redis.NewClient(&redis.Options{
	Addr:        "redis:6379",
	Password:    "", // no password set
	DB:          0,  // use default DB
	ReadTimeout: -1, // for verification
})

var ctx = context.Background()

func main() {
	e.GET("/", articleIndex)
	e.POST("/subscribe", subscribe)

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
	return c.String(http.StatusOK, "Hello, Sub!")
}

func subscribe(c echo.Context) error {
	var param SubscribeParam

	if err := c.Bind(&param); err != nil {
		return err
	}

	pubsub := rdb.Subscribe(ctx, param.Channels...)
	defer pubsub.Close()

	if _, err := pubsub.Receive(ctx); err != nil {
		panic(err)
	}

	ch := pubsub.Channel()

	for msg := range ch {
		// TODO: Unsubscribe

		user := User{}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", user)
	}

	return nil
}
