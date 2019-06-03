package algo

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
)

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

func ContFracNew(ints ...int64) ContFrac {
	res := make([]*big.Int, len(ints), len(ints))
	for i, ii := range ints {
		res[i] = big.NewInt(ii)
	}
	return res
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