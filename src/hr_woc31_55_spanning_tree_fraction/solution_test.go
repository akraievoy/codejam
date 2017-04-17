package main

import (
	"testing"
	"fmt"
)

func TestSimple(t *testing.T) {
	out := solve(In{
		3,
		2,
		/*
		[]int{0, 1, 2},
		*/
		[]Link{{0, 1, 2, 3, 2.0 / 3.0}, {1, 2, 3, 4, 3.0 / 4.0}} })

	if out.sumA != 5 {
		t.Errorf("sumA == %d != %d", out.sumA, 5)
	}
	if out.sumB != 7 {
		t.Errorf("sumB == %d != %d", out.sumB, 7)
	}
}

func TestHackerRank(t *testing.T) {
	out := solve(In{
		3,
		3,
		/*
		[]int{0, 1, 2, 3},
		*/
		[]Link{
			{0, 1, 1, 1, 1.0 / 1.0},
			{1, 2, 2, 4, 2.0 / 4.0},
			{2, 0, 1, 2, 1.0 / 2.0} } })

	expectedA := uint32(2)
	expectedB := uint32(3)
	if float64(out.sumA) / float64(out.sumB) < float64(expectedA) / float64(expectedB) {
		t.Errorf("sumA/sumB = %f < %f", float64(out.sumA) / float64(out.sumB), float64(expectedA) / float64(expectedB))
	}
	if out.sumA != expectedA {
		t.Errorf("sumA == %d != %d", out.sumA, expectedA)
	}
	if out.sumB != expectedB {
		t.Errorf("sumB == %d != %d", out.sumB, expectedB)
	}

}

func TestMultiLink(t *testing.T) {
	out := solve(In{
		3,
		2,
		/*
		[]int{0, 1, 3},
		*/
		[]Link{
			{0, 1, 1, 100, 1.0 / 100.0},
			{1, 2, 96, 10, 96.0 / 10.0},
			{1, 2, 88, 1, 88.0 / 1.0}} })

	expectedA := uint32(97)
	expectedB := uint32(110)
	if float64(out.sumA) / float64(out.sumB) < float64(expectedA) / float64(expectedB) {
		t.Errorf("sumA/sumB = %f < %f", float64(out.sumA) / float64(out.sumB), float64(expectedA) / float64(expectedB))
	}
	if out.sumA != expectedA {
		t.Errorf("sumA == %d != %d", out.sumA, expectedA)
	}
	if out.sumB != expectedB {
		t.Errorf("sumB == %d != %d", out.sumB, expectedB)
	}
}

func TestMultiLink2(t *testing.T) {
	out := solve(In{
		3,
		2,
		/*
		[]int{0, 1, 3},
		*/
		[]Link{
			{0, 1, 1, 100, 1.0 / 100.0},
			{1, 2, 96, 12, 96.0 / 12.0},
			{1, 2, 88, 1, 88.0 / 1.0}} })

	expectedA := uint32(89)
	expectedB := uint32(101)
	if float64(out.sumA) / float64(out.sumB) < float64(expectedA) / float64(expectedB) {
		t.Errorf("sumA/sumB = %f < %f", float64(out.sumA) / float64(out.sumB), float64(expectedA) / float64(expectedB))
	}
	if out.sumA != expectedA {
		t.Errorf("sumA == %d != %d", out.sumA, expectedA)
	}
	if out.sumB != expectedB {
		t.Errorf("sumB == %d != %d", out.sumB, expectedB)
	}
}

func TestRenderOptimality(t *testing.T) {
	t.SkipNow()
	renderPairs([][2]float64{{2, 2}, {4, 4}, {3, 4}, {6, 8}, {4, 7}, {8, 14}, {5, 11}, {10, 22}, {6, 17}, {12, 34}, {7, 24}})
	renderPairs([][2]float64{{2, 1}, {2, 2}, {4, 2}, {5, 3}, {8, 7}, {11, 20}, {15, 40}})
}

func renderPairs(pairs [][2]float64) {
	wasOptimal := make([]bool, len(pairs))
	optimalABs := make([][]float64, len(pairs))

	fmt.Println(fmt.Sprintf("%v", pairs))

	deltas := make([]float64, len(pairs))

	for i, p := range pairs {
		deltas[i] = p[0] - p[1]
	}

	fmt.Println(fmt.Sprintf("%v", deltas))

	rates := make([]float64, len(pairs))

	for i, p := range pairs {
		rates[i] = p[0] / p[1]
	}

	fmt.Println(fmt.Sprintf("%v", rates))

	results := make([]float64, len(pairs))
	for a := float64(1); a <= 3000; a += 1 {
		for b := float64(1); b <= 3000; b += 1 {
			for i, p := range pairs {
				results[i] = (a + p[0]) / (b + p[1])
			}
			max := results[0]
			maxI := 0
			for i := 1; i < len(results); i++ {
				if max < results[i] {
					maxI = i
					max = results[i]
				}
			}
			if !wasOptimal[maxI] {
				wasOptimal[maxI] = true
				optimalABs[maxI] = []float64{a, b}
			}
		}
	}

	fmt.Println(fmt.Sprintf("%v", wasOptimal))
	fmt.Println(fmt.Sprintf("%v", optimalABs))

	for i := range wasOptimal {
		if wasOptimal[i] {
			fmt.Println(fmt.Sprintf("%v is optimal for %v with delta %d and ratio %.3f", pairs[i], optimalABs[i], int64(deltas[i]), rates[i]))
		}
	}

	for a := 1; a < 50; a++ {
		for b := 1; b < 100; b++ {
			for i, p := range pairs {
				results[i] = (float64(a) + p[0]) / (float64(b) + p[1])
			}
			max := results[0]
			maxI := 0
			for i := 1; i < len(results); i++ {
				if max < results[i] {
					maxI = i
					max = results[i]
				}
			}
			fmt.Print(fmt.Sprintf("%d", maxI))
		}
		fmt.Println();
	}
}
