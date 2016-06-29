package vulcand

import "github.com/mailgun/vulcand/engine"

// Client defines the interface of the underlying vulcan client
type Client interface {
	DeleteServer(backendID, serverID string) error
	GetBackends() ([]string, error)
	GetFrontends() ([]string, error)
	GetServers(backendID string) ([]string, error)
}

// WrapperClient is an implementation of client that wraps
// an instance of github.com/vulcand/vulcand/api Client
type WrapperClient struct {
	wrapped WrappedClient
}

// WrappedClient defines the interface of the client the
// WrapperClient wraps
type WrappedClient interface {
	DeleteServer(sk engine.ServerKey) error
	GetBackends() ([]engine.Backend, error)
	GetFrontends() ([]engine.Frontend, error)
	GetServers(bk engine.BackendKey) ([]engine.Server, error)
}

// NewClient wrapps an API client in a simpler interface
func NewClient(client WrappedClient) Client {
	return &WrapperClient{client}
}

// DeleteServer deletes a server from vulcand
func (client *WrapperClient) DeleteServer(backendID, serverID string) error {
	backendKey := engine.BackendKey{Id: backendID}
	serverKey := engine.ServerKey{BackendKey: backendKey, Id: serverID}
	return client.wrapped.DeleteServer(serverKey)
}

// GetBackends returns the backends
func (client *WrapperClient) GetBackends() ([]string, error) {
	engineBackends, err := client.wrapped.GetBackends()
	if err != nil {
		return []string{}, err
	}

	backends := make([]string, len(engineBackends))
	for i, engineBackend := range engineBackends {
		backends[i] = engineBackend.GetId()
	}
	return backends, nil
}

// GetFrontends returns the backends
func (client *WrapperClient) GetFrontends() ([]string, error) {
	engineFrontends, err := client.wrapped.GetFrontends()
	if err != nil {
		return []string{}, err
	}

	frontends := make([]string, len(engineFrontends))
	for i, engineFrontend := range engineFrontends {
		frontends[i] = engineFrontend.GetId()
	}
	return frontends, nil
}

// GetServers returns servers for a backendID
func (client *WrapperClient) GetServers(backendID string) ([]string, error) {
	backendKey := engine.BackendKey{Id: backendID}
	engineServers, err := client.wrapped.GetServers(backendKey)
	if err != nil {
		return []string{}, err
	}

	servers := make([]string, len(engineServers))
	for i, engineServer := range engineServers {
		servers[i] = engineServer.GetId()
	}
	return servers, nil
}
