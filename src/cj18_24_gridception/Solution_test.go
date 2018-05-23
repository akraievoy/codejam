package main

import (
	"testing"
	"algo"
	"sort"
	"fmt"
)

func TestSimple(t *testing.T) {
	out := solveCase(caseInput{3, 3, 3, [][]bool{{false, false, false}, {false, true, false}, {false, false, false}}})

	{
		actual := out.index
		expected := 3

		if actual != expected {
			t.Errorf("out.index is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}

	{
		actual := out.largestPattern
		expected := uint32(8)

		if actual != expected {
			t.Errorf("out.largestPattern is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}
}

func TestDual(t *testing.T) {
	R := 3
	C := 3

	in := make([][]bool, R)
	inInv := make([][]bool, R)
	for r := 0; r < R; r++ {
		in[r] = make([]bool, C)
		inInv[r] = make([]bool, C)
	}

	flags := algo.SubsetFlagsNew(R * C)
	for {
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				in[r][c] = flags[r*C+c]
				inInv[r][c] = !flags[r*C+c]
			}
		}

		out := solveCase(
			caseInput{3, R, C, in},
		)
		outInv := solveCase(
			caseInput{3, R, C, inInv},
		)

		{
			actual := out
			expected := outInv

			if actual != expected {
				t.Errorf("out is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
			}
		}

		if !flags.Next() {
			break
		}
	}
}

type pos struct{ r, c int }

func connected(pattern [][]bool) (bool, int, int, int, int, int) {
	visited := make([][]bool, len(pattern))

	var firstSet *pos = nil
	minR, minC, maxR, maxC := len(pattern), len(pattern[0]), 0, 0
	setCount := 0
	for r, row := range pattern {
		visited[r] = make([]bool, len(row))
		for c, v := range row {
			if !v {
				continue
			}
			if c < minC {
				minC = c
			}
			if r < minR {
				minR = r
			}
			if c+1 > maxC {
				maxC = c + 1
			}
			if r+1 > maxR {
				maxR = r + 1
			}
			if setCount == 0 {
				firstSet = &pos{r, c}
			}
			setCount += 1
		}
	}

	if setCount == 0 {
		return false, -1, -1, -1, -1, -1
	}

	deltas := []pos{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	fillQ := append(make([]pos, 0), *firstSet)
	seenCount := 0
	for len(fillQ) > 0 {
		cur := fillQ[len(fillQ)-1]
		fillQ = fillQ[:len(fillQ)-1]
		if visited[cur.r][cur.c] {
			continue
		}
		visited[cur.r][cur.c] = true
		seenCount += 1

		for _, delta := range deltas {
			r := cur.r + delta.r
			c := cur.c + delta.c
			cValid := c >= 0 && c < len(pattern[cur.r])
			rValid := r >= 0 && r < len(pattern)

			if cValid && rValid && pattern[r][c] && !visited[r][c] {
				fillQ = append(fillQ, pos{r, c})
			}
		}
	}

	return setCount == seenCount, setCount, minR, minC, maxR, maxC
}

type pattern struct {
	mask                   [][]bool
	popcount               int
	minR, minC, maxR, maxC int
}

type patternz []pattern

func (p patternz) Len() int           { return len(p) }
func (p patternz) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p patternz) Less(i, j int) bool { return p[i].popcount > p[j].popcount }

func enumerateConnectedPatterns(R, C int) []pattern {
	patterns := make([]pattern, 0)
	flags := algo.SubsetFlagsNew(R * C)

	for flags.Next() {
		mask := make([][]bool, R)
		for r := 0; r < R; r++ {
			mask[r] = make([]bool, C)
			for c := 0; c < C; c++ {
				mask[r][c] = flags[r*C+c]
			}

		}

		connected, setCount, minR, minC, maxR, maxC := connected(mask)
		if connected {
			p := pattern{mask, setCount, minR, minC, maxR, maxC}
			patterns = append(patterns, p)
		}
	}

	sort.Sort(patternz(patterns))

	return patterns
}

func TestBrute(t *testing.T) {
	R := 3
	C := 4

	connectedPatterns := enumerateConnectedPatterns(R, C)

	in := make([][]bool, R)
	deepeeper := make([][]bool, R*4)
	for r := 0; r < R; r++ {
		in[r] = make([]bool, C)
	}
	for r := 0; r < R*4; r++ {
		deepeeper[r] = make([]bool, C*4)
	}

	flags := algo.SubsetFlagsNew(R * C)
	for {
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				in[r][c] = flags[r*C+c]
				for dr := 0; dr < 4; dr++ {
					for dc := 0; dc < 4; dc++ {
						deepeeper[r*4+dr][c*4+dc] = flags[r*C+c]
					}
				}
			}
		}

		out := solveCase(
			caseInput{3, R, C, in},
		)

		//	as we go through patterns in decreasing popcount order, the first we find is the max
		var firstMatch *pattern = nil
		for _, p := range connectedPatterns {
			patternHeight := p.maxR - p.minR
			patternWidth := p.maxC - p.minC

			/*
						t := true
						f := false
						//noinspection GoBoolExpressions
						if reflect.DeepEqual(in, [][]bool{{f, f, f, f}, {f, f, t, t}, {t, t, t, f}}) &&
							reflect.DeepEqual(p.mask, [][]bool{{t, t, t, t}, {t, t, f, f}, {t, t, t, t}}) {
							fmt.Printf("oi!\n")
						}
			*/
			for drb := 0; drb <= 4*R-patternHeight && firstMatch == nil; drb++ {
				for dcb := 0; dcb <= 4*C-patternWidth && firstMatch == nil; dcb++ {
					patternMatches := true
					for drd := 0; drd < patternHeight && patternMatches; drd++ {
						for dcd := 0; dcd < patternWidth; dcd++ {
							if p.mask[p.minR+drd][p.minC+dcd] && in[p.minR+drd][p.minC+dcd] != deepeeper[drb+drd][dcb+dcd] {
								patternMatches = false
								break
							}
						}
					}

					if patternMatches {
						firstMatch = &p
					}
				}
			}

			if firstMatch != nil {
				break
			}

		}

		{
			actual := int(out.largestPattern)
			expected := firstMatch.popcount

			if actual != expected {

				fmt.Printf("input:\n")
				for _, row := range in {
					fmt.Printf("\t")
					for _, v := range row {
						c := 'B'
						if v {
							c = 'W'
						}
						fmt.Printf("%c", c)
					}
					fmt.Printf("\n")
				}
				fmt.Printf("pattern:\n")
				for _, row := range firstMatch.mask {
					fmt.Printf("\t")
					for _, v := range row {
						c := '.'
						if v {
							c = '#'
						}
						fmt.Printf("%c", c)
					}
					fmt.Printf("\n")
				}
				t.Errorf("out.largestPattern is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)

				t.FailNow()
			}
		}

		if !flags.Next() {
			break
		}
	}
}

/*
func TestGen(t *testing.T) {
	R := 3
	C := 4

	in := make([][]bool, R)
	for r := 0; r < R; r++ {
		in[r] = make([]bool, C)
	}

	flags := algo.SubsetFlagsNew(R * C)
	file, e := os.Create("03.in")
	defer file.Close()

	if e != nil {
		t.Errorf("failed to open file: %v", e)
		t.FailNow()
	}

	fmt.Fprintf(file, "%d\n", int64(math.Pow(float64(2), float64(R*C))))
	for {
		fmt.Fprintf(file, "%d %d\n", R, C)
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				v := 'B'
				if flags[r*C+c] {
					v = 'W'
				}
				fmt.Fprintf(file, "%c", v)
			}
			fmt.Fprintf(file, "\n")
		}

		if !flags.Next() {
			break
		}
	}
}
*/

/*
func TestForFile(t *testing.T) {
	os.Setenv("CODEJAM_INPUT", "../../inout/cj18_24_gridception/01.in")
	main()
}
*/
