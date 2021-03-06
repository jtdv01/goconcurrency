package barycenter

import (
	"os"
	"time"
	"fmt"
	"errors"
	"sync"
	"bufio"
)

// It takes a string to work on, a channel through which to send results, and a WaitGroup pointer.
// The first thing we need to do is make an async version of the loading procedure.
// It takes a string to work on, a channel through which to send results, and a WaitGroup pointer.
func stringToPointAsync(s string, c chan<- MassPoint, wg *sync.WaitGroup) {
	// First off, we'll defer the WaitGroup finishing operation, so however this function exits
	// it will notify the WaitGroup that it's done.
	defer wg.Done()
	// We'll create a new MassPoint to hold the result
	var newMassPoint MassPoint
	// Then we'll use Sscanf to parse the line
	_, err := fmt.Sscanf(s, "%f:%f:%f:%f", &newMassPoint.x, &newMassPoint.y, &newMassPoint.z, &newMassPoint.mass)
	// If there's an error, just abort
	if err != nil {
		return
	}
	// If there wasn't an error, send the result through the channel
	c <- newMassPoint
}

// Now we need an async version of the actual computation.
// This will be exactly the same as the avgMassPointsWeighted except that it takes a channel
// and passes the result through that channel.
// No need for fancy WaitGroup syncronization, since this can't fail.
func avgMassPointsWeightedAsync(a MassPoint, b MassPoint, c chan<- MassPoint) {
	c <- avgMassPointsWeighted(a, b)
}

func ParallelBarycenter(filename string){
	file, err := os.Open(filename)
	handle(err)
	defer closeFile(file)

	var masspoints []MassPoint

	startLoading := time.Now()

	// Now we need to modify the file parsing logic.
	// Rather than scanf, we'll use the function we created earlier, along with a buffered reader.
	r := bufio.NewReader(file)
	// We also need a buffered channel for results
	masspointsChan := make(chan MassPoint, 128)
	// And a waitgroup for syncronization
	var wg sync.WaitGroup
	for {
		// To actually get a line, we'll use the ReadString function
		str, err := r.ReadString('\n')
		// If the result is empty or there's an error, there are no more lines to read
		if len(str) == 0 || err != nil {
			break
		}

		// Otherwise, we'll start off a goroutine to parse the line
		wg.Add(1)
		go stringToPointAsync(str, masspointsChan, &wg)
	}

	// Now we'll set up syncronization. We need a channel for this, unbuffered.
	syncChan := make(chan bool)
	// Then we'll run a goroutine which will send a value through this channel when
	// the WaitGroup finishes.
	go func() { wg.Wait(); syncChan <- false }()

	// Finally,  we'll receive the values in a loop
	// We'll have a boolean value telling us if the computations are still running
	run := true
	// If they're still running, or there are values in the channel, keep receiving values
	for run || len(masspointsChan) > 0 {
		select {
		// If a value is available, we'll put it in the masspoints buffer
		case val := <-masspointsChan:
			masspoints = append(masspoints, val)
			// If the computations are done, we'll toggle the switch off
		case _ = <-syncChan:
			run = false
		}
	}

	fmt.Printf("Loaded %d values from file in %s.\n", len(masspoints), time.Since(startLoading))
	if len(masspoints) <= 1 {
		handle(errors.New("Insufficient number of values; there must be at least one "))
	}

	// Just before the processing loop, we'll create a channel.
	// It'll be buffered, and the larger the buffer, the faster the program will run,
	// up to half the size of the input.
	c := make(chan MassPoint, len(masspoints)/2)

	startCalculation := time.Now()
	for len(masspoints) > 1 {
		var newMasspoints []MassPoint
		// We need a new variable here to keep track of how many goroutines we've
		// spun off.
		goroutines := 0
		for i := 0; i < len(masspoints)-1; i += 2 {
			// Now, rather than doing the actual processing here, we'll just spin off a goroutine
			// for each pair of points.
			go avgMassPointsWeightedAsync(masspoints[i], masspoints[i+1], c)
			goroutines++
		}

		// Now that all the goroutines are running, we'll recieve from them in a loop.
		for i := 0; i < goroutines; i++ {
			newMasspoints = append(newMasspoints, <-c)
		}

		if len(masspoints)%2 != 0 {
			newMasspoints = append(newMasspoints, masspoints[len(masspoints)-1])
		}

		masspoints = newMasspoints
	}
	systemAverage := masspoints[0]

	fmt.Printf("System barycenter is at (%f, %f, %f) and the system's mass is %f.\n",
		systemAverage.x,
		systemAverage.y,
		systemAverage.z,
		systemAverage.mass)
	fmt.Printf("Calculation took %s.\n", time.Since(startCalculation))
}
