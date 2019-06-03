package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

var DEBUG = false

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF

	T := l()
	for t := int64(1); t <= T; t++ {
		d("TEST %d\n", t)
		n := l()
		c, j := make([]int64, n, n), make([]int64, n, n)
		for i := int64(0); i < n; i++ {
			c[i], j[i] = l(), l()
		}
		valid := true
		hiJ, hiC := int64(1), int64(math.MinInt32)
		loJ, loC := int64(1), int64(0)
		for wee := int64(0); wee < n - 1 && valid; wee++ {
			for jay := int64(wee+1); jay < n && valid; jay++ {
				dC, dJ := c[jay]-c[wee], j[jay]-j[wee]
				d("lo=%d/%d hi=%d/%d: %d<%d dC=%d dJ=%d", -loC, loJ, -hiC, hiJ, wee, jay, dC, dJ)
				if dJ == 0 {
					if dC <= 0 {
						d("INVALID")
						valid = false
					}
				} else /*if dC == 0 {
					if dJ < 0 {
						d("INVALID")
						valid = false
					}
				} else */ if dJ > 0 {
					d(" > %d/%d (%s)", -dC, dJ, ""/*big.NewRat(-dC, dJ).FloatString(9)*/)
					/* -loC/loJ < -dC/dJ */
					if compare(-dC, dJ, -loC, loJ) > 0 {
						loC, loJ = dC, dJ
						d("NEW LO: %d/%d (%s)", -dC, dJ,""/*big.NewRat(-dC, dJ).FloatString(9)*/)
					} else {
						d("STALE LO: %d/%d (%s)", -dC, dJ, ""/*big.NewRat(-dC, dJ).FloatString(9)*/)
					}
				} else {
					d(" < %d/%d (%s)", -dC, dJ, ""/*big.NewRat(-dC, dJ).FloatString(9)*/)
					/* -dC/dJ < -hiC/hiJ */
					if compare(-dC, dJ, -hiC, hiJ) < 0 {
						hiC, hiJ = dC, dJ
						d("NEW HI: %d/%d (%s)", -dC, dJ, ""/*big.NewRat(-dC, dJ).FloatString(9)*/)
					} else {
						d("STALE HI: %d/%d (%s)", -dC, dJ, ""/*big.NewRat(-dC, dJ).FloatString(9)*/)
					}
				}
				d("\n")
			}
		}
		if valid && compare(-loC, loJ, -hiC, hiJ) < 0 {
			d("approx(%d/%d .. %d/%d) --> ", -loC, loJ, -hiC, hiJ)
			approx := BestApproximate(big.NewRat(-loC, loJ), big.NewRat(-hiC, hiJ))
			if approx == nil {
				p("Case #%d: IMPOSSIBLE\n", t)
			} else {
				approxNum, approxDenom := approx.Denom().Int64(), approx.Num().Int64()
				d("%d/%d\n", approxNum, approxDenom)
				p("Case #%d: %d %d\n", t, approxNum, approxDenom)
			}
		} else {
			p("Case #%d: IMPOSSIBLE\n", t)
		}
	}

	if false {
		d("%v", []interface{}{s, l, ui, f, d, p, pf})
	}
}

func compare(aNum, aDenom, bNum, bDenom int64) int {
	res := aNum*bDenom - bNum*aDenom
	if res == 0 {
		return 0
	}
	if (res > 0) == (aDenom * bDenom > 0) {
		return 1
	} else {
		return -1
	}
}

func main() {
	jam, closeFunc := JamNew()
	defer closeFunc()
	solveAll(jam)
}

type Jam interface {
	Scanner() *bufio.Scanner
	Writer() *bufio.Writer
	Close()

	Str() string
	Long() int64
	Int() uint32
	Float() float64

	D(format string, values ...interface{})
	P(format string, values ...interface{})
	PF(format string, values ...interface{})
}

func JamNew() (Jam, func()) {
	if len(os.Args) > 1 {
		panic("running with input file path is not supported")
	}

	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriterSize(os.Stdout, 1024*1024)
	jam := &jam{scanner, writer}
	return jam, jam.Close
}

type jam struct {
	sc *bufio.Scanner
	wr *bufio.Writer
}

func (j *jam) Close() {
	if err := j.wr.Flush(); err != nil {
		panic(err)
	}
}

func (j *jam) Scanner() *bufio.Scanner {
	return j.sc
}

func (j *jam) Writer() *bufio.Writer {
	return j.wr
}

func (j *jam) Str() string {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	return j.sc.Text()
}

func (j *jam) Long() int64 {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(j.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (j *jam) Int() uint32 {
	return uint32(j.Long())
}

func (j *jam) Float() float64 {
	j.sc.Scan()
	res, err := strconv.ParseFloat(j.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (j *jam) D(format string, values ...interface{}) {
	//noinspection GoBoolExpressions
	if !DEBUG {
		return
	}
	_ /*bytesWritten*/, err := fmt.Fprintf(os.Stderr, format, values...)
	if err != nil {
		panic(err)
	}
}

func (j *jam) P(format string, values ...interface{}) {
	_, err := fmt.Fprintf(j.wr, format, values...)
	if err != nil {
		panic(err)
	}
}

func (j *jam) PF(format string, values ...interface{}) {
	_, err := fmt.Fprintf(j.wr, format, values...)
	if err != nil {
		panic(err)
	}
	if err = j.wr.Flush(); err != nil {
		panic(err)
	}
}

type ContFrac []*big.Int

func (cf ContFrac) String() string {
	b := new(bytes.Buffer)
	wr := bufio.NewWriterSize(b, 512)
	for i, ii := range cf {
		if _, err := fmt.Fprintf(wr, "%v", ii); err != nil {
			panic(err)
		}
		if i == 0 {
			if _, err := fmt.Fprintf(wr, ";"); err != nil {
				panic(err)
			}
		} else if i + 1 < len(cf) {
			if _, err := fmt.Fprintf(wr, ","); err != nil {
				panic(err)
			}
		}
	}
	if err := wr.Flush(); err != nil {
		panic(err)
	}
	return b.String()
}

// https://www.math.uci.edu/~ndonalds/math180b/1contfrac.pdf

func RatToContFrac(a *big.Rat) ContFrac {
	zero := big.NewInt(0)
	res := make([]*big.Int, 0, 1)

	num := big.NewInt(0).Abs(a.Num())
	denom := big.NewInt(0).Set(a.Denom())
	for denom.Cmp(zero) > 0 {
		div, newDenom := big.NewInt(0).DivMod(num, denom, num)
		num, denom = denom, newDenom
		if len(res) == 0 && a.Denom().Sign() < 0 {
			res = append(res, big.NewInt(0).Neg(div))
		} else {
			res = append(res, div)
		}
	}

	if len(res) == 0 {
		res = append(res, big.NewInt(0))
	}

	return res
}

func ContFracToRat(cf ContFrac) *big.Rat {
	var res *big.Rat = nil
	for i := len(cf) - 1; i >= 0; i-- {
		if res == nil {
			res = big.NewRat(1,1).SetInt(cf[i])
		} else {
			l := big.NewRat(1, 1).SetInt(cf[i])
			r := res.Inv(res)
			res = l.Add(l, r)
		}
	}
	return res
}

func TrailingOne(a ContFrac) ContFrac {
	one := big.NewInt(1)
	last := len(a) - 1
	if len(a) == 0 || a[last].Cmp(one) == 0 && len(a) > 1{
		return a
	}

	res := make([]*big.Int, len(a) + 1, len(a) + 1)
	for i := range a {
		res[i] = a[i]
	}
	res[last] = big.NewInt(0).Sub(a[last], one)
	res[len(a)] = one

	return res
}

// https://en.wikipedia.org/wiki/Continued_fraction#Best_rational_within_an_interval

func BestApproximate(lowerExclusive, upperExclusive *big.Rat) *big.Rat {
	var best *big.Rat = nil
    lecf := RatToContFrac(lowerExclusive)
    uecf := RatToContFrac(upperExclusive)
	for _, l := range []ContFrac{lecf, TrailingOne(lecf)} {
		for _, u := range []ContFrac{uecf, TrailingOne(uecf)} {
			//fmt.Printf("%v .. %v -> ", l, u)
			pos, approx := 0, make([]*big.Int, 0, 0)
			for pos < len(l) && pos < len(u) && l[pos].Cmp(u[pos]) == 0 {
				approx = append(approx, l[pos])
				pos += 1
			}
			var last *big.Int = nil
			if pos < len(l) && pos < len(u) {
				if l[pos].Cmp(u[pos]) < 0 {
					last = l[pos]
				} else {
					last = u[pos]
				}
			} else if pos < len(l) {
				last = l[pos]
			} else if pos < len(u) {
				last = u[pos]
			} else {
				// possible only if lowerExclusive is strictly equal to upperExclusive
				continue
			}
			approx = append(approx, big.NewInt(0).Add(last, big.NewInt(1)))
			approxRat := ContFracToRat(approx)
			//fmt.Printf(
			//	"%v (%v<%d/%d<%v)\n",
			//	approx, lowerExclusive, approxRat.Num().Int64(), approxRat.Denom().Int64(), upperExclusive,
			//)
			if approxRat.Cmp(lowerExclusive) <= 0 || approxRat.Cmp(upperExclusive) >= 0 {
				continue
			}
			if best == nil {
				best = approxRat
			} else {
				denomCmp := best.Denom().Cmp(approxRat.Denom())
				if denomCmp > 0 || denomCmp == 0 && best.Num().Cmp(approxRat.Num()) > 0 {
					best = approxRat
				}
			}
		}
	}
	return best
}
