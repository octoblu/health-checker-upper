package vulcand_test

import "github.com/mailgun/vulcand/engine"

type FakeApiClient struct {
	DeleteServerCallCount      int
	DeleteServerLastCalledWith engine.ServerKey
	DeleteServerReturns        error

	GetBackendsCallCount       int
	GetBackendsReturnsBackends []engine.Backend
	GetBackendsReturnsError    error

	GetFrontendsCallCount        int
	GetFrontendsReturnsFrontends []engine.Frontend
	GetFrontendsReturnsError     error

	GetServersCallCount      int
	GetServersLastCalledWith engine.BackendKey
	GetServersReturnsServers []engine.Server
	GetServersReturnsError   error
}

func NewFakeApiClient() *FakeApiClient {
	return &FakeApiClient{DeleteServerCallCount: 0, GetBackendsCallCount: 0}
}

func (client *FakeApiClient) DeleteServer(serverKey engine.ServerKey) error {
	client.DeleteServerCallCount++
	client.DeleteServerLastCalledWith = serverKey
	return client.DeleteServerReturns
}

func (client *FakeApiClient) GetBackends() ([]engine.Backend, error) {
	client.GetBackendsCallCount++
	return client.GetBackendsReturnsBackends, client.GetBackendsReturnsError
}

func (client *FakeApiClient) GetFrontends() ([]engine.Frontend, error) {
	client.GetFrontendsCallCount++
	return client.GetFrontendsReturnsFrontends, client.GetFrontendsReturnsError
}

func (client *FakeApiClient) GetServers(backendKey engine.BackendKey) ([]engine.Server, error) {
	client.GetServersCallCount++
	client.GetServersLastCalledWith = backendKey
	return client.GetServersReturnsServers, client.GetServersReturnsError
}
