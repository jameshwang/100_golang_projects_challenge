package tmp

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var NUM_OF_WORKERS int = 10

func worker(q <-chan string) {
	wg.Add(1)
	defer wg.Done()
	for {
		work, ok := <-q
		if !ok {
			fmt.Println("Worker finished. Going home :)")
			break
		}
		time.Sleep(time.Second)
		fmt.Println("Work on: ", work)
	}
}

func main() {
	queue := make(chan string)

	for i := 0; i < NUM_OF_WORKERS; i++ {
		go worker(queue)
	}

	queue <- "homework"
	queue <- "watch tv"
	queue <- "eat"

	close(queue)

	wg.Wait()
}
