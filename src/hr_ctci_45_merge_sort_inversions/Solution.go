package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	nums []int32
}

type Out struct {
	inversions int
}

func mergeSort(numsPtr, tempPtr *[]int32, l, r int) int {
	if (l >= r - 1) {
		return 0
	}
	var nums, temp = *numsPtr, *tempPtr

	mid := (l + r) / 2
	inversions := mergeSort(numsPtr, tempPtr, l, mid) + mergeSort(numsPtr, tempPtr, mid, r)
	i, j, k := l, mid, l

	for k := l ; k < r ; k++ {
		temp[k] = nums[k]
	}

	for i < mid || j < r {
		if i == mid {
			nums[k] = temp[j]
			j += 1
			k += 1
		} else if j == r {
			nums[k] = temp[i]
			i += 1
			k += 1
		} else if temp[i] <= temp[j] {
			nums[k] = temp[i]
			i += 1
			k += 1
		} else {
			inversions += mid - i

			nums[k] = temp[j]
			j += 1
			k += 1
		}
	}

	// println(fmt.Sprintf("%v %v [%d %d) %d", nums, temp, l, r , inversions))

	return inversions
}

func solve(in In) Out {
	temp := make([]int32, len(in.nums))
	return Out{mergeSort(&(in.nums), &temp, 0, len(in.nums))}
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

	tests := ReadInt16(scanner)
	for t := int16(0); t < tests; t++  {
		size := ReadInt32(scanner)
		nums := make([]int32, size)
		for i := range nums {
			nums[i] = ReadInt32(scanner)
		}
		in := In{nums}

		out := solve(in)

		Writef(writer, "%d\n", out.inversions)
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

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
