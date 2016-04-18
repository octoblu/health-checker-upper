package vulcand

import "github.com/vulcand/vulcand/engine"

// Server stores a vulcan server record
type Server struct {
	ID, URL string
}

// NewServer constructs a new Server
func NewServer() *Server {
	return &Server{}
}

// newServerFromVCTL constructs a new Server instance
// from a vctl instance
func newServerFromVCTL(server engine.Server) *Server {
	return &Server{ID: server.GetId(), URL: server.URL}
}
