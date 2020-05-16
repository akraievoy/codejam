package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	routes = make(map[[2]int64]string)
)

func probe(x, y, delta int64, route string, stepsLeft int64) {
	pos := [2]int64{x, y}
	if stepsLeft == 0 {
		_, prs := routes[pos]
		if prs {
			return
		}
		routes[pos] = route
		return
	}
	probe(x+delta, y, delta<<1, route+"E", stepsLeft-1)
	probe(x-delta, y, delta<<1, route+"W", stepsLeft-1)
	probe(x, y+delta, delta<<1, route+"N", stepsLeft-1)
	probe(x, y-delta, delta<<1, route+"S", stepsLeft-1)
}

func main() {
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for steps := int64(1); steps <= 8; steps++ {
		probe(0, 0, 1, "", steps)
	}

	Writef(writer, "%d\n", 201*201-1)

	caseNumber := 1
	for x := int64(-100); x <= 100; x++ {
		for y := int64(-100); y <= 100; y++ {
			if x == 0 && y == 0 {
				continue
			}
			pos := [2]int64{x, y}
			route, prs := routes[pos]
			if prs {
				Writef(writer, "Case #%d: %s\n", caseNumber, route)
			} else {
				Writef(writer, "Case #%d: IMPOSSIBLE\n", caseNumber)
			}
			caseNumber += 1
		}
	}
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
