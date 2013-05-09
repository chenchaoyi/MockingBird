package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"dummy"
	"sync"
)


// use “go run dummy_servers.go -pAmount 10 -lAmount 10” to run
// -pAmount means how many parent nodes to be created
// -lAmount means how many leaves you want to attach to each parent node
// node starts from 0 by default

// use localhost:48888/servers/#{server_num}/kill to kill presence node to of specific parent node
// use localhost:48888/servers/#{server_num}/revive to revive presence node to of specific parent node

var servers = make(map[int]*dummy.DummyStaticServer)

type RESTHandler func(*dummy.DummyStaticServer, *http.Request) (string, string)

func (self RESTHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var (
		result string
		error  string
	)
	ser_id, _ := strconv.Atoi(mux.Vars(r)["server"])

	if ser_id > 1000 {
		result, error = "", "serverid > 1000"
	} else {
		result, error = self(servers[ser_id], r)
	}

	w := rw.(io.Writer)

	rw.Header().Set("Content-Type", "application/json")

	if result != "" {
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode("{'error':'" + error + "'}")
	}
}

func killHandle(server *dummy.DummyStaticServer, r *http.Request) (string, string) {	
	var wg sync.WaitGroup
	log.Printf("removing all palantiri dummy servers")
	    	
	for i := 0; i < *parent_node_amount; i++ {
		wg.Add(1)
	    go func(index int){
			servers[index].Trigger_chan <- 4
			log.Printf(strconv.Itoa(index) + " deleted")
			wg.Done()
	    }(i)
	}
	wg.Wait()

	return "", ""
}

func handle(server *dummy.DummyStaticServer, r *http.Request) (string, string) {
	switch {
	// case strings.Contains(r.URL.String(), "deleteNode"):
	// 	// stop target server
	// 	server.Trigger_chan <- 0
	// case strings.Contains(r.URL.String(), "reviveNode"):
	// 	// start target server
	// 	server.Trigger_chan <- 1
	case strings.Contains(r.URL.String(), "kill"):
		// kill presence
		server.Trigger_chan <- 2
	case strings.Contains(r.URL.String(), "revive"):
		// revive presence
		server.Trigger_chan <- 3
	}

	return r.URL.String(), ""
}

var (
	parent_node_amount = flag.Int("pAmount", 10, "pAmount")
	leaf_node_amount = flag.Int("lAmount", 10, "lAmount")
)

// main func
func main() {
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < *parent_node_amount; i++ {
		wg.Add(1)
		go func(index int){
			s := new(dummy.DummyStaticServer)
			s.Init("http://localhost:3000", strconv.Itoa(index), strconv.Itoa(20000+index), *leaf_node_amount)
			s.LaunchAndListen()
			servers[index] = s
			wg.Done()
		}(i)
	}
	wg.Wait()

	r := mux.NewRouter()
	r.Handle("/kill", RESTHandler(killHandle)).Methods("GET")

	s := r.PathPrefix("/servers/{server:[0-9]+}").Subrouter()
	// s.Handle("/deleteNode", RESTHandler(handle)).Methods("GET")
	// s.Handle("/reviveNode", RESTHandler(handle)).Methods("GET")
	s.Handle("/kill", RESTHandler(handle)).Methods("GET")
	s.Handle("/revive", RESTHandler(handle)).Methods("GET")


	// r.Handle("/removeAll", RESTHandler(removeAll)).Methods("GET")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		r.ServeHTTP(w, req)
		duration := time.Since(startTime)
		log.Printf("%s %s %s %s %s", req.RemoteAddr, req.Method, req.URL, duration, req.Header.Get("X-Request-ID"))
	})

	if err := http.ListenAndServe(":48888", nil); err != nil {
		log.Fatal(err)
	}

}