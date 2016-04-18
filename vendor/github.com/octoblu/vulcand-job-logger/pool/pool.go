package pool

import "sync"

// Pool is a redis connection pool
type Pool struct {
	redisChannel     chan message
	redisChannelOnce sync.Once
}

// message contains a message to be published to the pool
type message struct {
	redisURI       string
	redisQueueName string
	data           []byte
}

// New constructs a new Pool
func New() *Pool {
	return &Pool{}
}

// Publish publishes to redis
func (pool *Pool) Publish(redisURI, redisQueueName string, data []byte) {
	channel := pool.channel()
	channel <- message{redisURI, redisQueueName, data}
}

func (pool *Pool) channel() chan message {
	pool.redisChannelOnce.Do(func() {
		pool.redisChannel = make(chan message)
		manager := NewManager(pool.redisChannel)
		go manager.Manage()
	})

	return pool.redisChannel
}
