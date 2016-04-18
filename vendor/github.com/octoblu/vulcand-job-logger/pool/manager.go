package pool

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Manager manages different redis connections
type Manager struct {
	messageChan chan message
	connections map[string]*redis.Pool
}

// NewManager constructs a new Manager
func NewManager(messageChan chan message) *Manager {
	connections := make(map[string]*redis.Pool)
	return &Manager{messageChan, connections}
}

// Manage tells the manager to manage. I'm a
// people person
func (manager *Manager) Manage() {
	for {
		msg := <-manager.messageChan
		manager.message(msg)
	}
}

func (manager *Manager) message(msg message) {
	connection := manager.connection(msg.redisURI)
	defer connection.Close()

	_, err := connection.Do("LPUSH", msg.redisQueueName, msg.data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "LPUSH failed", err.Error())
	}
}

func (manager *Manager) connection(redisURI string) redis.Conn {
	pool := manager.pool(redisURI)
	return pool.Get()
}

func (manager *Manager) pool(redisURI string) *redis.Pool {
	pool, ok := manager.connections[redisURI]
	if ok {
		return pool
	}

	pool = newPool(redisURI)
	manager.connections[redisURI] = pool
	return pool
}

func newPool(redisURI string) *redis.Pool {
	fmt.Println("newPool", redisURI)
	return &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(redisURI)
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}
