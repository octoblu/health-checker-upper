package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/octoblu/health-checker-upper/vulcand"
)

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

	return response.StatusCode == 200
}
