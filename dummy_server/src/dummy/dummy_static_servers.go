package dummy

import (
	"log"
	"strconv"
	// "strings"
	"git.gree-dev.net/stanislav-vishnevski/go-overseer"
)

// to attach multiple leaf nodes (amount a) to each node (amount b)
type DummyStaticServer struct {
	DummyServer

	leafNodeAmount int
	host           string
	watcher        *overseer.Presence
}

func (self *DummyStaticServer) Init(host string, id string, port string, a int) {
	self.status = 0
	self.server_id = id
	self.port = port
	self.Trigger_chan = make(chan int, 0)
	self.host = host
	self.leafNodeAmount = a

	self.palantiri_keeper, _ = overseer.New("palantiri", host)

}

func (self *DummyStaticServer) createAndWatch(host, key string) (*overseer.Presence, error) {
	log.Printf("creating node in %s/registry%s", host, key)
	self.palantiri_keeper.Create(key, nil)

	for j := 0; j < self.leafNodeAmount; j++ {
		self.palantiri_keeper.Create(key+"/leaf"+strconv.Itoa(j), nil)
	}
	// log.Printf("creating emphemerial node in %s/registry%s", host, key + "/presence")
	return self.palantiri_keeper.Presence(key + "/presence")
}

func (self *DummyStaticServer) die() {
	for j := 0; j < self.leafNodeAmount; j++ {
		self.palantiri_keeper.Delete("/QA/performance/parentNode" + self.server_id + "/leaf" + strconv.Itoa(j))
	}
	self.palantiri_keeper.Delete("/QA/performance/parentNode" + self.server_id)
}

func (self *DummyStaticServer) LaunchAndListen() {
	// server launches
	self.status = 0

	self.watcher, _ = self.createAndWatch(self.host, "/QA/performance/parentNode"+self.server_id)
	log.Println("presence loaded for node" + self.server_id)
	// go self.palantiri_keeper.Presence("/QA/performance/~server" + self.server_id)

	go func() {
		// do server kill/revive
		for {
			if command := <-self.Trigger_chan; command != self.status {
				switch command {
				// case 0:
				// 	// delete node
				// 	log.Printf("delete node" + self.server_id)
				// 	self.palantiri_keeper.Delete("/QA/performance/parentNode" + self.server_id)
				// 	// watcher.Die()
				// 	log.Printf("done")
				// case 1:
				// 	// revive node
				// 	log.Printf("reviving node" + self.server_id)
				// 	self.watcher, _ = self.createAndWatch(self.host, "/QA/performance/parentNode" + self.server_id)

				// 	// go self.palantiri_keeper.Presence("/QA/performance/~server" + self.server_id)
				// 	log.Printf("done")
				case 2:
					// kill presence
					log.Printf("killing present node" + self.server_id)
					self.watcher.Die()
					log.Printf("done")
				case 3:
					//revive presence
					log.Printf("creating present node in %s/registry%s", self.host, "/QA/performance/parentNode"+self.server_id+"/presence")
					self.watcher, _ = self.palantiri_keeper.Presence("/QA/performance/parentNode" + self.server_id + "/presence")
					log.Printf("done")
				case 4:
					log.Printf("deleting node" + self.server_id)
					self.watcher.Die()
					self.die()
				}
				self.status = command
			}
		}
	}()
}
