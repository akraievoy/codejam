package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

type Link struct {
	from, into int
}

type In struct {
	n, s    int
	fingers []int
	links   []Link
}

type LinkSort []Link

func (ls LinkSort) Len() int {
	return len(ls)
}
func (ls LinkSort) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}
func (ls LinkSort) Less(i, j int) bool {
	lsi, lsj := ls[i], ls[j]
	if (lsi.from == lsj.from) {
		return lsi.into < lsj.into
	}
	return lsi.from < lsj.from
}

type Out struct {
	distances []int
}

func solve(in In) (out Out) {
	distances := make([]int, in.n)
	for i := range distances {
		distances[i] = -1
	}
	distances[in.s] = 0
	queue := append(make([]int, 0, in.n), in.s)
	queuePos := 0
	for queuePos < len(queue) {
		v := queue[queuePos]
		for l := in.fingers[v]; l < in.fingers[v+1]; l++ {
			into := in.links[l].into
			if distances[into] >= 0 {
				continue
			}
			distances[into] = distances[v] + 6

			queue = append(queue, into)
		}

		queuePos ++
	}

	return Out{distances}
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

	testCount := int(ReadInt64(scanner))
	for t := 0; t < testCount; t++ {
		n := int(ReadInt64(scanner))
		m := int(ReadInt64(scanner))
		links := make([]Link, 2 * m)
		for l := range links {
			if l % 2 == 0 {
				links[l] = Link{
					int(ReadInt64(scanner)) - 1,
					int(ReadInt64(scanner)) - 1}
			} else {
				links[l] = Link{
					links[l - 1].into,
					links[l - 1].from,
				}
			}
		}
		s := int(ReadInt64(scanner)) - 1

		sort.Sort(LinkSort(links))
		fingers := make([]int, n + 1)
		from := 0
		for f, l := range links {
			for from < l.from{
				from++
				fingers[from] = f
			}
		}
		for from < n {
			from++
			fingers[from] = 2*m
		}

		in := In{n: n, s: s, fingers: fingers, links: links}
		out := solve(in)

		first := true
		for i, d := range out.distances {
			if i != in.s {
				if (!first) {
					Writef(writer, " ")
				}
				Writef(writer, "%d", d)
				first = false
			}
		}
		Writef(writer, "\n")
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

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
