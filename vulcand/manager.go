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
	client          Client
	cachedFrontends map[string]bool
}

// NewManager constructs a new manager
func NewManager(uri string) (Manager, error) {
	reg, err := registry.GetRegistry()
	if err != nil {
		return nil, err
	}

	client := NewClient(api.NewClient(uri, reg))
	return NewManagerWithClient(client), nil
}

// NewManagerWithClient constructs a new manager with a client
func NewManagerWithClient(client Client) Manager {
	return &HTTPManager{
		client:          client,
		cachedFrontends: make(map[string]bool),
	}
}

// ServerRm removes the server from vulcan, using the vulcan API
func (manager *HTTPManager) ServerRm(server *Server) error {
	return manager.client.DeleteServer(server.BackendID(), server.ServerID())
}

// Servers returns all the servers from vulcand
func (manager *HTTPManager) Servers() ([]*Server, error) {
	var allServers []*Server

	backends, err := manager.client.GetBackends()
	if err != nil {
		return allServers, err
	}

	for _, backend := range backends {
		hasFrontend, err := manager.hasFrontend(backend)
		if err != nil {
			return allServers, err
		}
		if !hasFrontend {
			continue
		}

		servers, err := manager.serversForBackend(backend)
		if err != nil {
			return allServers, err
		}

		allServers = append(allServers, servers...)
	}

	return allServers, nil
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

func (manager *HTTPManager) getFrontends() (map[string]bool, error) {
	if len(manager.cachedFrontends) > 0 {
		return manager.cachedFrontends, nil
	}

	frontends, err := manager.client.GetFrontends()
	if err != nil {
		return make(map[string]bool), err
	}

	for _, frontend := range frontends {
		manager.cachedFrontends[frontend] = true
	}

	return manager.cachedFrontends, nil
}

func (manager *HTTPManager) hasFrontend(backend string) (bool, error) {
	frontends, err := manager.getFrontends()
	if err != nil {
		return false, err
	}

	_, ok := frontends[backend]
	return ok, nil
}

// servers returns all servers for a particular backend
func (manager *HTTPManager) serversForBackend(backendID string) ([]*Server, error) {
	empty := []*Server{}

	serverIDs, err := manager.client.GetServers(backendID)
	if err != nil {
		return empty, err
	}

	servers := []*Server{}

	for _, serverID := range serverIDs {
		url, err := manager.client.GetServerURL(backendID, serverID)
		if err != nil {
			return empty, nil
		}

		servers = append(servers, NewServer(backendID, serverID, url))
	}

	return servers, nil
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
