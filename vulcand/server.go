package vulcand

import "github.com/vulcand/vulcand/engine"

// Server stores a vulcan server record
type Server struct {
	backend engine.Backend
	server  engine.Server
}

// NewServer constructs a new Server from an engine.Backend and engine.Server
func NewServer(backend engine.Backend, server engine.Server) *Server {
	return &Server{backend, server}
}

// BackendID returns the server's backend id
func (server *Server) BackendID() string {
	return server.backend.GetId()
}

// ServerID returns the server's id
func (server *Server) ServerID() string {
	return server.server.GetId()
}

// URL returns the server's URL
func (server *Server) URL() string {
	return server.server.URL
}
