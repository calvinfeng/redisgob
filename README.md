# Inter-process Communication

## via Direct Network

Launch the server in one process

    redisgob serve

Dial the server in another process

    redisgob dial

A cat will be passed from client to server. The communication is inter-process.

```text
2019/02/14 16:26:41 server received cat Ripperthread Spur which is 8 years old
2019/02/14 16:26:43 server received cat Takerroot Bee which is 6 years old
2019/02/14 16:26:43 server received cat Koalafast Eye which is 2 years old
2019/02/14 16:26:44 server received cat Chopperpetal Shirt which is 19 years old
2019/02/14 16:26:44 server received cat Moosenight Lifter which is 8 years old
2019/02/14 16:26:44 server received cat Soareriridescent Frill which is 5 years old
2019/02/14 16:26:45 server received cat Catcherruby Witch which is 12 years old
2019/02/14 16:26:45 server received cat Browbush Leg which is 8 years old
2019/02/14 16:26:45 server received cat Whimseysavage Kangaroo which is 5 years old
2019/02/14 16:29:51 server received cat Flamephase Koala which is 17 years old
2019/02/14 16:29:52 server received cat Houndballistic Stealer which is 4 years old
2019/02/14 16:29:53 server received cat Llamasavage Tiger which is 0 years old
2019/02/14 16:29:53 server received cat Venomevening Jackal which is 9 years old
2019/02/14 16:29:54 server received cat Healerheavy Snake which is 7 years old
2019/02/14 16:29:54 server received cat Sparrowchisel Face which is 10 years old
2019/02/14 16:29:54 server received cat Robingiant Drop which is 14 years old
```

## via Redis

Push the data to Redis queue in one process

    redisgob push

Pull the data from Redis queue in another process

    redisgob pull

The queue itself implements `io.Reader` and `io.Writer`. We can pass the queue to the `gob` encoder
and decoder.

```golang
cfg := queue.Config{
    RedisAddr: "localhost:6379",
    QueueName: "cats",
}

q := queue.NewFIFO(cfg)
encoder := gob.NewEncoder(q)
```

```golang
cfg := queue.Config{
    RedisAddr: "localhost:6379",
    QueueName: "cats",
}

q := queue.NewFIFO(cfg)
decoder := gob.Decoder(q)
```