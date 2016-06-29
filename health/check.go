package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/octoblu/health-checker-upper/vulcand"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("health-checker-upper/health:check")

// Check returns true if the uri responds with an HTTP 200
// status code, false otherwise. The healthcheck URI is expected
// to be the uri on the passed in server with '/healthcheck'
// appended to it
func Check(server *vulcand.Server) bool {
	uri := fmt.Sprintf("%s/healthcheck", server.URL())
	client := &http.Client{Timeout: time.Second * 1}
	response, err := client.Get(uri)

	if err != nil {
		return false
	}

	debug("status: %v, uri: %v", response.StatusCode, uri)
	return response.StatusCode == 200
}
