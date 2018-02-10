package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

var (
	DEBUG=false
)

type Runner struct {
	Height int32
	Price  int64
}

func (r *Runner) String() string {
	return fmt.Sprintf("%d@%d", r.Height, r.Price)
}

type In struct {
	Runners []Runner
}

type Out struct {
	MinTotalCost int64
}

type Route struct {
	Height       int32
	Cost         int64
}

func (r *Route) Premium() int64 {
	return r.Cost - int64(r.Height)
}

func (r *Route) String() string {
	return fmt.Sprintf("%d\t%d\t%d", r.Height, r.Cost, r.Premium())
}

func MinByPremium(routes []Route) Route {
	var minRoute Route
	for idx, route := range routes {
		if idx == 0 || minRoute.Premium() > route.Premium() {
			minRoute = route
		}
	}
	return minRoute
}

func solve(in In) (out Out) {
	routesByHeightDesc := make([]Route, 1, 1)
	routesByHeightDesc[0] =
		Route{
			in.Runners[0].Height,
			in.Runners[0].Price,
		}

	for pos, runner := range in.Runners {
		if pos == 0 {
			if DEBUG {
				dumpRoutes("INIT", routesByHeightDesc)
			}
			continue // init above does what we need for step 0
		}

		if DEBUG {
			fmt.Printf("\nRUNNER %v\n", &runner)
		}

		routeHighest := routesByHeightDesc[0]
		routeLowest := routesByHeightDesc[len(routesByHeightDesc)-1]

		if runner.Height > routeHighest.Height {
			//	new highest route: all routes collapse to one new single possibility
			ascendRoute := MinByPremium(routesByHeightDesc)
			newRoute :=
				Route{
					runner.Height,
					ascendRoute.Cost + runner.Price + int64(runner.Height-ascendRoute.Height),
				}
			routesByHeightDesc = routesByHeightDesc[:1]
			routesByHeightDesc[0] = newRoute
			if DEBUG {
				dumpRoutes("HIGHEST NEW", routesByHeightDesc)
			}
			continue // processed
		}

		if routeHighest.Height == runner.Height {
			//	alternative topmost runner
			ascendRoute := MinByPremium(routesByHeightDesc)
			ascendCost := ascendRoute.Cost + runner.Price + int64(runner.Height - ascendRoute.Height)
			if ascendCost < routeHighest.Cost {
				routesByHeightDesc[0] = Route{runner.Height, ascendCost}
				routesByHeightDesc = routesByHeightDesc[:1]
				if DEBUG {
					dumpRoutes("HIGHEST UPDATE", routesByHeightDesc)
				}
			} else {
				routesByHeightDesc = routesByHeightDesc[:1]
				if DEBUG {
					dumpRoutes("HIGHEST UPDATE NOOP", routesByHeightDesc)
				}
				continue //	can't be part of optimal route
			}
		} else if routeLowest.Height > runner.Height {
			//	new lowest route
			if runner.Price < 0 {
				newRoute :=
					Route{
						runner.Height,
						routeLowest.Cost + runner.Price + int64(routeLowest.Height-runner.Height),
					}
				routesByHeightDesc = append(routesByHeightDesc, newRoute)
				if DEBUG {
					dumpRoutes("LOWEST NEW", routesByHeightDesc)
				}
			} else {
				if DEBUG {
					dumpRoutes("LOWEST NEW NOOP", routesByHeightDesc)
				}
				continue // can't be part of optimal route
			}
		} else if routeLowest.Height == runner.Height {
			//	alternative lowest runner
			if runner.Price < 0 {
				routesByHeightDesc[len(routesByHeightDesc)-1] =
					Route{
						routeLowest.Height,
						routeLowest.Cost + runner.Price,
					}
				if DEBUG {
					dumpRoutes("LOWEST SWAP", routesByHeightDesc)
				}
			} else {
				if DEBUG {
					dumpRoutes("LOWEST SWAP NOOP", routesByHeightDesc)
				}
				continue // can't be part of optimal route
			}
		} else {
			routeIdx :=
				sort.Search(
					len(routesByHeightDesc),
					func(pos int) bool {
						return routesByHeightDesc[pos].Height <= runner.Height
					},
				)
			if routesByHeightDesc[routeIdx].Height == runner.Height {
				ascendRoute := MinByPremium(routesByHeightDesc[:routeIdx+1])
				ascendCost := ascendRoute.Cost + runner.Price + int64(runner.Height - ascendRoute.Height)

				routeSameHeight := routesByHeightDesc[routeIdx]
				if ascendCost < routeSameHeight.Cost {
					newRoute :=
						Route{
							runner.Height,
							ascendCost,
						}
					routesByHeightDesc[routeIdx] = newRoute
					routesByHeightDesc = routesByHeightDesc[:routeIdx+1]
					if DEBUG {
						dumpRoutes("MID SWAP", routesByHeightDesc)
					}
				} else {
					routesByHeightDesc = routesByHeightDesc[:routeIdx+1]
					if DEBUG {
						dumpRoutes("MID SWAP NOOP", routesByHeightDesc)
					}
					continue // can't be part of optimal route
				}
			} else {
				ascendRoute := MinByPremium(routesByHeightDesc[routeIdx:])
				routeHigher := routesByHeightDesc[routeIdx-1]

				costAscent := ascendRoute.Cost + int64(runner.Height-ascendRoute.Height)
				costDescent := routeHigher.Cost + int64(routeHigher.Height-runner.Height)
				minCost := min64(costAscent, costDescent) + runner.Price

				if minCost >= costDescent {
					routesByHeightDesc = routesByHeightDesc[:routeIdx]
					if DEBUG {
						dumpRoutes("MID NEW NOOP", routesByHeightDesc)
					}
					continue // can't be part of optimal route
				}

				newRoute :=
					Route{
						runner.Height,
						minCost,
					}

				routesByHeightDesc = append(routesByHeightDesc[:routeIdx], newRoute)
				if DEBUG {
					dumpRoutes("MID NEW", routesByHeightDesc)
				}
			}
		}
	}

	minTotalCost := int64(0)
	for idx, route := range routesByHeightDesc {
		if idx == 0 || minTotalCost > route.Cost {
			minTotalCost = route.Cost
		}
	}
	return Out{minTotalCost + int64(len(in.Runners))}
}
func dumpRoutes(action string, routesByHeightDesc []Route) {
	routesByHeightDescStr := fmt.Sprintf("%s:\n", action)
	for idx, route := range routesByHeightDesc {
		routesByHeightDescStr += fmt.Sprintf("%d:\t%v\n", idx, &route)
	}
	fmt.Print(routesByHeightDescStr)
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

	size := ReadInt(scanner)
	runners := make([]Runner, size)
	runners[0].Height = int32(ReadInt(scanner))
	for i := range runners {
		if i == 0 {
			continue
		}
		runners[i].Height = int32(ReadInt(scanner))
	}
	for i := range runners {
		if i == 0 {
			continue
		}
		runners[i].Price = int64(ReadInt(scanner))
	}
	in := In{runners}

	out := solve(in)

	Writef(writer, "%d\n", out.MinTotalCost)
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

func ReadFloat64(sc *bufio.Scanner) float64 {
	sc.Scan()
	res, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
