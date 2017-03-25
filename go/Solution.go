package main

import (
	"fmt"
	"time"
	"sort"
)

var primes = []uint16{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199 }

const debug = false

type ModRem struct {
	Modulo    uint16
	Remainder uint16
}

type ModRemCompare []ModRem

func (c ModRemCompare) Len() int {
	return len(c)
}
func (c ModRemCompare) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c ModRemCompare) Less(i, j int) bool {
	mri := c[i]
	mrj := c[j]
	if mri.Modulo == mrj.Modulo {
		return mri.Remainder < mrj.Remainder
	} else {
		return mri.Modulo < mrj.Modulo
	}
}

func main() {
	//  the idea is to cache matching indexes
	//    for modRemKeys of all prime power / modulo pairs
	//    that we actually use in queries
	var modRemToIndexes = make(map[ModRem][]uint16)
	var n, queries uint16;
	fmt.Scan(&n, &queries)

	var a = make([]uint16, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	var lefts = make([]uint16, n)
	var rights = make([]uint16, n)

	var queryModRems = make([][]ModRem, 0)
	var q uint16 = 0;
	for ; q < queries; q++ {
		fmt.Scan(&lefts[q], &rights[q])
		var queryModulo, queryRemainder uint16
		fmt.Scan(&queryModulo, &queryRemainder)

		var queryModRemsLocal = make([]ModRem, 0)
		var primeTriedIdx = 0
		for queryModulo > 1 && primeTriedIdx < len(primes) {
			var primeTried = primes[primeTriedIdx]
			primeTriedIdx++
			var power uint16 = 1
			for queryModulo % primeTried == 0 {
				queryModulo /= primeTried
				power *= primeTried
			}
			if power > 1 {
				queryModRemsLocal = append(queryModRemsLocal, ModRem{power, queryRemainder % power})
			}
		}
		//  that's a prime, we tested all divisors below sqrt(40k)
		if queryModulo > 1 {
			queryModRemsLocal = append(queryModRemsLocal, ModRem{queryModulo, queryRemainder % queryModulo})
		}
		sort.Sort(ModRemCompare(queryModRemsLocal))
		queryModRems = append(queryModRems, queryModRemsLocal)
	}

	//  a bit of lousy collections code here:
	//    we won't have more than 120k unique mod/rem combinations
	//      even for the most carefully crafted adverse/evil test case
	//        each query has to use products of globally disjoint sets of mod/rems with mods we keep globally co-prime
	//          worst case is 7 mods (2,3,5,7,11,13 = 30030), average would be like 3 for 40k queries
	var dedupeStartedNanos = time.Now().UnixNano()
	var uniqueModRemsSet = make(map[ModRem](interface{}))
	for _, modRems := range queryModRems {
		for _, modRem := range modRems {
			uniqueModRemsSet[modRem] = nil
		}
	}
	var uniqueModRems []ModRem = make([]ModRem, 0/*, len(uniqueModRemsSet*/)
	for modRem := range uniqueModRemsSet {
		uniqueModRems = append(uniqueModRems, modRem)
	}
	sort.Sort(ModRemCompare(uniqueModRems))
	var dedupeCompletedNanos = time.Now().UnixNano()

	if debug {
		fmt.Printf("uniqueModRems.length = %d in %d nanos\n", 0, (dedupeCompletedNanos - dedupeStartedNanos))
	}

	var divisionScans uint16 = 0
	if (len(uniqueModRems) > 0) {
		var interestingRems []uint16 = make([]uint16, 0, 40000)
		var mod = uniqueModRems[0].Modulo
		interestingRems = append(interestingRems, uniqueModRems[0].Remainder)
		for _, modRem := range uniqueModRems {
			if modRem.Modulo == mod {
				interestingRems = append(interestingRems, modRem.Remainder)
			} else {
				divisionScans++
				populateIndexes(&a, mod, &interestingRems, &modRemToIndexes)
				mod = modRem.Modulo
				interestingRems = append(interestingRems[0:0], modRem.Remainder)
			}
		}
		populateIndexes(&a, mod, &interestingRems, &modRemToIndexes)
		divisionScans++
	}

	if debug {
		fmt.Printf("divisionScans = %d\n", divisionScans)
	}

	intersection, intersectionTemp := make([]uint16, 0, 40000), make([]uint16, 0, 40000)
	for q = 0; q < queries; q++ {
		left := lefts[q]
		right := rights[q]
		modRems := queryModRems[q]

		if (len(modRems) == 0) {
			//  someone tries to query by module = 1?
			//  ohhhh... ohhhh...
			//  at least module is not zero, that would be really PATHETIC
			fmt.Printf("%d\n", right - left + 1)
			continue
		}

		sparsest := modRemToIndexes[modRems[len(modRems) - 1]]
		if len(modRems) == 1 {
			leftPos, rightPos := sliceRange(sparsest, left, right)
			fmt.Printf("%d\n", rightPos - leftPos)
		} else {
			leftPos, rightPos := sliceRange(sparsest, left, right)
			intersection = append(intersection[0:0], sparsest[leftPos:rightPos]...)

			for modRemI := len(modRems) - 2; len(intersection) > 0 && modRemI >= 0; modRemI-- {
				lefter, righter := intersection[0], intersection[len(intersection) - 1]
				modRemIdx := modRemToIndexes[modRems[modRemI]]
				lefterPos, righterPos := sliceRange(modRemIdx, lefter, righter)

				intersectionPos := 0
				intersectionTemp = intersectionTemp[0:0]
				for idxPos := lefterPos; intersectionPos < len(intersection) && idxPos < righterPos; {
					elemI := intersection[intersectionPos]
					elemM := modRemIdx[idxPos]

					if (elemI < elemM) {
						intersectionPos++
					} else if (elemI == elemM) {
						intersectionTemp = append(intersectionTemp, elemI)
						idxPos++
						intersectionPos++
					} else {
						idxPos++
					}
				}

				intersection, intersectionTemp = intersectionTemp, intersection
			}

			fmt.Printf("%d\n", len(intersection))
		}
	}
}

func sliceRange(sorted []uint16, left uint16, right uint16) (leftPos int, rightPos int) {
	sLen := len(sorted)
	leftPos = sort.Search(sLen + 1, func(idx int) bool {
		return idx == sLen || sorted[idx] >= left
	})
	rightPos = sort.Search(sLen + 1, func(idx int) bool {
		return idx == sLen || sorted[idx] > right
	})
	return
}

func populateIndexes(a *[]uint16, mod uint16, interestingRems *[]uint16, modRemToIndexes *map[ModRem][]uint16) {
	var indexes [][]uint16 = make([][]uint16, mod)
	for _, interestingRem := range *interestingRems {
		indexes[interestingRem] = make([]uint16, 0, len(*a))
	}
	for i, aI := range *a {
		aRem := aI % mod
		if indexes[aRem] == nil {
			continue
		}
		indexes[aRem] = append(indexes[aRem], uint16(i))
	}
	for _, interestingRem := range *interestingRems {
		indexesI := indexes[interestingRem]
		(*modRemToIndexes)[ModRem{mod, interestingRem}] = append(make([]uint16, 0, len(indexesI)), indexesI...)
	}
}
