package barycenter

import (
	"os"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

func NaiveBarycenterDataset() {
	if len(os.Args) < 2{
		fmt.Println(fmt.Errorf("Not enough args!"))
		os.Exit(1)
	}

	nBodies, err := strconv.Atoi(os.Args[1])

	if err != nil {
		os.Exit(1)
	}

	rand.Seed(time.Now().Unix())

	posMax := 100
	massMax := 5

	for i := 0; i < nBodies; i++ {
		posX := rand.Intn(posMax * 2) - posMax
		posY := rand.Intn(posMax * 2) - posMax
		posZ := rand.Intn(posMax * 2) - posMax
		mass := rand.Intn(massMax - 1) + 1

		fmt.Printf("%d:%d:%d:%d\n",posX,posY,posZ,mass)
	}
}