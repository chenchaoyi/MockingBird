package main

import (
	"git.gree-dev.net/stanislav-vishnevski/go-overseer"
	"flag"
	"strconv"
	"sync"
)

var (
	parent_node_amount = flag.Int("pAmount", 10, "pAmount")
	leaf_node_amount = flag.Int("lAmount", 10, "lAmount")
	palantiri_keeper, _ = overseer.New("palantiri", "http://localhost:3000")
)

// main func
func main() {
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < *parent_node_amount; i++ {
		wg.Add(1)
		go func(index int){
			for j := 0; j < *leaf_node_amount; j++ {
				palantiri_keeper.Delete("/QA/performance/parentNode" + strconv.Itoa(index) + "/leaf" + strconv.Itoa(j))
			}
			palantiri_keeper.Delete("/QA/performance/parentNode" + strconv.Itoa(index))
			
			wg.Done()
		}(i)
	}
	wg.Wait()
}