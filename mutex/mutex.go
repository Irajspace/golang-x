package main

import(
	"fmt"
	"sync"
)

type post struct{
	views int
	mu sync.Mutex
}

func(p *post)inc(wg *sync.WaitGroup){
	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views++
	
}
func main() {

	var wg sync.WaitGroup
	p := post{views: 0}

	wg.Add(2)
	go p.inc(&wg)
	go p.inc(&wg)

	wg.Wait()
	fmt.Println("Total views:", p.views)	
	
}
