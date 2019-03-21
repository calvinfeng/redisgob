package cmd

import (
	"log"

	"github.com/calvinfeng/redisgob/payload"
	"github.com/calvinfeng/redisgob/queue"
	"github.com/spf13/cobra"
)

var Pull = &cobra.Command{
	Use:   "pull",
	Short: "continously pull some random cats from redis queue",
	RunE:  pull,
}

func pull(cmd *cobra.Command, args []string) error {
	cfg := queue.Config{
		RedisAddr: "localhost:6379",
		QueueName: "cats",
	}

	q := queue.NewFIFO(cfg)
	for {
		cat := &payload.Cat{}
		if err := q.Dequeue(cat); err != nil {
			log.Printf("failed to dequeue %v", err)
		} else {
			log.Printf("pulled %d years old cat %s from queue\n", cat.Age, cat.Name)
		}
	}
}
