package intro

/*
	Channels are like pipes that move messages around

	Buffered channels can hold some messages in them unless they're fully blocked.
	Sending is non blocking, and unless they are empty, receiving is non-blocking
*/

import (
	"sync"
	"net/http"
	"fmt"
)

/**
	Channel can only receive in
**/
func webGetWorker(in <- chan string, wg *sync.WaitGroup) {
	for {
		url := <- in
		res, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("GET %s: %d\n", url, res.StatusCode)
		}

		// within a for loop, only the unit of work is done
		wg.Done()
	}

}

func MainBufferedChannels() {

	// Only x amount of items can sit inside the string
	// Most of the time it will be non blocking
	work := make(chan string, 1024)

	var wg sync.WaitGroup

	// Its cheap to add more workers!
	numWorkers := 1000

	for i := 0; i < numWorkers; i++ {
		// multiple workers
		go webGetWorker(work, &wg)
	}

	urls := [6]string{"http://example.com", "http://google.com", "http://reddit.com", "http://twitter.com", "http://bing.com", "http://apple.com"}

	for i := 0; i < 100; i++{
		for _, url := range urls {
			wg.Add(1)
			work <- url
		}
	}

	// wait until all work is done
	wg.Wait()
}