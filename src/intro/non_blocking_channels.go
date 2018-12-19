package intro

import (
	"fmt"
	"math/rand"
	"time"
)

func cakeMaker(kind string, number int, to chan<- string) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < number; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		to <- kind
	}

	// once all cakes are made, we can close the channel
	// cannot send / receive from it anymore
	close(to)
}

func MainNonBlocking() {
	chocolateChan := make(chan string, 8)
	redVelvetChan := make(chan string, 8)

	go cakeMaker("chocolate", 4, chocolateChan)
	go cakeMaker("red velvet", 3, redVelvetChan)

	moreChocolate := true
	moreRedVelvet := true

	var cake string

	for moreChocolate || moreRedVelvet {
		select {
		case cake, moreChocolate = <- chocolateChan:
			// fmt.Printf("Are there more %s? %s\n", cake, moreChocolate)
			if moreChocolate {
				fmt.Printf("Got a cake from the first factory: %s\n", cake)
			}
		case cake, moreRedVelvet = <- redVelvetChan:
			// fmt.Printf("Are there more %s? %s\n", cake, moreRedVelvet)
			if moreRedVelvet {
				fmt.Printf("Got a cake from the second factory: %s\n", cake)
			}
		}
	}

}