package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"math"
	"time"
)

func readUint16(sc *bufio.Scanner) uint16 {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return uint16(res)
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}

const maxBlockModulo = 300

func blockFloor(block int, idx uint16) uint16 {
	return idx - (idx % uint16(block))
}

func blockCeil(block int, idx uint16, n int) uint16 {
	var blockCeil = blockFloor(block, idx) + uint16(block - 1)
	if int(blockCeil) >= n {
		return uint16(n - 1)
	}
	return blockCeil
}

type ei struct {
	elems   []uint16
	indexes []uint16
}

func (c ei) Len() int {
	return len(c.elems)
}
func (c ei) Swap(i, j int) {
	c.elems[i], c.elems[j] = c.elems[j], c.elems[i]
	c.indexes[i], c.indexes[j] = c.indexes[j], c.indexes[i]
}
func (c ei) Less(i, j int) bool {
	return c.elems[i] < c.elems[j]
}

func main() {
	var scanner *bufio.Scanner;
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if (err != nil) {
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

	var n, queries int = int(readUint16(scanner)), int(readUint16(scanner))
	var block = int(math.Floor(math.Sqrt(float64(n))))
	var blocks = 1 + n / block

	var a = make([]uint16, n)
	for i := range a {
		a[i] = readUint16(scanner)
	}

	rootStarted := time.Now().UnixNano()
	var divToModToBlockCounts = make([][][]uint16, maxBlockModulo + 1)
	for div := 1; div <= maxBlockModulo; div++ {
		divToModToBlockCounts[div] = make([][]uint16, div)
		for mod := 0; mod < div; mod++ {
			divToModToBlockCounts[div][mod] = make([]uint16, blocks, blocks)
		}
	}
	for i := range a {
		iBlock := i / block
		for div := 1; div <= maxBlockModulo; div++ {
			divToModToBlockCounts[div][int(a[i] % uint16(div))][iBlock] += 1
		}
	}
	rootFinished := time.Now().UnixNano()

	sortStarted := time.Now().UnixNano()
	var indexes [][]uint16 = make([][]uint16, 40001)
	for i := range indexes {
		indexes[i] = make([]uint16, 0, 0)
	}
	for i,v := range a {
		indexes[v] = append(indexes[v], uint16(i))
	}
	sortFinished := time.Now().UnixNano()

	var queriesRoot, queriesSort = 0, 0
	var queriesRootNano, queriesSortNano int64 = 0, 0
	for q := 0; q < queries; q++ {
		var left, right uint16 = readUint16(scanner), readUint16(scanner)
		var queryModulo, queryRemainder uint16 = readUint16(scanner), readUint16(scanner)

		res := 0
		if queryModulo <= maxBlockModulo {
			queryRootStarted := time.Now().UnixNano();
			leftFloor := blockFloor(block, left)
			for i := leftFloor; i < left; i++ {
				if a[i] % queryModulo == queryRemainder {
					res -= 1
				}
			}
			rightCeil := blockCeil(block, right, n)
			for i := rightCeil; i > right; i-- {
				if a[i] % queryModulo == queryRemainder {
					res -= 1
				}
			}
			bLeft := int(left) / block
			bRight := int(right) / block
			for b := bLeft; b <= bRight; b++ {
				res += int(divToModToBlockCounts[queryModulo][queryRemainder][b])
			}
			queryRootFinished := time.Now().UnixNano();
			queriesRootNano += (queryRootFinished - queryRootStarted)
			queriesRoot += 1
		} else {
			querySortStarted := time.Now().UnixNano()
			for search := int(queryRemainder); search <= 40000; search += int(queryModulo) {

				l := len(indexes[search])
				if (l == 0) {
					if (false) {
						os.Stderr.WriteString(
							fmt.Sprintf(
								"HOUSTON: %d --> %v SKIPPED\n",
								search, indexes[search]))
					}
					continue
				}
				start :=
					sort.Search(
						l,
						func(i int) bool {
						  return indexes[search][i] >= left
						})
				end :=
					start +
						sort.Search(
							l - start,
							func(i int) bool {
								return indexes[search][i+start] > right
							})
				if (false) {
					os.Stderr.WriteString(
						fmt.Sprintf(
							"HOUSTON: %d --> %v\n%d %d -> %d %d\n",
							search,
							indexes[search],
							left,
							right,
							start,
							end))
				}
				res += (end - start)
			}
			querySortFinished := time.Now().UnixNano()
			queriesSortNano += (querySortFinished - querySortStarted)
			queriesSort += 1
		}

		writef(writer, "%d\n", res)
	}

	/*
	os.Stderr.WriteString(
	*/
	fmt.Sprintf(
		"root constructed in %d millis\n" +
			"sort constructed in %d millis\n" +
			"%d root queries in %d millis (%f nanos each)\n" +
			"%d sort queries in %d millis (%f nanos each)\n" +
			"\n",
		(rootFinished - rootStarted) / 1000000,
		(sortFinished - sortStarted) / 1000000,
		queriesRoot, queriesRootNano / 1000000, float64(queriesRootNano) / float64(queriesRoot),
		queriesSort, queriesSortNano / 1000000, float64(queriesSortNano) / float64(queriesSort)) /* )
	*/
}
