package vulcand

// Server stores a vulcan server record
type Server struct {
	backendID string
	serverID  string
	url       string
}

// NewServer constructs a new Server from an string and string
func NewServer(backendID, serverID, url string) *Server {
	return &Server{backendID, serverID, url}
}

// BackendID returns the server's backendID id
func (server *Server) BackendID() string {
	return server.backendID
}

// ServerID returns the server's id
func (server *Server) ServerID() string {
	return server.serverID
}

// URL returns the server's URL
func (server *Server) URL() string {
	return server.url
}
