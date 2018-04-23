package main

import (
	"testing"
	"math/rand"
	"fmt"
	"sort"
	"time"
)

func TestSimple(t *testing.T) {
	testForSeed(time.Now().UnixNano(), 3, 4, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 4, 4, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 5, 4, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 3, 5, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 4, 5, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 5, 5, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 3, 6, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 4, 6, 10000, true, t)
	testForSeed(time.Now().UnixNano(), 5, 6, 10000, true, t)
}

func TestSimpleUpTo12(t *testing.T) {
	testForSeed(time.Now().UnixNano(), 3, 10, 50, true, t)
	testForSeed(time.Now().UnixNano(), 4, 10, 50, true, t)
	testForSeed(time.Now().UnixNano(), 5, 10, 50, true, t)
	testForSeed(time.Now().UnixNano(), 3, 11, 25, true, t)
	testForSeed(time.Now().UnixNano(), 4, 11, 25, true, t)
	testForSeed(time.Now().UnixNano(), 5, 11, 25, true, t)
	testForSeed(time.Now().UnixNano(), 3, 12, 10, true, t)
	testForSeed(time.Now().UnixNano(), 4, 12, 10, true, t)
	testForSeed(time.Now().UnixNano(), 5, 12, 10, true, t)
}

func TestLarge(t *testing.T) {
	testForSeed(time.Now().UnixNano(), 3, 20, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 4, 20, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 5, 20, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 3, 40, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 4, 40, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 5, 40, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 3, 80, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 4, 80, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 5, 80, 10000, false, t)
}

func TestFullAlNumLarge(t *testing.T) {
	testForSeed(time.Now().UnixNano(), 26, 20, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 26, 40, 10000, false, t)
	testForSeed(time.Now().UnixNano(), 26, 80, 10000, false, t)
}

func testForSeed(seed int64, letters int, wordLen int, tests int, permutations bool, t *testing.T) {
	fmt.Printf("seed = %d\n", seed)
	var rnd = rand.New(rand.NewSource(seed))
	freq := make([]int, letters)
	word := make([]uint8, wordLen)
	for test := 0; test < tests; test++ {
		for i := range freq {
			freq[i] = 0
		}
		//fmt.Printf("test #%d\n", test)
		for i := 0; i < wordLen; i++ {
			letter := uint8(rnd.Intn(len(freq)))
			freq[letter] += 1
			word[i] = letter
		}
		sort.Sort(Uint8Slice(word))

		bestSuperiorsPerm := superiors(word)
		permWord := make([]uint8, wordLen)
		if permutations {
			perm := PermNew(wordLen)
			//fmt.Printf("S %4d\t%v\n", bestSuperiorsPerm, word)
			for true {
				for i := range word {
					permWord[i] = word[perm[i]]
				}
				superiorsPerm := superiors(permWord)

				if superiorsPerm > bestSuperiorsPerm {
					//fmt.Printf("B %4d\t%v\n", superiorsPerm, permWord)
					bestSuperiorsPerm = superiorsPerm
				}

				if !perm.Next() || superiorsPerm == (wordLen-1)/2 {
					break
				}
			}
		}
		for i := range permWord {
			permWord[i] = 255
		}
		for i := range word {
			if i*2 < len(word) {
				permWord[i*2] = word[i]
			} else {
				target := 1 + (i-(len(word)+1)/2)*2
				if len(word)%2 == 0 /*&& word[i/2] == word[0]*/ {
					if i == len(word)-1 {
						permWord[target] = word[len(word)/2]
					} else {
						permWord[target] = word[i+1]
					}
				} else {
					permWord[target] = word[i]
				}
			}
		}
		superiorsOpt := superiors(permWord)

		if superiorsOpt < (wordLen-1)/2 {
			fmt.Printf("%v\t%4d\t%v\n", freq, superiorsOpt, permWord)
		}

		if permutations {
			actual := bestSuperiorsPerm
			expected := superiorsOpt

			if actual != expected {
				t.Errorf("bestSuperiorsPerm is not correct: expected %v, actual %v", expected, actual)
				t.FailNow()
			}
		}
		{
			actual := solveCase(caseInput{0, freq}).superiors
			expected := superiorsOpt

			if actual != expected {
				t.Errorf("solveCase(...).superiors is not correct: expected %v, actual %v", expected, actual)
				t.FailNow()
			}
		}
	}
}

func superiors(word []uint8) int {
	superiorsPerm := 0
	for i := range word {
		if i-1 < 0 || i+1 == len(word) {
			continue
		}
		if word[i] > word[i-1] && word[i] > word[i+1] {
			superiorsPerm += 1
		}
	}
	return superiorsPerm
}

type Perm []uint32

func PermNew(size int) Perm {
	perm := make([]uint32, size)
	for i := range perm {
		perm[i] = uint32(i)
	}
	return perm
}

//	https://www.nayuki.io/page/next-lexicographical-permutation-algorithm

func (perm Perm) Next() bool {
	rollPos := len(perm) - 2
	for rollPos >= 0 && perm[rollPos] > perm[rollPos+1] {
		rollPos -= 1
	}
	if rollPos < 0 {
		return false
	}

	swapPos := rollPos + 1
	for pos := rollPos + 2; pos < len(perm); pos++ {
		if perm[rollPos] < perm[pos] && perm[pos] < perm[swapPos] {
			swapPos = pos
		}
	}
	perm[rollPos], perm[swapPos] = perm[swapPos], perm[rollPos]

	for i, j := rollPos+1, len(perm)-1; i < j; i, j = i+1, j-1 {
		perm[i], perm[j] = perm[j], perm[i]
	}

	return true
}

type Uint8Slice []uint8

func (p Uint8Slice) Len() int           { return len(p) }
func (p Uint8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
