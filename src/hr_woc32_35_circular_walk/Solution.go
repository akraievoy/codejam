package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	n, s, t        int
	r0, g, seed, p int
}

type Out struct {
	seconds int
}

var debug = false

func within(l, r, x int) bool {
	if l == r {
		return true
	}

	if l < r {
		return l <= x && x < r
	}

	return l <= x || x < r
}

func merge(l1, r1, l2, r2 int) (int, int) {
	if l2 == r1 {
		if (debug) {
			fmt.Println(fmt.Sprintf("[%d %d) [%d %d) -> [%d %d)", l1, r1, l2, r2, l1, r2))
		}
		return l1, r2
	} else if l1 == r2 {
		if (debug) {
			fmt.Println(fmt.Sprintf("[%d %d) [%d %d) -> [%d %d)", l1, r1, l2, r2, l2, r1))
		}
		return l2, r1
	} else if within(l1, r1, l2) {
		if within(l1, r1, r2) || r1 == r2 {
			if debug {
				fmt.Println(fmt.Sprintf("[%d %d) [%d %d) -> [%d %d)", l1, r1, l2, r2, l1, r1))
			}
			return l1, r1
		} else if within(l1, r2, r1) {
			if debug {
				fmt.Println(fmt.Sprintf("[%d %d) [%d %d) -> [%d %d)", l1, r1, l2, r2, l1, r2))
			}
			return l1, r2
		} else {
			panic(fmt.Sprintf("[%d %d) [%d %d) won't merge - [l1 r1) has l2", l1, r1, l2, r2))
		}
	} else if within(l2, r2, l1) {
		if within(l2, r2, r1) || r2 == r1 {
			if debug {
				fmt.Println(fmt.Sprintf("[%d %d) [%d %d) -> [%d %d)", l1, r1, l2, r2, l2, r2))
			}
			return l2, r2
		} else if within(l2, r1, r2) {
			if debug {
				fmt.Println(fmt.Sprintf("[%d %d) [%d %d) -> [%d %d)", l1, r1, l2, r2, l2, r1))
			}
			return l2, r1
		} else {
			panic(fmt.Sprintf("[%d %d) [%d %d) won't merge - [l2 r2) has l1", l1, r1, l2, r2))
		}
	} else {
		panic(fmt.Sprintf("[%d %d) [%d %d) won't merge at all", l1, r1, l2, r2))
	}
}

func norm(l, r, n int) (int, int) {
	if l < 0 {
		l += n
	}
	return l, r % n
}

func solve(in In) (out Out) {
	r := make([]int, in.n)
	r[0] = in.r0

	for i := 1; i < len(r); i++ {
		r[i] = (r[i - 1] * in.g + in.seed) % in.p
	}

	if debug {
		fmt.Println(fmt.Sprintf("r=%v", r))
	}

	curL, curR := in.s, in.s + 1
	seenL, seenR := in.s, in.s + 1
	time := 0

	for !within(curL, curR, in.t) {
		nextL, nextR := curL, curR

		if time == 0 {
			if curR <= curL {
				curR += in.n
			}
			for i := curL; i < curR; i++ {
				nextL, nextR = merge(nextL, nextR, (i - r[i%in.n] + in.n) % in.n, (i + r[i%in.n] + 1) % in.n)
			}
		} else {
			if curL > seenL {
				seenL += in.n
			}
			if seenR > curR {
				curR += in.n
			}
			for i := curL; i < seenL; i++ {
				nextL, nextR = merge(nextL, nextR, (i - r[i%in.n] + in.n) % in.n, (i + r[i%in.n] + 1) % in.n)
			}
			for i := seenR; i < curR; i++ {
				nextL, nextR = merge(nextL, nextR, (i - r[i%in.n] + in.n) % in.n, (i + r[i%in.n] + 1) % in.n)
			}
		}

		nextL, nextR = norm(nextL, nextR, in.n)
		curL, curR = norm(curL, curR, in.n)
		if nextL == curL && nextR == curR {
			if debug {
				fmt.Println(fmt.Sprintf("[%v %v) -> [%v %v) BREAK", seenL, seenR, curL, curR))
			}

			return Out{seconds: -1}
		}

		seenL, seenR = norm(curL, curR, in.n)
		curL, curR = norm(nextL, nextR, in.n)

		if debug {
			fmt.Println(fmt.Sprintf("[%v %v) -> [%v %v)", seenL, seenR, curL, curR))
		}

		time++
	}

	return Out{seconds: time}
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

	out := solve(In{
		n: ReadInt(scanner),
		s: ReadInt(scanner),
		t: ReadInt(scanner),
		r0: ReadInt(scanner),
		g: ReadInt(scanner),
		seed: ReadInt(scanner),
		p: ReadInt(scanner)})

	Writef(writer, "%d\n", out.seconds)
}

func ReadInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func ReadInt(sc *bufio.Scanner) int {
	return int(ReadInt64(sc))
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
