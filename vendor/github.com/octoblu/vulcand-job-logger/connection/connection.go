package connection

import (
	"fmt"
	"sync"
)

import "github.com/octoblu/vulcand-job-logger/pool"

var redisPool *pool.Pool
var redisPoolOnce sync.Once

// Connection connects to redis and
type Connection struct {
	redisURI, redisQueueName string
}

// New constructs a new Connection
func New(redisURI, redisQueueName string) *Connection {
	redisPoolOnce.Do(func() {
		redisPool = pool.New()
	})
	return &Connection{redisURI, redisQueueName}
}

// Publish puts the thing in the redis queue
func (connection *Connection) Publish(data []byte) {
	go redisPool.Publish(connection.redisURI, connection.redisQueueName, data)
}

// String will be called by loggers inside Vulcand and command line tool.
func (connection *Connection) String() string {
	return fmt.Sprintf("redis-uri=%v, redis-queue-name=%v", connection.redisURI, connection.redisQueueName)
}
