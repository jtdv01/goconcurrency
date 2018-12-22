package barycenter

import (
	"math/rand"
	"time"
	"fmt"
)

func GenerateBarycenterDatasets (nBodies int) {


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


