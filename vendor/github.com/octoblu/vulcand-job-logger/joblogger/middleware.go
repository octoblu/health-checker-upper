package joblogger

import (
	"fmt"
	"net/http"

	"github.com/octoblu/vulcand-job-logger/connection"
)

// Middleware is a vulcand middleware that logs to redis
type Middleware struct {
	BackendID, RedisURI, RedisQueueName string
	Conn                                *connection.Connection
}

// NewMiddleware constructs new Middleware instances
func NewMiddleware(RedisURI, RedisQueueName, BackendID string) (*Middleware, error) {
	if RedisURI == "" || RedisQueueName == "" || BackendID == "" {
		return nil, fmt.Errorf("RedisURI, RedisQueueName, and BackendID are all required. received '%v', '%v', and '%v'", RedisURI, RedisQueueName, BackendID)
	}

	Conn := connection.New(RedisURI, RedisQueueName)

	return &Middleware{BackendID, RedisURI, RedisQueueName, Conn}, nil
}

// NewHandler returns a new Handler instance
func (middleware *Middleware) NewHandler(next http.Handler) (http.Handler, error) {
	return NewHandler(middleware.Conn, middleware.BackendID, next), nil
}

// String will be called by loggers inside Vulcand and command line tool.
func (middleware *Middleware) String() string {
	return middleware.Conn.String()
}
