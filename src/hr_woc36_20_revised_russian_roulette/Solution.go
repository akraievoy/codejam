package hr_woc36_20_revised_russian_roulette

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	doors []bool
}

type Out struct {
	movesMin   int
	movesMax   int
}

func solve(in In) (out Out) {
	movesMin := 0
	movesMax := 0
	prevDoors := 0
	for _, door := range in.doors {
		if door {
			//	most INEFFICIENT strategy opens rightmost closed door, each time
			movesMax += 1
			//  most efficient, OTOH, opens leftmost door,
			// 		which also opens all even-numbered doors in the same contiguous block
			//	hence we have to count only odd-numbered doors as even-numbered are open by preceding odd ones
			if prevDoors % 2 == 0 {
				movesMin += 1
			}
			prevDoors += 1
		} else {
			prevDoors = 0
		}
	}
	return Out{movesMin, movesMax}
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

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	size := int16(ReadInt(scanner))
	doors := make([]bool, size)
	for i := range doors {
		doors[i] = ReadInt(scanner) > 0
	}
	in := In{doors}

	out := solve(in)

	Writef(writer, "%d %d\n", out.movesMin, out.movesMax)
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
