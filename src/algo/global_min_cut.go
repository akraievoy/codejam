package algo

import "math"

/*
	http://www.e-maxx-ru.1gb.ru/algo/stoer_wagner_mincut#4

	http://www.e-maxx-ru.1gb.ru/bookz/files/stoer_wagner_mincut.pdf
	http://www.e-maxx-ru.1gb.ru/bookz/files/mehlhorn_mincut_stoer_wagner.pdf
 */
func (cm *CapacityMatrix) MinCutStoerWagner() float64 {
	minCutSet := false
	minCut := float64(math.MaxFloat64)

	g := *cm
	n := len(g)
	w := make([]float64, n)
	exist := make([]bool, n)
	for i := range exist {
		exist[i] = true
	}
	inA := make([]bool, n)

	for phase := 0; phase < n-1; phase++ {
		for i := range inA {
			inA[i] = false
			w[i] = 0
		}

		prev := -1
		for it := 0; it < n-phase; it++ {

			sel := -1
			for i := 0; i < n; i++ {
				if exist[i] && !inA[i] && (sel == -1 || w[i] > w[sel]) {
					sel = i;
				}
			}

			if it == n-phase-1 {
				if !minCutSet || w[sel] < minCut {
					minCutSet = true
					minCut = w[sel]
				}

				for i := 0; i < n; i++ {
					g[i][prev] += g[sel][i]
					g[prev][i] = g[i][prev]
				}
				exist[sel] = false
			} else {
				inA[sel] = true
				for i := 0; i < n; i++ {
					w[i] += g[sel][i]
				}
				prev = sel
			}

		}
	}

	return minCut;
}
