package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type In struct {
	index int
	R     int
	C     int
	cake  [][]rune
}

type Out struct {
	index int
	cake  [][]rune
}

func min(a, b int) int {
	if (a < b) {
		return a
	}
	return b
}

func max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func allFree(rS, cS, rE, cE int, cakePtr *[][]rune) bool {
	cake := *cakePtr
	for r := rS; r < rE; r++ {
		for c := cS; c < cE; c++ {
			if cake[r][c] != '?' {
				return false
			}
		}
	}

	return true
}

func mark(rS, cS, rE, cE int, cakePtr *[][]rune, letter int) {
	cake := *cakePtr
	l := 'A' + int32(letter)
	for r := rS; r < rE; r++ {
		for c := cS; c < cE; c++ {
			if cake[r][c] != '?' {
				panic(fmt.Sprintf("trying to set %c over %c at %d %d", l, cake[r][c], r, c))
			}
			cake[r][c] = l
		}
	}
}

func solve(in In) (out Out) {
	letterFlags := make([]bool, 26)
	minR := make([]int, 26)
	minC := make([]int, 26)
	maxR := make([]int, 26)
	maxC := make([]int, 26)
	for l := 0; l < 26; l++ {
		minR[l], minC[l], maxR[l], maxC[l] = in.R, in.C, -1, -1
	}
	for r, line := range in.cake {
		for c, l := range line {
			if l == '?' {
				continue;
			}
			letter := l - 'A'
			letterFlags[letter] = true
			minR[letter] = min(minR[letter], r)
			minC[letter] = min(minC[letter], c)
			maxR[letter] = max(maxR[letter], r + 1)
			maxC[letter] = max(maxC[letter], c + 1)
		}
	}
	letters := make([]int, 0, 26)
	for li, u := range letterFlags {
		if u {
			letters = append(letters, li)
		}
	}
	directions := []string{"lrdu", "rlud", "ludr", "rdul", "lrdu", "lurd", "ulrd", "lrud"}

	for _, dir := range directions {
		filled := false
		/*
		fmt.Println(fmt.Sprintf("DIR: %s", dir))
		*/
		minRDir := make([]int, 26)
		minCDir := make([]int, 26)
		maxRDir := make([]int, 26)
		maxCDir := make([]int, 26)
		for i := 0; i < 26; i++ {
			minRDir[i], minCDir[i], maxRDir[i], maxCDir[i] = minR[i], minC[i], maxR[i], maxC[i]
		}
		cakeDir := make([][]rune, len(in.cake))
		for r, s := range in.cake {
			cakeDir[r] = make([]rune, len(s))
			for c, l := range s {
				cakeDir[r][c] = l
			}
		}

		/*
		for _, line := range cakeDir {
			fmt.Println(fmt.Sprintf("%s", string(line)))
		}
		*/
		for widen := 0; widen <= 25; widen++ {
			changed := false
			for _, d := range ([]rune(dir)) {
				for _, l := range letters {
					/*
					prefix := fmt.Sprintf("%c: [%d:%d %d:%d)", 'A' + l, minRDir[l], minCDir[l], maxRDir[l], maxCDir[l])
					*/
					scanChanged := false
					if (d == 'u') {
						for minRDir[l] > 0 &&
							allFree(minRDir[l] - 1, minCDir[l], minRDir[l], maxCDir[l], &cakeDir) {
							mark(minRDir[l] - 1, minCDir[l], minRDir[l], maxCDir[l], &cakeDir, l)
							minRDir[l] -= 1
							scanChanged = true
							changed = true
						}
					} else if (d == 'l') {
						for minCDir[l] > 0 &&
							allFree(minRDir[l], minCDir[l] - 1, maxRDir[l], minCDir[l], &cakeDir) {
							mark(minRDir[l], minCDir[l] - 1, maxRDir[l], minCDir[l], &cakeDir, l)
							minCDir[l] -= 1
							scanChanged = true
							changed = true
						}
					} else if (d == 'd') {
						for maxRDir[l] < in.R &&
							allFree(maxRDir[l], minCDir[l], maxRDir[l] + 1, maxCDir[l], &cakeDir) {
							mark(maxRDir[l], minCDir[l], maxRDir[l] + 1, maxCDir[l], &cakeDir, l)
							maxRDir[l] += 1
							scanChanged = true
							changed = true
						}
					} else if (d == 'r') {
						for maxCDir[l] < in.C &&
							allFree(minRDir[l], maxCDir[l], maxRDir[l], maxCDir[l] + 1, &cakeDir) {
							mark(minRDir[l], maxCDir[l], maxRDir[l], maxCDir[l] + 1, &cakeDir, l)
							maxCDir[l] += 1
							scanChanged = true
							changed = true
						}
					}
					if scanChanged {
						/*
						fmt.Print(fmt.Sprintf("%s --> [%d:%d %d:%d)\n", prefix, minRDir[l], minCDir[l], maxRDir[l], maxCDir[l]))
						*/
					}
				}
			}
			if (!changed) {
				break
			}
			/*
			for _, line := range cakeDir {
				fmt.Println(fmt.Sprintf("%s", string(line)))
			}
			*/
		}

		filled = true
		for _, line := range cakeDir {
			/*
			fmt.Println(fmt.Sprintf("%s", string(line)))
			*/
			for _, l := range line {
				if l == '?' {
					filled = false;
				}
			}
		}

		if (filled) {
			for r, s := range cakeDir {
				for c, l := range s {
					in.cake[r][c] = l
				}
			}

			break
		}
	}

	return Out{in.index, in.cake}
}

func solveChannel(ins <-chan In, outs chan <- Out) {
	for in := range ins {
		outs <- solve(in)
	}
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer reader.Close()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)

	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	caseCount := int(ReadInt32(scanner))

	cores := runtime.NumCPU()
	var ins = make(chan In, cores)
	var outs = make(chan Out, caseCount)
	for t := 0; t < cores; t++ {
		go solveChannel(ins, outs)
	}

	outsSlice := make([]Out, caseCount)
	for index := 0; index < caseCount; index++ {
		r, c := int(ReadInt16(scanner)), int(ReadInt16(scanner))
		cake := make([][]rune, r)
		for i := range cake {
			scanner.Scan()
			cake[i] = []rune(scanner.Text())
		}
		in := In{index, r, c, cake}
		ins <- in
	}
	close(ins)

	for index := 0; index < caseCount; index++ {
		out := <-outs
		outsSlice[out.index] = out
	}
	close(outs)

	for _, out := range outsSlice {
		Writef(writer, "Case #%d:\n", out.index + 1)
		for _, l := range out.cake {
			Writef(writer, "%s\n", string(l))
		}
	}
}

//	boring IO
func ReadInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return int64(res)
}

func ReadInt32(sc *bufio.Scanner) int32 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(res)
}

func ReadInt16(sc *bufio.Scanner) int16 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(res)
}

func ReadInt8(sc *bufio.Scanner) int8 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 8)
	if err != nil {
		panic(err)
	}
	return int8(res)
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
