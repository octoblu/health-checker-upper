package health

import "github.com/octoblu/health-checker-upper/vulcand"

// Check returns true if the uri responds with an HTTP 200
// status code, false otherwise. The healthcheck URI is expected
// to be the uri on the passed in server with '/healthcheck'
// appended to it
func Check(server *vulcand.Server) (bool, error) {
	return false, nil
}
