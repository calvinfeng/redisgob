package queue

import (
	"bytes"
	"encoding/gob"

	"github.com/go-redis/redis"
)

// Config provides configuration for connection to Redis.
type Config struct {
	QueueName string
	RedisAddr string
}

// FIFO is a first in first out queue.
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

// Enqueue adds an element to Redis queue.
func (q *FIFO) Enqueue(val interface{}) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(val); err != nil {
		return err
	}

	len, err := q.cli.LPush(q.id, string(buf.Bytes())).Result()
	if err != nil {
		return err
	}

	q.len = len
	return nil
}

// Dequeue pops an element from Redis queue.
func (q *FIFO) Dequeue(val interface{}) error {
	elems, err := q.cli.BRPop(0, q.id).Result()
	if err != nil {
		return err
	}

	if len(elems) != 2 {
		panic("elemnts should have key and value")
	}

	return decode([]byte(elems[1]), val)
}

func decode(value []byte, result interface{}) error {
	buf := bytes.NewBuffer(value)
	enc := gob.NewDecoder(buf)
	err := enc.Decode(result)
	if err != nil {
		return err
	}

	return nil
}
