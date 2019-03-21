package cmd

import (
	"log"
	"math/rand"
	"time"

	sillyname "github.com/Pallinder/sillyname-go"
	"github.com/calvinfeng/redisgob/payload"
	"github.com/calvinfeng/redisgob/queue"
	"github.com/spf13/cobra"
)

var Push = &cobra.Command{
	Use:   "push",
	Short: "continuously push some random cats onto redis queue",
	RunE:  push,
}

func push(cmd *cobra.Command, args []string) error {
	cfg := queue.Config{
		RedisAddr: "localhost:6379",
		QueueName: "cats",
	}

	q := queue.NewFIFO(cfg)
	for {
		cat := &payload.Cat{
			Name: sillyname.GenerateStupidName(),
			Age:  rand.Intn(20),
		}

		if err := q.Enqueue(cat); err != nil {
			log.Printf("failed to push to redis, %v", err)
		} else {
			log.Printf("pushed %d years old cat %s to queue", cat.Age, cat.Name)
		}

		time.Sleep(1 * time.Second)
	}
}
