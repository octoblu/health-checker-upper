package vulcand_test

import (
	"fmt"

	vulcand "github.com/octoblu/health-checker-upper/vulcand"
	"github.com/vulcand/vulcand/engine"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var wrappedClient *FakeApiClient
	var sut vulcand.Client

	BeforeEach(func() {
		wrappedClient = NewFakeApiClient()
		sut = vulcand.NewClient(wrappedClient)
	})

	Describe("NewClient", func() {
		It("should create a new client instance", func() {
			Expect(sut).NotTo(BeNil())
		})
	})

	Describe("client.DeleteServer", func() {
		Describe("When the wrapped client returns an error", func() {
			var err error

			BeforeEach(func() {
				wrappedClient.DeleteServerReturns = fmt.Errorf("No, I'm blind")
				err = sut.DeleteServer("hello", "is it me you're looking for?")
			})

			It("Should call wrapped.DeleteServer with an ServerKey", func() {
				var serverKey engine.ServerKey
				serverKey = wrappedClient.DeleteServerLastCalledWith

				Expect(wrappedClient.DeleteServerCallCount).To(Equal(1))
				Expect(serverKey.BackendKey.Id).To(Equal("hello"))
				Expect(serverKey.Id).To(Equal("is it me you're looking for?"))
			})

			It("Should return the error value", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("No, I'm blind"))
			})
		})

		Describe("When the wrapped client returns a '100: Key not found' error", func() {
			var err error

			BeforeEach(func() {
				wrappedClient.DeleteServerReturns = fmt.Errorf("100: Key not found (/foo/bar)")
				err = sut.DeleteServer("hello", "is it me you're looking for?")
			})

			It("Should call wrapped.DeleteServer with an ServerKey", func() {
				var serverKey engine.ServerKey
				serverKey = wrappedClient.DeleteServerLastCalledWith

				Expect(wrappedClient.DeleteServerCallCount).To(Equal(1))
				Expect(serverKey.BackendKey.Id).To(Equal("hello"))
				Expect(serverKey.Id).To(Equal("is it me you're looking for?"))
			})

			It("Should no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("client.GetBackends", func() {
		Describe("when the wrapped client returns some backends", func() {
			var backends []string
			var err error

			BeforeEach(func() {
				backend := engine.Backend{Id: "Foo"}
				wrappedClient.GetBackendsReturnsBackends = []engine.Backend{backend}
				backends, err = sut.GetBackends()
			})

			It("should call wrappedClient.GetBackends", func() {
				Expect(wrappedClient.GetBackendsCallCount).To(Equal(1))
			})

			It("should return the backends", func() {
				Expect(backends).To(Equal([]string{"Foo"}))
			})

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("when the wrapped client returns an error", func() {
			var backends []string
			var err error

			BeforeEach(func() {
				wrappedClient.GetBackendsReturnsError = fmt.Errorf("Baby got no back(end)")
				backends, err = sut.GetBackends()
			})

			It("should call wrappedClient.GetBackends", func() {
				Expect(wrappedClient.GetBackendsCallCount).To(Equal(1))
			})

			It("should return an empty array", func() {
				Expect(backends).To(HaveLen(0))
			})

			It("should return the error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Baby got no back(end)"))
			})
		})
	})

	Describe("client.GetFrontends", func() {
		Describe("when the wrapped client returns some frontends", func() {
			var frontends []string
			var err error

			BeforeEach(func() {
				frontend := engine.Frontend{Id: "Foo"}
				wrappedClient.GetFrontendsReturnsFrontends = []engine.Frontend{frontend}
				frontends, err = sut.GetFrontends()
			})

			It("should call wrappedClient.GetFrontends", func() {
				Expect(wrappedClient.GetFrontendsCallCount).To(Equal(1))
			})

			It("should return the frontends", func() {
				Expect(frontends).To(Equal([]string{"Foo"}))
			})

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("when the wrapped client returns an error", func() {
			var frontends []string
			var err error

			BeforeEach(func() {
				wrappedClient.GetFrontendsReturnsError = fmt.Errorf("Baby got no back(end)")
				frontends, err = sut.GetFrontends()
			})

			It("should call wrappedClient.GetFrontends", func() {
				Expect(wrappedClient.GetFrontendsCallCount).To(Equal(1))
			})

			It("should return an empty array", func() {
				Expect(frontends).To(HaveLen(0))
			})

			It("should return the error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Baby got no back(end)"))
			})
		})

		Describe("when the wrapped client returns a '100: Key not found' error", func() {
			var frontends []string
			var err error

			BeforeEach(func() {
				wrappedClient.GetFrontendsReturnsError = fmt.Errorf("100: Key not found (/some/path)")
				frontends, err = sut.GetFrontends()
			})

			It("should call wrappedClient.GetFrontends", func() {
				Expect(wrappedClient.GetFrontendsCallCount).To(Equal(1))
			})

			It("should return an empty array", func() {
				Expect(frontends).To(HaveLen(0))
			})

			It("should return no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("client.GetServers", func() {
		Describe("when the wrapped client returns some servers", func() {
			var servers []string
			var err error

			BeforeEach(func() {
				server := engine.Server{Id: "um"}
				wrappedClient.GetServersReturnsServers = []engine.Server{server}
				servers, err = sut.GetServers("jup")
			})

			It("should call wrappedClient.GetServers with the backend key", func() {
				backendKey := wrappedClient.GetServersLastCalledWith

				Expect(wrappedClient.GetServersCallCount).To(Equal(1))
				Expect(backendKey.Id).To(Equal("jup"))
			})

			It("should return the servers", func() {
				Expect(servers).To(Equal([]string{"um"}))
			})

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("when the wrapped client returns an error", func() {
			var servers []string
			var err error

			BeforeEach(func() {
				wrappedClient.GetServersReturnsError = fmt.Errorf("You got served")
				servers, err = sut.GetServers("meiru")
			})

			It("should call wrappedClient.GetServers", func() {
				Expect(wrappedClient.GetServersCallCount).To(Equal(1))
				Expect(wrappedClient.GetServersLastCalledWith.Id).To(Equal("meiru"))
			})

			It("should return an empty array", func() {
				Expect(servers).To(HaveLen(0))
			})

			It("should return the error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("You got served"))
			})
		})

		Describe("when the wrapped client returns an '100: Key not found'", func() {
			var servers []string
			var err error

			BeforeEach(func() {
				wrappedClient.GetServersReturnsError = fmt.Errorf("100: Key not found (/vulcand/something)")
				servers, err = sut.GetServers("meiru")
			})

			It("should call wrappedClient.GetServers", func() {
				Expect(wrappedClient.GetServersCallCount).To(Equal(1))
				Expect(wrappedClient.GetServersLastCalledWith.Id).To(Equal("meiru"))
			})

			It("should return an empty array", func() {
				Expect(servers).To(HaveLen(0))
			})

			It("should not return the error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("GetServerURL", func() {
		Describe("When wrappedClient.GetServer returns an engine.Server", func() {
			var url string
			var err error

			BeforeEach(func() {
				wrappedClient.GetServerReturnsServer = &engine.Server{Id: "server", URL: "http://biz.apps"}
				url, err = sut.GetServerURL("backend", "server")
			})

			It("Should call wrappedClient.GetServer", func() {
				serverKey := wrappedClient.GetServerLastCalledWith
				Expect(wrappedClient.GetServerCallCount).To(Equal(1))
				Expect(serverKey.BackendKey.Id).To(Equal("backend"))
				Expect(serverKey.Id).To(Equal("server"))
			})

			It("Should return the url", func() {
				Expect(url).To(Equal("http://biz.apps"))
			})

			It("Should return no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
