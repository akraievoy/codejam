package main

import (
	"testing"
	"algo"
	"fmt"
	"time"
	"math/rand"
)

func TestSimple(t *testing.T) {
	out :=
		solveCase(
			caseInput{
				4,
				[]group{
					{[]uint32{1, 2}, 4},
					{[]uint32{1, 3}, 5},
					{[]uint32{2, 3, 4}, 10},
				},
			},
		)

	{
		actual := out.efficiencyTotal
		expected := uint64(10)

		if actual != expected {
			t.Errorf("out.efficiencyTotal is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}

}

func TestProbing(t *testing.T) {
	probe(2, 2, 1, 5, t)
	probe(2, 5, 3, 5, t)
	probe(2, 10, 5, 5, t)
	probe(2, 15, 6, 5, t)
}

func probe(timeoutSecs float64, persons, groups int, efficiencyCap int, t *testing.T) {
	started := time.Now()
	try := 0
	for ; time.Since(started).Seconds() < timeoutSecs; try++ {
		seed := try*101*103 + 1536*107*109
		rnd := rand.New(rand.NewSource(int64(seed)))

		gs := make([]group, groups)
		for g := range gs {
			eligible := make([]uint32, persons)
			for i := range eligible {
				eligible[i] = uint32(i + 1);
			}
			groupSize := rnd.Intn(2) + 2
			members := make([]uint32, groupSize)
			for m := range members {
				chosenIdx := rnd.Intn(len(eligible))
				members[m] = eligible[chosenIdx]
				eligible = append(eligible[chosenIdx:], eligible[:chosenIdx+1]...)
			}
			efficiency := uint64(1 + rnd.Intn(efficiencyCap))
			gs[g] = group{members, efficiency}
		}

		out := solveCase(caseInput{uint32(persons), gs})

		validSeen := false
		bestEfficiency := uint64(0)

		membersIn := make([]map[bool]bool, groups)

		for m := range membersIn {
			membersIn[m] = make(map[bool]bool)
		}

		flags := algo.SubsetFlagsNew(persons)
		for {
			for _, mIn := range membersIn {
				mIn[false], mIn[true] = false, false
			}

			for groupIdx, group := range gs {
				for _, mIdx := range group.members {
					membersIn[groupIdx][flags[mIdx-1]] = true
				}
			}
			globalGroupIn := make(map[bool]bool)
			for _, g := range flags {
				globalGroupIn[g] = true
			}

			if globalGroupIn[false] && globalGroupIn[true] {
				efficiency := uint64(0)
				for gIdx, g := range gs {
					if membersIn[gIdx][false] && membersIn[gIdx][true] {
						continue
					}
					efficiency += g.efficiency
				}

				if !validSeen || efficiency > bestEfficiency {
					validSeen = true
					bestEfficiency = efficiency
				}
			}

			if !flags.Next() {
				break
			}
		}

		{
			actual := out.efficiencyTotal
			expected := bestEfficiency

			if actual != expected {
				t.Errorf("out.efficiencyTotal is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
			}
		}
	}

	fmt.Printf("...executed %d probes for %d persons and %d groups\n", try, persons, groups)
}
