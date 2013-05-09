package dummy

import (
	"encoding/json"
	"io"
	// "fmt"
	// "net"
	"net/http"
	// "strings"
	"git.gree-dev.net/stanislav-vishnevski/go-overseer"
)

type ServerCoreHandler func()

func (self ServerCoreHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	w := rw.(io.Writer)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r.URL.String())
}

type DummyServer struct {
	server_id    string
	status       int
	port         string
	Trigger_chan chan int

	palantiri_keeper overseer.Interface
	server_handler   ServerCoreHandler
}
