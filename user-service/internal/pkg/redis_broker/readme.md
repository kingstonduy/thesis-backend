```
package main

import (
	"context"
	"encoding/json"
	"fmt"
	errorx "go-playground/lib/error"
	redix "go-playground/lib/redis_pubsub_server"
	transport "go-playground/lib/transport/model"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	address  = "10.96.20.152:6379"
	password = ""
	database = 0
)

func main() {
	app := fiber.New()
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})

	app.Use(limiter.New(limiter.Config{
		Max:        10000,           // max count of connections
		Expiration: 1 * time.Second, // expiration time of the limit
		LimitReached: func(c *fiber.Ctx) error {
			return errorx.FailedErrorx(errorx.WithDetail("too many request"))
		},
	}))

	rApp := redix.New(rdb)

	go func() {
		rApp.Listen(context.Background())
	}()

	// waiting the redis to subcribe
	time.Sleep(time.Second * 1)

	cnt := 1
	app.Post("/redix", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		// init
		aggID := uuid.New().String()

		event := transport.Event[string]{
			AggregateID:   aggID,
			AggregateType: "agg type",
			EventID:       uuid.New().String(),
			EventType:     "event type",
			PayLoad:       fmt.Sprintf("%d", cnt),
			Trace: transport.Trace{
				Cid:  uuid.New().String(),
				Cts:  time.Now().UnixMilli(),
				Sid:  uuid.New().String(),
				Sts:  time.Now().UnixMilli(),
				From: "from",
				To:   "to",
			},
		}
		eventBytes, _ := json.Marshal(event)

		rApp.Publish(ctx, redix.RedisMessage{Key: aggID, Value: string(eventBytes)})

		val, err := rApp.GetValue(ctx, aggID, time.Second*10)
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(errorx.InternalServerErrorx(errorx.WithDetail(err.Error())))
		} else {
			c.Status(http.StatusOK).JSON(transport.ResponseType[string]{
				Result: transport.GetSuccessResult(),
				Data:   val,
			})
		}

		return nil
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

```