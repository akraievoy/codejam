package main

import (
	"testing"
	"time"
	"math/rand"
)

func TestSimple(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	out := solveCase(caseInput{3, 470, 143, 5, []int{11, 24, 420, 6, 9}})

	{
		actual := out.sectionForK
		expected := 3

		if actual != expected {
			t.Errorf("out.sectionForK is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}


}

func TestSimpleBoundaryLeft(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	out := solveCase(caseInput{3, 470, 36, 5, []int{11, 24, 420, 6, 9}})

	{
		actual := out.sectionForK
		expected := 3

		if actual != expected {
			t.Errorf("out.sectionForK is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}


}

func TestSimpleBoundaryRight(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	out := solveCase(caseInput{3, 470, 455, 5, []int{11, 24, 420, 6, 9}})

	{
		actual := out.sectionForK
		expected := 3

		if actual != expected {
			t.Errorf("out.sectionForK is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}


}

func TestSimpleBoundaryNextLeft(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	out := solveCase(caseInput{3, 470, 456, 5, []int{11, 24, 420, 6, 9}})

	{
		actual := out.sectionForK
		expected := 4

		if actual != expected {
			t.Errorf("out.sectionForK is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}


}

func TestSimpleBoundaryLeftMost(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	out := solveCase(caseInput{3, 470, 470, 5, []int{11, 24, 420, 6, 9}})

	{
		actual := out.sectionForK
		expected := 5

		if actual != expected {
			t.Errorf("out.sectionForK is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}


}

func TestSimpleBoundaryLeftmost1(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	out := solveCase(caseInput{3, 470, 470, 5, []int{11, 24, 420, 14, 1}})

	{
		actual := out.sectionForK
		expected := 5

		if actual != expected {
			t.Errorf("out.sectionForK is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}


}

func TestSimpleBoundaryRightmost1(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	out := solveCase(caseInput{3, 470, 1, 5, []int{1, 34, 420, 14, 1}})

	{
		actual := out.sectionForK
		expected := 1

		if actual != expected {
			t.Errorf("out.sectionForK is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}


}
