package main

import (
	"testing"
	"reflect"
	"math/rand"
	"time"
	"fmt"
)

func TestFactorizeOne(t *testing.T) {
	factors, powers := factorize(1, make([]int, 20), make([]int, 20))

	{
		expected := []int{1};
		actual := factors;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("factors is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := []int{1};
		actual := powers;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("powers is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
func TestFactorize(t *testing.T) {
	factors, powers := factorize(498677, make([]int, 20), make([]int, 20))

	{
		expected := []int{53, 97};
		actual := factors;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("factors is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := []int{1, 2};
		actual := powers;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("powers is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestFactorizePrime(t *testing.T) {
	factors, powers := factorize(402763, make([]int, 20), make([]int, 20))

	{
		expected := []int{402763};
		actual := factors;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("factors is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := []int{1};
		actual := powers;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("powers is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimpleFull(t *testing.T) {
	out := solve(In{[]rune("ccaccbbbaccccca")})

	{
		expected := 2;
		actual := out.triples;

		if actual != expected {
			t.Errorf("out.triples is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func testForSizeAndSeed(t *testing.T, size int, seed int64) {
	r := rand.New(rand.NewSource(seed))
	abc := make([]rune, size)
	_abc := []rune("abc")
	for i := range abc {
		abc[i] = _abc[r.Intn(len(_abc))]
	}

	expected := 0
	for i := range abc {
		if abc[i] != 'a' {
			continue
		}
		for j := range abc {
			jSq := (j + 1) * (j + 1)
			if abc[j] != 'b' || jSq % (i + 1) > 0 {
				continue
			}
			for k := range abc {
				if abc[k] != 'c' {
					continue
				}
				if (i + 1) * (k + 1) == jSq {
					expected += 1
				}
			}
		}
	}

	{
		expected := expected;
		actual := solve(In{abc}).triples;
		fmt.Println(fmt.Sprintf("%v -> %v" , string(abc), expected))

		if actual != expected {
			t.Errorf("solve(In{abc}).triples is not correct for seed %v: expected %v, returned %v", seed, expected, actual)
		}
	}
}

func TestLargeLoop(t *testing.T) {
	for i := 0; i < 20; i++ {
		testForSizeAndSeed(t, 2000, time.Now().UnixNano())
	}
}

func TestMedLoop(t *testing.T) {
	for i := 0; i < 100; i++ {
		testForSizeAndSeed(t, 40, time.Now().UnixNano())
	}
}

func TestLargerLoop(t *testing.T) {
	t.SkipNow()
	for i := 0; i < 10; i++ {
		testForSizeAndSeed(t, 5000, time.Now().UnixNano())
	}
}

func TestLargest(t *testing.T) {
	t.SkipNow()

	testForSizeAndSeed(t, 500000, time.Now().UnixNano())
}

func TestGenerate(t *testing.T) {
	t.SkipNow()

	r := rand.New(rand.NewSource(10305))
	abc := make([]rune, 500000)
	_abc := []rune("abc")
	for i := range abc {
		abc[i] = _abc[r.Intn(len(_abc))]
	}
	fmt.Println("GENERATED")
	fmt.Println(fmt.Sprintf("%s", string(abc)))
}
