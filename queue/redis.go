package queue

import (
	"github.com/go-redis/redis"
)

type Config struct {
	QueueName string
	RedisAddr string
}

// FIFO implements first in first out queue backed by Redis.
type FIFO struct {
	id  string
	len int64
	cli *redis.Client
}

// NewFIFO returns a FIFO queue.
func NewFIFO(cfg Config) *FIFO {
	cli := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
	return &FIFO{cfg.QueueName, 0, cli}
}

// Enqueue puts an item into queue.
func (q *FIFO) Enqueue(value interface{}) error {
	len, err := q.cli.LPush(q.id, value).Result()
	if err != nil {
		return err
	}

	q.len = len
	return nil
}

// Write implements the io.Writer interface.
func (q *FIFO) Write(p []byte) (n int, err error) {
	err = q.Enqueue(p)
	if err != nil {
		return
	}

	return len(p), nil
}

// Dequeue takes an item out of the queue.
func (q *FIFO) Dequeue() (string, error) {
	elems, err := q.cli.BRPop(0, q.id).Result()
	if err != nil {
		return "", err
	}

	if len(elems) != 2 {
		panic("elemnts should have key and value")
	}

	return elems[1], nil
}

func (q *FIFO) Read(p []byte) (n int, err error) {
	var val string
	val, err = q.Dequeue()
	if err != nil {
		return
	}

	for i, b := range []byte(val) {
		p[i] = b
	}

	return len([]byte(val)), nil
}
