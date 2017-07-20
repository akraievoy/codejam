package main

import (
	"testing"
	"time"
	"math/rand"
	"fmt"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]int{3, 1, 4, 2, 8}, []int{5, 2, 12, 8, 3}})

	{
		expected := 16;
		actual := out.sum;

		if actual != expected {
			t.Errorf("out.sum is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestLinear(t *testing.T) {
	A := make([]int, 100000, 100000)
	B := make([]int, 100000, 100000)
	for i := range A {
		A[i] = (i + 1) * 9
		B[i] = (i + 1) * 9
	}
	out := solve(In{A, B})

	{
		expected := 100000 * 9 * 2;
		actual := out.sum;

		if actual != expected {
			t.Errorf("out.sum is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestRandomFailure(t *testing.T) {
	A := make([]int, 2, 2)
	B := make([]int, 2, 2)
	A[0] = 33556
	A[1] = 69131
	B[0] = 38924
	B[1] = 73705

	out := solve(In{A, B})

	{
		expected := 72480;
		actual := out.sum;

		if actual != expected {
			t.Errorf("out.sum is not correct: expected %v, returned %v", expected, actual)
		}
	}
}


func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func testForSize(t *testing.T, size int) bool {
	var rnd = rand.NewSource(time.Now().UnixNano())
	A := make([]int, size, size)
	B := make([]int, size, size)
	for i := range A {
		A[i] = abs(int(rnd.Int63())) % 100000 + 1
		B[i] = abs(int(rnd.Int63())) % 100000 + 1
	}

	out := solve(In{A, B})

	maxSum := 0
	maxGcd := 0
	for _, x := range A {
		for _, y := range B {
			curGcd := gcd(x, y)
			curSum := x + y
			if curGcd > maxGcd || (curGcd == maxGcd && curSum > maxSum) {
				maxGcd = curGcd
				maxSum = curSum
			}
		}
	}

	{
		expected := maxSum;
		actual := out.sum;

		if actual != expected {
			fmt.Println(fmt.Sprintf("A = %v", A))
			fmt.Println(fmt.Sprintf("B = %v", B))
			t.Errorf("out.sum is not correct: expected %v, returned %v", expected, actual)
			return false
		}
	}

	return true
}

func TestVersusBrute(t *testing.T) {
	for i := 0; i < 2; i++ {
		if !testForSize(t, 2) {
			break
		}
	}
}
