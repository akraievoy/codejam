package algo

import (
	"testing"
	"math/rand"
	"fmt"
	"time"
	"math"
)

//	http://www.diss.fu-berlin.de/docs/servlets/MCRFileNodeServlet/FUDOCS_derivate_000000000270/1994_12.pdf
func TestMinCut1(t *testing.T) {
	mcNet := CapacityMatrixNew(8)

	mcNet.AddLink(0,1, 2)
	mcNet.AddLink(0,4, 3)
	mcNet.AddLink(1,2, 3)
	mcNet.AddLink(1,4, 2)
	mcNet.AddLink(1,5, 2)
	mcNet.AddLink(2,3, 4)
	mcNet.AddLink(2,6, 2)
	mcNet.AddLink(3,6, 2)
	mcNet.AddLink(3,7, 2)
	mcNet.AddLink(4,5, 3)
	mcNet.AddLink(5,6, 1)
	mcNet.AddLink(6,7, 3)

	{
		actual := mcNet.MinCutStoerWagner()
		expected := float64(4)

		if actual != expected {
			t.Errorf("mcNet.MinCutStoerWagner() is not correct: expected %v, actual %v", expected, actual)
		}
	}
}

func TestRandomProbes(t *testing.T) {
	randomProbe(3, 5, t)
	randomProbe(4, 5, t)
	randomProbe(5, 5, t)
	randomProbe(6, 5, t)
	randomProbe(7, 5, t)
	randomProbe(8, 5, t)
	randomProbe(12, 5, t)
	randomProbe(16, 5, t)
	randomProbe(24, 5, t)
	randomProbe(36, 5, t)
	randomProbe(50, 5, t)
}

func randomProbe(size uint32, timeoutSeconds float64, t *testing.T) {
	started := time.Now()
	tests := 0
	for time.Since(started).Seconds() < timeoutSeconds {
		seed := (65536+tests)*101*103 + 107*109;
		cm := CapacityMatrixNew(size)

		rnd := rand.New(rand.NewSource(int64(seed)))

		for townI := uint32(0); townI < size; townI++ {
			for townJ := townI + 1; townJ < townI + 3 && townJ < size; townJ++ {
				linkCap := float64(rnd.Intn(1023) + 1)
				cm.AddLink(townI, townJ, linkCap)
				cm.AddLink(townJ, townI, linkCap)
			}
		}

		clonedForMinCut := cm.Clone()
		globalMinCut := clonedForMinCut.MinCutStoerWagner()

		globalMinOfMaxFlow := math.MaxFloat64
		for s := uint32(0); s < size; s++ {
			for t := s + 1; t < size; t++ {
				clonedForMaxFlow := cm.Clone()
				maxFlowST := clonedForMaxFlow.MaxFlowFordFulkerson(s, t)
				if maxFlowST < globalMinOfMaxFlow {
					globalMinOfMaxFlow = maxFlowST
//					fmt.Printf("narrower cut of %d from %d to %d\n", maxFlowST, s, t)
				}
			}
		}

		{
			actual := globalMinCut
			expected := globalMinOfMaxFlow

			if actual != expected {
				for _,r := range cm {
					for _,linkCap := range r {
						fmt.Printf("%8f", linkCap)
					}
					fmt.Printf("\n")
				}

				t.Errorf(
					"globalMinCut is not correct for seed %d:\n\texpected %v\n\tactual %v",
					seed,
					expected,
					actual)
				t.FailNow()
			}
		}
		tests += 1
	}

	fmt.Printf("ran %d tests for Size %d\n", tests, size)

}


