package vulcand

// Manager provides server management functions
// for vulcan
type Manager struct {
	uri string
}

// NewManager constructs a new manager
func NewManager(uri string) *Manager {
	return &Manager{uri}
}

// ServerRM removes the server from vulcan, using the vulcan API
func (manager *Manager) ServerRM(server *Server) error {
	return nil
}

// ShuffledServers returns all the servers from vulcand
// in random order
func (manager *Manager) ShuffledServers() ([]*Server, error) {
	var servers []*Server
	return servers, nil
}
