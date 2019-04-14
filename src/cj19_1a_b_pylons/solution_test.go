package main

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	for _, rc := range [][2]int64{
		{2, 5}, {2, 6}, {2, 7}, {2, 8}, {2, 9}, {2, 10}, {2, 11}, {2, 12},
		{3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 9},
		{4, 4}, {4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9}, {4, 10}, {4, 11}, {4, 12},
		{5, 5}, {5, 6}, {5, 7}, {5, 8}, {5, 9}, {5, 10}, {5, 11}, {5, 12},
		{6, 6}, {6, 7}, {6, 8}, {6, 9}, {6, 10}, {6, 11}, {6, 12}, {6, 13},
		{7, 7}, {7, 8}, {7, 9}, {7, 10}, {7, 11}, {7, 12}, {7, 13}, {7, 14}, {7, 15}, {7, 16},
		{8, 8}, {8, 9}, {8, 10}, {8, 11}, {8, 12}, {8, 13}, {8, 14}, {8, 15}, {8, 16}, {8, 17},
		{9, 9}, {9, 10}, {9, 11}, {9, 12}, {9, 13},
		{10, 10}, {10, 11}, {10, 12}, {10, 13},
		{11, 11}, {11, 12}, {11, 13}, {11, 14},
	} {
		R := rc[0]
		C := rc[1]

		orderDump := make([][]int, R)
		order := 1;
		rPrev, cPrev := int64(-1), int64(-1)
		failAt := -1

		fill(R, C, func(ri, ci int64) {
			if orderDump[ri] == nil {
				orderDump[ri] = make([]int, C)
			}
			orderDump[ri][ci] = order

			if rPrev >= 0 && failAt < 0 {
				if rPrev == ri || cPrev == ci || abs64(rPrev-ri) == abs64(cPrev-ci) {
					failAt = order
				}
			}

			rPrev = ri
			cPrev = ci
			order += 1
		})

		if failAt != -1 {
			fmt.Printf("%dx%d:\n", R, C)
			for r := int64(0); r < R; r++ {
				for c := int64(0); c < C; c++ {
					fmt.Printf("\t%d", orderDump[r][c])
				}
				fmt.Printf("\n")
			}

			t.Fatalf("fill failed at %d", failAt)
		}

	}
}
