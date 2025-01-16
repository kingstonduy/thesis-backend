package redix

// import (
// 	"context"
// 	"encoding/json"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"testing"
// 	"time"

// 	"git.ocb.vn/ktthud/microservices/library/golang/mcs-go-core.git/logger"
// 	"git.ocb.vn/ktthud/microservices/library/golang/mcs-go-core.git/transport"
// 	"github.com/gammazero/workerpool"
// 	"github.com/google/uuid"
// 	"github.com/redis/go-redis/v9"
// )

// const (
// 	address  = "10.96.20.152:6379"
// 	password = ""
// 	database = 0
// )

// func TestBasic(t *testing.T) {
// 	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
// 	defer stop()

// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     address,
// 		Password: password,
// 		DB:       database,
// 	})

// 	redix := New(rdb)

// 	wg := sync.WaitGroup{}

// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()
// 		Sub(ctx, redix)
// 	}()

// 	// waiting for the redis to init subscribe
// 	time.Sleep(time.Second * 2)

// 	Pub(ctx, redix)

// 	// waiting for the redis to process
// 	time.Sleep(time.Second * 2)
// }

// func Pub(ctx context.Context, redix PubSubBroker) {
// 	// init
// 	aggID := uuid.New().String()

// 	event := transport.Command{
// 		AggregateID:   aggID,
// 		AggregateType: "agg type",
// 		CommandID:     uuid.New().String(),
// 		CommandType:   "event type",
// 		Payload:       "hello",
// 		Trace: transport.Trace{
// 			Cid:  uuid.New().String(),
// 			Cts:  time.Now().UnixMilli(),
// 			Sid:  uuid.New().String(),
// 			Sts:  time.Now().UnixMilli(),
// 			From: "from",
// 			To:   "to",
// 		},
// 	}
// 	eventBytes, _ := json.Marshal(event)

// 	redix.Publish(ctx, RedisMessage{Key: aggID, Value: string(eventBytes)})
// }

// func Sub(ctx context.Context, redix PubSubBroker) {
// 	srvErr := make(chan error, 2)
// 	wp := workerpool.New(1)

// 	wp.Submit(func() {
// 		srvErr <- redix.Listen(ctx)
// 	})

// 	select {
// 	case err := <-srvErr:
// 		logger.Info(ctx, err)
// 		return
// 	case <-ctx.Done():
// 		// Wait for first CTRL+C. Stop receiving signal notifications as soon as possible.
// 		srvErr <- redix.Shutdown(ctx)
// 		return
// 	}
// }
