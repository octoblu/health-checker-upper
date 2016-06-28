package vulcand

import (
	"math/rand"
	"strings"

	"github.com/octoblu/vulcand-bundle/registry"
	"github.com/vulcand/vulcand/api"
	"github.com/vulcand/vulcand/engine"
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
	client          *api.Client
	cachedFrontends map[string]bool
}

// NewManager constructs a new manager
func NewManager(uri string) (Manager, error) {
	reg, err := registry.GetRegistry()
	if err != nil {
		return nil, err
	}

	return &HTTPManager{
		client:          api.NewClient(uri, reg),
		cachedFrontends: make(map[string]bool),
	}, nil
}

// ServerRm removes the server from vulcan, using the vulcan API
func (manager *HTTPManager) ServerRm(server *Server) error {
	backendKey := engine.BackendKey{Id: server.BackendID()}
	serverKey := engine.ServerKey{BackendKey: backendKey, Id: server.ServerID()}

	err := manager.client.DeleteServer(serverKey)

	if manager.isKeyNotFoundError(err) {
		return nil
	}

	return err
}

// Servers returns all the servers from vulcand
func (manager *HTTPManager) Servers() ([]*Server, error) {
	var allServers []*Server

	manager.client.GetBackends()
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
	if manager.isKeyNotFoundError(err) {
		return make(map[string]bool), nil
	}
	if err != nil {
		return make(map[string]bool), err
	}

	for _, frontend := range frontends {
		manager.cachedFrontends[frontend.BackendId] = true
	}

	return manager.cachedFrontends, nil
}

func (manager *HTTPManager) hasFrontend(backend engine.Backend) (bool, error) {
	frontends, err := manager.getFrontends()
	if err != nil {
		return false, err
	}

	_, ok := frontends[backend.GetId()]
	return ok, nil
}

// servers returns all servers for a particular backend
func (manager *HTTPManager) serversForBackend(backend engine.Backend) ([]*Server, error) {
	var servers []*Server

	vctlServers, err := manager.client.GetServers(backend.GetUniqueId())
	if manager.isKeyNotFoundError(err) {
		return make([]*Server, 0), nil
	}
	if err != nil {
		return make([]*Server, 0), err
	}

	for _, vctlServer := range vctlServers {
		servers = append(servers, NewServer(backend, vctlServer))
	}

	return servers, nil
}

func (manager *HTTPManager) isKeyNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	errorMessage := err.Error()
	return strings.HasPrefix(errorMessage, "Key not found")
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
