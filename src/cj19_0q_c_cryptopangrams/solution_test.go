package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	t.Skip()

	testin, _ := os.Create("02.in")
	testout, _ := os.Create("02.out")

	_, _ = fmt.Fprintf(testin, "%d\n", 100)
	for tc := 1; tc <= 100; tc++ {
		seed := time.Now().UnixNano()
		fmt.Printf("seed for test %d is %d\n", tc, seed)
		rnd := rand.New(rand.NewSource(seed))

		cipherTextLen := 101
		cleartext := make([]int, 0, cipherTextLen)
		for l := 0; l < cipherTextLen; l++ {
			pos := rnd.Intn(len(cleartext) + 1)
			nextLetter := l
			if l > 25 {
				nextLetter = rnd.Intn(26)
			}
			cleartext = append(cleartext[:pos], append([]int{nextLetter}, cleartext[pos:]...)...)
		}

		alphabet := make([]*big.Int, 0, 26)
		primeLimit := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(100), nil)
		for len(alphabet) < 26 {
			alpha := big.NewInt(0).Rand(rnd, primeLimit)
			if !alpha.ProbablyPrime(100) {
				continue
			}

			searchIdx := sort.Search(
				len(alphabet),
				func(i int) bool { return alphabet[i].Cmp(alpha) >= 0 },
			)
			if searchIdx < len(alphabet) {
				if alphabet[searchIdx].Cmp(alpha) > 0 {
					alphabet = append(alphabet[:searchIdx], append([]*big.Int{alpha}, alphabet[searchIdx:]...)...)
				} else {
					continue
				}
			} else {
				alphabet = append(alphabet, alpha)
			}
		}

		fmt.Print("ALPHABET:\n")
		for a := 0; a < 25; a++ {
			fmt.Printf("%c = %s\n", rune('A'+a), alphabet[a].String())
		}

		_, _ = fmt.Fprintf(testout, "Case #%d: ", tc)
		for _, l := range cleartext {
			_, _ = fmt.Fprintf(testout, "%c", rune('A'+l))
		}
		_, _ = fmt.Fprintf(testout, "\n")

		_, _ = fmt.Fprintf(testin, "%s %d\n", primeLimit.String(), 100)
		for i, l := range cleartext {
			if i+1 == len(cleartext) {
				_, _ = fmt.Fprintf(testin, "\n")
			} else if i == 0 {
				_, _ = fmt.Fprintf(testin, "%v", big.NewInt(0).Mul(alphabet[l], alphabet[cleartext[i+1]]).String())
			} else {
				_, _ = fmt.Fprintf(testin, " %v", big.NewInt(0).Mul(alphabet[l], alphabet[cleartext[i+1]]).String())
			}
		}
	}
}
