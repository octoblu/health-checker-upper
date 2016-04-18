package joblogger

import (
	"net/http"
	"time"

	"github.com/octoblu/vulcand-job-logger/connection"
	"github.com/octoblu/vulcand-job-logger/wrapper"
)

// Handler implements http.Handler
type Handler struct {
	conn      *connection.Connection
	backendID string
	next      http.Handler
}

// NewHandler constructs a new handler
func NewHandler(conn *connection.Connection, backendID string, next http.Handler) *Handler {
	return &Handler{conn, backendID, next}
}

// ServeHTTP will be called each time the request
// hits the location with this middleware activated
func (handler *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	wrapped := wrapper.New(rw, time.Now(), handler.backendID, handler.conn.Publish)
	handler.next.ServeHTTP(wrapped, r)
}
