package hr_woc36_10_acid_naming

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type In struct {
	names []string
}

type Out struct {
	kinds []string
}

func solve(in In) (out Out) {
	kinds := make([]string, len(in.names))
	for idx, name := range in.names {
		kinds[idx] = solveForName(name)
	}
	return Out{kinds}
}

func solveForName(name string) string {
	if strings.HasSuffix(name, "ic") {
		if strings.HasPrefix(name, "hydro") {
			return "non-metal acid"
		}
		return "polyatomic acid"
	}
	return "not an acid"
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
	names := make([]string, size)
	for i := range names {
		names[i] = ReadString(scanner)
	}
	in := In{names}

	out := solve(in)

	for _, kind := range out.kinds {
		Writef(writer, "%s\n", kind)
	}
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

func ReadString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
