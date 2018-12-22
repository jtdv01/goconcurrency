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
func stringToPointAsync(s string, c chan<- MassPoint, wg *sync.WaitGroup){
	defer wg.Done()
	var newMassPoint MassPoint

	_, err := fmt.Sscanf(s, "%f:%f:%f:%f", &newMassPoint.x, &newMassPoint.y, &newMassPoint.z, &newMassPoint.mass)
	if err != nil {
		return
	}

	fmt.Println(newMassPoint)

	// If there wasn't an error, send the result through the channel
	c <- newMassPoint
}

func avgMassPointAsync(a, b MassPoint, c chan<- MassPoint){
	c <- avgMassPointsWeighted(a,b)
}

func ParallelBarycenter(filename string){
	file, err := os.Open(filename)
	handle(err)
	defer closeFile(file)

	var masspoints []MassPoint

	startLoading := time.Now()

	// Async reader
	r := bufio.NewReader(file)
	masspointsChan := make(chan MassPoint, 128)
	var wg sync.WaitGroup

	for {
		// To actually get a line, we'll use the ReadString function
		str, err := r.ReadString('\n')
		// Stop if error
		if len(str) == 0 || err != nil {
			break
		}
		// Otherwise, we'll start off a goroutine to parse the line
		wg.Add(1)
		go stringToPointAsync(str, masspointsChan, &wg)
	}

	// Now we'll set up syncronization. We need a channel for this, unbuffered.
	// Unbuffered channel, will be blocked if we try to read from it
	syncChan := make(chan bool)
	// Then we'll run a goroutine which will send a value through this channel when
	// the WaitGroup finishes.
	go func() {wg.Wait(); syncChan <- false}()
	run := true

	// While computations or try to drain out the channel
	for run || len(masspointsChan) > 0 {
		select {
		// If a value is available, we'll put it in the masspoints buffer
		case value := <-masspointsChan:
			masspoints = append(masspoints, value)
			// If the computations are done, we'll toggle the switch off
		case _ = <-syncChan:
			run = false
		}
	}

	fmt.Printf("Loaded %d values from file in %s.\n", len(masspoints), time.Since(startLoading))
	if len(masspoints) <= 1{
		handle(errors.New("Insufficient values."))
	}

	// Could be a large memory allocated
	// In prod, you'd porbably add a limit to this buffer
	c := make(chan MassPoint, len(masspoints)/2)

	startCalculation := time.Now()
	var systemAverage MassPoint
	for len(masspoints) > 1 {
		var newMasspoints []MassPoint

		goroutines := 0

		for i := 0; i < len(masspoints) - 1; i += 2 {
			go avgMassPointAsync(masspoints[i], masspoints[i+1], c)
			goroutines++
		}

		for i := 0; i < goroutines; i++{
			newMasspoints = append(newMasspoints, <-c)
		}


		// Then we check to make sure we didn't leave off one
		if len(masspoints) % 2 != 0 {
			newMasspoints = append(newMasspoints, masspoints[len(masspoints) - 1])
		}

		masspoints = newMasspoints
	}
	systemAverage = masspoints[0]
	fmt.Println("System average: ", systemAverage)
	fmt.Printf("System barycenter is at (%f, %f, %f) and the system's mass is %f.\n",
	systemAverage.x,
	systemAverage.y,
	systemAverage.z,
	systemAverage.mass)
	// Finally, we just want to print out the time the calculation has taken.
	fmt.Printf("Calculation took %s.\n", time.Since(startCalculation))
}
