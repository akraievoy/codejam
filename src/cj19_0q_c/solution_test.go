package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	t.Skip()

	seed := time.Now().UnixNano()
	fmt.Printf("my seed is %d\n", seed)
	rnd := rand.New(rand.NewSource(seed))

	cipherTextLen := 101
	cleartext := make([]int, 0, cipherTextLen)
	for l := 0; l < cipherTextLen; l++ {
		pos := rnd.Intn(len(cleartext)+1)
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
		if !alpha.ProbablyPrime(1000) {
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
	fmt.Print("\nCLEARTEXT:\n")
	for _, l := range cleartext {
		fmt.Printf("%c", rune('A'+l))
	}
	fmt.Print("\n\nCIPHER:\n")
	for i, l := range cleartext {
		if i + 1 == len(cleartext) {
			fmt.Printf("\n")
		} else if i == 0 {
			fmt.Printf("%v", big.NewInt(0).Mul(alphabet[l], alphabet[cleartext[i+1]]).String())
		} else {
			fmt.Printf(" %v", big.NewInt(0).Mul(alphabet[l], alphabet[cleartext[i+1]]).String())
		}
	}
}
