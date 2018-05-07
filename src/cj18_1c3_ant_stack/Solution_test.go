package main

import (
	"testing"
	"math/rand"
	"time"
	"algo"
	"fmt"
)

func TestSimple(t *testing.T) {
	out := solveCase(caseInput{3, []int{10, 10, 10, 10, 10, 10, 10, 10, 100}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}

	{
		actual := out.longestSubsequence
		expected := 8

		if actual != expected {
			t.Errorf("out.longestSubsequence is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}

}

func TestSimple2(t *testing.T) {
	out := solveCase(caseInput{3, []int{1, 2, 1, 2, 1, 2, 1, 1, 1, 1}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}

	{
		actual := out.longestSubsequence
		expected := 7

		if actual != expected {
			t.Errorf("out.longestSubsequence is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}

}

func TestSimpleFailing(t *testing.T) {
	out := solveCase(caseInput{3, []int{8, 1}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}

	{
		actual := out.longestSubsequence
		expected := 1

		if actual != expected {
			t.Errorf("out.longestSubsequence is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}
}

func TestFiverFailing(t *testing.T) {
	out := solveCase(caseInput{3, []int{7, 7, 9, 3, 3}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}

	{
		actual := out.longestSubsequence
		expected := 4

		if actual != expected {
			t.Errorf("out.longestSubsequence is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}
}

func TestProbing(t *testing.T) {
	probe(2, 2, t)
	probe(2, 5, t)
	probe(2, 10, t)
	probe(2, 20, t)
}

func probe(timeoutSecs float64, weightLen int, t *testing.T) {
	started := time.Now()
	try := 0
	for ; time.Since(started).Seconds() < timeoutSecs; try++ {
		seed := try*101*103 + 1536*108*109
		rnd := rand.New(rand.NewSource(int64(seed)))

		weights := make([]int, weightLen, weightLen)
		for i := range weights {
			weights[i] = rnd.Intn(1000000000) + 1
		}

		out := solveCase(caseInput{3, weights})

		{
			actual := out.index
			expected := 3

			if actual != expected {
				t.Errorf("out.index is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
			}
		}

		flags := algo.SubsetFlagsNew(weightLen)
		longest := 0
		solution := ""
		for flags.Next() {
			valid := true
			currentLen := 0
			weightUpToI := 0
			for i, weight := range weights {
				if flags[i] {
					if weightUpToI <= weight*6 {
						currentLen += 1
						weightUpToI += weight
					} else {
						valid = false
						break
					}
				}
			}
			if valid && longest < currentLen {
				longest = currentLen
				solution = fmt.Sprintf("%v", flags)
			}
		}

		{
			actual := out.longestSubsequence
			expected := longest

			if actual != expected {
				t.Logf("failing input and solution:\n\t%v\n\t%v", weights, solution)
				t.Errorf(
					"out.longestSubsequence is not correct for seed %d\n\texpected:\t%v\n\tactual:\t%v",
					seed,
					expected,
					actual,
				)
				t.FailNow()
			}
		}
	}
	fmt.Printf("...executed %d probes for size %d\n", try, weightLen)
}
