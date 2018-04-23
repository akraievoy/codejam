package main

import (
	"testing"
	"time"
	"math/rand"
	"fmt"
)

type kb struct{ k, b uint32 }

func TestMed(t *testing.T) {
	testTries(32000, 2, t)
	testTries(16000, 3, t)
	testTries(8000, 10, t)
	testTries(4000, 100, t)
	testTries(2000, 1000, t)
}

func TestLarge(t *testing.T) {
	testTries(100, 100000, t)
}

func testTries(tries int, queryCount int, t *testing.T) {
	fmt.Printf("tries = %d queries = %d\n", tries, queryCount)
	queries := make([]query, queryCount)
	validRemoves := make([]addQuery, 0, queryCount)
	for try := 0; try < tries; try++ {
		if (queryCount < 10000 && try % 1000 == 0 || queryCount >= 10000 && try % 10 == 0) && try > 0 {
			fmt.Printf("\tdone for tries = %d...\n", try)
		}
		seed := time.Now().UnixNano() + int64(try)
		rnd := rand.New(rand.NewSource(seed))

		queries = queries[:0]
		validRemoves = validRemoves[:0]

		for i := 0; i < queryCount; i++ {
			var queType int
			if i * 2 >= queryCount {
				queType = 2 //	only queries second half, please
			} else {
				queType = rnd.Intn(3)
			}

			switch queType {
			case 0:
				que := addQuery{
					uint32(rnd.Intn(100000) + 1),
					uint32(rnd.Intn(100001)),
				}
				queries = append(queries, que)
				validRemoves = append(validRemoves, que)
			case 1:
				if len(validRemoves) == 0 {
					que := addQuery{
						uint32(rnd.Intn(100000) + 1),
						uint32(rnd.Intn(100001)),
					}
					queries = append(queries, que)
					validRemoves = append(validRemoves, que)
				} else {
					removed := rnd.Intn(len(validRemoves))
					que := removeQuery{validRemoves[removed].k, validRemoves[removed].b}
					queries = append(queries, que)
					validRemoves =
						append(
							validRemoves[:removed],
							validRemoves[removed+1:]...,
						)
				}
			case 2:
				switch rnd.Intn(3){
				case 0:
					queries = append(queries, askQuery{0 })
				case 1:
					queries = append(queries, askQuery{uint32(rnd.Intn(100001))})
				case 2:
					queries = append(queries, askQuery{100000 })
				}
			}
		}

		out := solveCase(caseInput{queries})

		askIdx := 0
		kbs := make([]kb, 0, 100000)
		for _, que := range queries {
			if que.isAdd() {
				kbs = append(kbs, kb{que.getK(), que.getB()})
			} else if que.isRemove() {
				remIdx := -1
				for i, kb := range kbs {
					if kb.k == que.getK() && kb.b == que.getB() {
						remIdx = i
						break
					}
				}
				kbs = append(kbs[:remIdx], kbs[remIdx+1:]...)
			} else if que.isAsk() {
				expected := uint32(0)
				for _, kb := range kbs {
					if que.getQ() % kb.k == kb.b % kb.k {
						expected += 1
					}
				}

				{
					actual := out.intersections[askIdx]

					if actual != expected {
						dumpAskIdx := 0
						for dumpI, dumpQue := range queries {
							if dumpQue.isAdd() {
								fmt.Printf("%d\t+ %d %d\n", dumpI, dumpQue.getK(), dumpQue.getB())
							} else if dumpQue.isRemove() {
								fmt.Printf("%d\t- %d %d\n", dumpI, dumpQue.getK(), dumpQue.getB())
							} else {
								thisFails := ""
								if dumpAskIdx == askIdx {
									thisFails = " <<< THIS FAILS"
								}
								fmt.Printf("%d\t? %d%s\n", dumpI, dumpQue.getQ(), thisFails)
								dumpAskIdx += 1
							}
						}

						t.Errorf(
							"seed=%d\nout.intersections[%d] is not correct:\n\texpected %v\n\tactual %v",
							seed,
							askIdx,
							expected,
							actual,
						)

						t.FailNow()
					}
				}

				askIdx += 1
			}
		}
	}
}
