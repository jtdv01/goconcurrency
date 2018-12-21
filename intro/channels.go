package intro

import "fmt"

func lucasoid(a, b, n int) int {
	if n == 0{
		return a
	}
	if n == 1{
		return b
	}
	return lucasoid(a, b, n-1) + lucasoid(a, b, n -2)
}

func MainChannels() {
	/*
	Race but always expect 1 to finish fist
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() { ch1 <- lucasoid(0,1,20)}()
	go func() { ch2 <- lucasoid(0,1,30)}()
	go func() { ch3 <- lucasoid(0,1,40)}()

	select {
		case msg := <- ch1:
			fmt.Printf("First finisher: ch1 returned %d\n", msg)
		case msg := <- ch2:
			fmt.Printf("First finisher: ch2 returned %d\n", msg)
		case msg := <- ch3:
			fmt.Printf("First finisher: ch3 returned %d\n", msg)
	}

}

func MainSingleChannel() {
	/*
	Just having one channel instead
	Think of it as a pipe. Can forgo waitgroups
	*/
	ch := make(chan int)

	go func() { ch <- lucasoid(0,1,20)}()
	go func() { ch <- lucasoid(0,1,30)}()
	go func() { ch <- lucasoid(0,1,40)}()

	fmt.Printf("First finisher: %d\n", <- ch)
}