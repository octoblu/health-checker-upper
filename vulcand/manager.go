package vulcand

import (
	"math/rand"

	"github.com/octoblu/vulcand-bundle/registry"
	"github.com/vulcand/vulcand/api"
)

// Manager provides server management functions
// for vulcan
type Manager interface {
	ShuffledServers() ([]*Server, error)
	ServerRm(*Server) error
}

// HTTPManager implements manager over Vulcan's HTTP
// API
type HTTPManager struct {
	client *api.Client
}

// NewManager constructs a new manager
func NewManager(uri string) (Manager, error) {
	reg, err := registry.GetRegistry()
	if err != nil {
		return nil, err
	}

	client := api.NewClient(uri, reg)
	return &HTTPManager{client}, nil
}

// ServerRm removes the server from vulcan, using the vulcan API
func (manager *HTTPManager) ServerRm(server *Server) error {
	return nil
}

// Servers returns all the servers from vulcand
func (manager *HTTPManager) Servers() ([]*Server, error) {
	var servers []*Server
	return servers, nil
}

// ShuffledServers returns all the servers from vulcand
// in random order
func (manager *HTTPManager) ShuffledServers() ([]*Server, error) {
	servers, err := manager.Servers()
	if err != nil {
		return servers, err
	}

	return manager.shuffle(servers), nil
}

// shuffle shuffles the servers in place, then returns the altered
// slice
func (manager *HTTPManager) shuffle(servers []*Server) []*Server {
	for i := range servers {
		j := rand.Intn(i + 1)
		servers[i], servers[j] = servers[j], servers[i]
	}

	return servers
}
