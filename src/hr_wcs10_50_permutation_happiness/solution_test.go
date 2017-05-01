package main

import (
	"testing"
	"sort"
)

func TestGeneration(t *testing.T) {
	limit := 10

	preco := precompute()
	for n := 1; n <= limit; n++ {
		for k := 0; k <= n; k++ {
			if false {
				print(preco[n][k], "\t")
			}
		}
		if false {
			println()
		}
	}

	for n := 1; n <= limit; n++ {
		counts := make([]int, n + 1)
		perm := make([]int, n)
		for i := 0; i < n; i++ {
			perm[i] = i
		}

		for {
			happyCount := 0
			for i := 0; i < n; i++ {
				if (i > 0 && perm[i] < perm[i - 1] || i + 1 < n && perm[i] < perm[i + 1]) {
					happyCount++
				}
			}
			counts[happyCount]++

			rollPos := n - 2
			for rollPos >= 0 && perm[rollPos] > perm[rollPos + 1] {
				rollPos--
			}
			if (rollPos < 0) {
				break
			}
			nextPos := -1
			for i := rollPos + 1; i < n; i++ {
				if perm[i] > perm[rollPos] && (nextPos < 0 || perm[nextPos] > perm[i]) {
					nextPos = i
				}
			}
			perm[rollPos], perm[nextPos] = perm[nextPos], perm[rollPos]
			sort.Ints(perm[rollPos + 1:])
		}

		if false {
			for i := 0; i <= n; i++ {
				print(counts[i])
				if (i == n) {
					print("\n")
				} else {
					print("\t")
				}
			}
		}

		for i := 0; i <= n; i++ {
			{
				expected := preco[n][i];
				actual := counts[i];

				if actual != expected {
					t.Errorf("counts[i] is not correct: expected %v, returned %v", expected, actual)
				}
			}
		}
	}
}
