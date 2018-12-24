package intro

import (
	"fmt"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:15"))
}

func listenForever() {
	for {
		fmt.Println("Listening...")
	}

}

func writeMail1(wg *sync.WaitGroup) {
	printTime("Done writing mail #1")
	wg.Done()
}

func writeMail2(wg *sync.WaitGroup) {
	printTime("Done writing mail #2")
	wg.Done()
}

func MainParallelism() {
	var wg sync.WaitGroup
	wg.Add(1)

	go listenForever()

	go writeMail1(&wg)
	go writeMail2(&wg)

	wg.Wait()
}