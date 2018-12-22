package barycenter

import (
	"os"
	"time"
	"io"
	"fmt"
	"errors"
)

type MassPoint struct {
	x, y, z, mass int
}

func addMassPoints(a, b MassPoint) MassPoint {
	return MassPoint {
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
		a.mass + b.mass}
}

func avgMassPoints(a, b MassPoint) MassPoint {
	sum := addMassPoints(a, b)
	return MassPoint{
		sum.x / 2,
		sum.y / 2,
		sum.z / 2,
		sum.mass}
}

func toWeightedSubspace(a MassPoint) MassPoint {
	return MassPoint {
		a.x * a.mass,
		a.y * a.mass,
		a.z * a.mass,
		a.mass}
}

func fromWeightedSubspace(a MassPoint) MassPoint {
	return MassPoint{
		a.x / a.mass,
		a.y / a.mass,
		a.z / a.mass,
		a.mass}
}

func avgMassPointsWeighted(a, b MassPoint) MassPoint {
	aWeighted := toWeightedSubspace(a)
	bWeighted := toWeightedSubspace(b)
	return fromWeightedSubspace(avgMassPoints(aWeighted, bWeighted))
}

func handle(err error){
	if err != nil {
		panic(err)
	}
}

func closeFile(fi *os.File){
	err := fi.Close()
	handle(err)
}

func NaiveBarycenter(filename string){
	file, err := os.Open(filename)
	handle(err)
	defer closeFile(file)

	var masspoints []MassPoint

	startLoading := time.Now()
	for {
		var newMassPoint MassPoint
		_, err := fmt.Fscanf(file, "%d:%d:%d:%d", &newMassPoint.x, &newMassPoint.y, &newMassPoint.z, &newMassPoint.mass)
		if err == io.EOF{
			break
		} else if err != nil {
			continue
		}

		masspoints = append(masspoints, newMassPoint)
	}

	fmt.Printf("Loaded %d values from file in %s.\n", len(masspoints), time.Since(startLoading))
	if len(masspoints) <= 1{
		handle(errors.New("Insufficient values."))
	}

	startCalculation := time.Now()
	var systemAverage MassPoint
	for len(masspoints) > 1 {
		var newMasspoints []MassPoint
		for i := 0; i < len(masspoints) - 1; i += 2 {
			newMasspoints = append(newMasspoints, avgMassPointsWeighted(masspoints[i], masspoints[i+1]))
		}

		// Then we check to make sure we didn't leave off one
		if len(masspoints) % 2 != 0 {
			newMasspoints = append(newMasspoints, masspoints[len(masspoints) - 1])
		}

		masspoints = newMasspoints
	}
	systemAverage = masspoints[0]
	fmt.Println("System average: ", systemAverage)
	fmt.Printf("System barycenter is at (%d, %d, %d) and the system's mass is %d.\n",
	systemAverage.x,
	systemAverage.y,
	systemAverage.z,
	systemAverage.mass)
	// Finally, we just want to print out the time the calculation has taken.
	fmt.Printf("Calculation took %s.\n", time.Since(startCalculation))
}
