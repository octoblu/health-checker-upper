package wrapper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/octoblu/go-logentry/logentry"
)

// Wrapper wraps in silence
type Wrapper struct {
	rw          http.ResponseWriter
	onStatus    OnStatus
	startTime   time.Time
	backendName string
}

// OnStatus is called whenever WriteHeader is called on the response writer
type OnStatus func([]byte)

// New constructs a new Wrapper instance
func New(rw http.ResponseWriter, startTime time.Time, backendName string, onStatus OnStatus) *Wrapper {
	return &Wrapper{rw, onStatus, startTime, backendName}
}

// Header returns the header map that will be sent by
// WriteHeader
func (wrapper *Wrapper) Header() http.Header {
	return wrapper.rw.Header()
}

// Write writes the data to the connection as part of an HTTP reply.
func (wrapper *Wrapper) Write(data []byte) (int, error) {
	return wrapper.rw.Write(data)
}

// WriteHeader sends an HTTP response header with status code.
func (wrapper *Wrapper) WriteHeader(statusCode int) {
	wrapper.logTheEntry(statusCode)
	wrapper.rw.WriteHeader(statusCode)
}

func (wrapper *Wrapper) logTheEntry(statusCode int) {
	elapsedTimeNano := time.Now().UnixNano() - wrapper.startTime.UnixNano()
	elapsedTime := int(elapsedTimeNano / 1000000)

	logEntry := logentry.New("metric:vulcand", "http", wrapper.backendName, "anonymous", statusCode, elapsedTime)
	logEntryBytes, err := json.Marshal(logEntry)
	logError("NewJobLogger failed: %v\n", err)

	if err != nil {
		return
	}

	wrapper.onStatus(logEntryBytes)
}

func logError(fmtMessage string, err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, fmtMessage, err.Error())
}
