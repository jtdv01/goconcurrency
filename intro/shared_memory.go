package intro

import (
	"fmt"
	"sync"
)

func printEven(x int, wg *sync.WaitGroup){
	if x%2 == 0{
		fmt.Printf("%d is even\n", x)
	}
}

func increment(ptr *int, wg *sync.WaitGroup){
	/**
	Bad!
	**/
	i := *ptr

	fmt.Printf("i is %d\n", i)
	*ptr = i + 1
	wg.Done()
}

func MainSharedMemory(){
	var wg sync.WaitGroup

	for i:=0; i < 10; i++ {
		wg.Add(1)
		// pass a pointer to the wait group
		go printEven(i, &wg)
		go increment(&i, &wg)
	}
	wg.Wait()
}