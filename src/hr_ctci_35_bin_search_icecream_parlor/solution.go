package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

type In struct {
	money  int16
	prices []int16
}

type Out struct {
	index1 int
	index2 int
}

type PriceIndex struct {
	price int16
	index int
}

type PriceIndexSort []PriceIndex

func (ops PriceIndexSort) Len() int {
	return len(ops)
}
func (ops PriceIndexSort) Swap(i, j int) {
	ops[i], ops[j] = ops[j], ops[i]
}
func (ops PriceIndexSort) Less(i, j int) bool {
	return ops[i].price < ops[j].price
}

func solve(in In) (out Out) {
	pricesSorted := make([]PriceIndex, len(in.prices))
	for i, price := range in.prices {
		pricesSorted[i] = PriceIndex{price, i}
	}
	sort.Sort(PriceIndexSort(pricesSorted))

	limit := in.money / 2
	for i := 0; i < len(pricesSorted) - 1 && pricesSorted[i].price <= limit; i++ {
		found :=
			i + 1 + sort.Search(
				len(pricesSorted) - i - 1,
				func(idx int) bool {
					return pricesSorted[i + 1 + idx].price >= in.money - pricesSorted[i].price
				})

		if (found < len(pricesSorted) && pricesSorted[found].price + pricesSorted[i].price == in.money) {
			smallerPriceIndex := pricesSorted[i].index
			largerPriceIndex := pricesSorted[found].index
			//	that stuff is requested to be returned by indexes ordering, not by prices
			if (smallerPriceIndex < largerPriceIndex) {
				return Out{smallerPriceIndex, largerPriceIndex}
			} else {
				return Out{largerPriceIndex, smallerPriceIndex}
			}
		}
	}

	return Out{-2, -2}
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

	t := ReadInt16(scanner)
	for i := int16(0); i < t; i++ {
		money := ReadInt16(scanner)
		priceLen := ReadInt16(scanner)
		prices := make([]int16, priceLen)
		for j := int16(0); j < priceLen; j++ {
			prices[j] = ReadInt16(scanner)
		}
		out := solve(In{money, prices})

		Writef(writer, "%d %d\n", out.index1 + 1, out.index2 + 1)
	}
}

//	boring IO

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
