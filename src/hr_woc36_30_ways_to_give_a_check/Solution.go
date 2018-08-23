package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tile uint8

const (
	TileEmpty       = Tile(iota)
	TileKingWhite
	TileQueenWhite
	TileKnightWhite
	TileBishopWhite
	TileRookWhite
	TilePawnWhite
	TileKingBlack
	TileQueenBlack
	TileKnightBlack
	TileBishopBlack
	TileRookBlack
	TilePawnBlack
)

var (
	CodeToTile = map[rune]Tile{
		'#': TileEmpty,
		'K': TileKingWhite,
		'Q': TileQueenWhite,
		'N': TileKnightWhite,
		'B': TileBishopWhite,
		'R': TileRookWhite,
		'P': TilePawnWhite,
		'k': TileKingBlack,
		'q': TileQueenBlack,
		'n': TileKnightBlack,
		'b': TileBishopBlack,
		'r': TileRookBlack,
		'p': TilePawnBlack}
)

func (t Tile) String() string {
	for code, tile := range CodeToTile {
		if tile == t {
			return string(code)
		}
	}
	return "?"
}

type Board [8][8]Tile

func (b *Board) String() string {
	res := ""
	for _, row := range b {
		for _, tile := range row {
			res = res + tile.String()
		}
		res = res + "\n"
	}
	return res
}

type Pos struct {
	Row int8 //	Rank is 8 to 1 (as per input order), Row is 0 to 7
	Col int8 // File is 1 to 8 (as per input order), Col is 0 to 7
}

func (board *Board) moveWithPromotion(oldPos Pos, newPos Pos, newTile Tile) Board {
	newBoard := [8][8]Tile{}
	for row, rowTiles := range board {
		for col, tile := range rowTiles {
			newBoard[row][col] = tile
		}
	}

	newBoard[oldPos.Row][oldPos.Col] = TileEmpty
	newBoard[newPos.Row][newPos.Col] = newTile

	return newBoard
}

type Index struct {
	Board Board
	Index map[Tile][]Pos
}

func (board *Board) Index() Index {
	index := make(map[Tile][]Pos, len(CodeToTile))
	for _, tile := range CodeToTile {
		index[tile] = make([]Pos, 0, 2)
	}
	for row, rowTiles := range board {
		for col, tile := range rowTiles {
			index[tile] =
				append(
					index[tile],
					Pos{
						int8(row),
						int8(col)})
		}
	}
	return Index{*board, index}
}

func (board *Index) checkTestRook(rookishTiles []Tile, checkedKing Tile) bool {
	for _, tile := range rookishTiles {
		for _, pos := range board.Index[tile] {
			{
				upperRow := pos.Row - 1
				for upperRow >= 0 {
					tile := board.Board[upperRow][pos.Col]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					upperRow --
				}
			}
			{
				lowerRow := pos.Row + 1
				for lowerRow < 8 {
					tile := board.Board[lowerRow][pos.Col]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					lowerRow ++
				}
			}
			{
				leftCol := pos.Col - 1
				for leftCol >= 0 {
					tile := board.Board[pos.Row][leftCol]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					leftCol --
				}
			}
			{
				rightCol := pos.Col + 1
				for rightCol < 8 {
					tile := board.Board[pos.Row][rightCol]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					rightCol ++
				}
			}
		}
	}
	return false
}

func (board *Index) checkTestBishop(bishopishTiles []Tile, checkedKing Tile) bool {
	for _, tile := range bishopishTiles {
		for _, pos := range board.Index[tile] {
			{
				leftUpperRow := pos.Row - 1
				leftUpperCol := pos.Col - 1
				for leftUpperRow >= 0 && leftUpperCol >= 0 {
					tile := board.Board[leftUpperRow][leftUpperCol]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					leftUpperRow --
					leftUpperCol --
				}
			}
			{
				leftLowerRow := pos.Row + 1
				leftLowerCol := pos.Col - 1
				for leftLowerRow < 8 && leftLowerCol >= 0 {
					tile := board.Board[leftLowerRow][leftLowerCol]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					leftLowerRow ++
					leftLowerCol --
				}
			}
			{
				rightLowerRow := pos.Row - 1
				rightLowerCol := pos.Col + 1
				for rightLowerRow >= 0 && rightLowerCol < 8 {
					tile := board.Board[rightLowerRow][rightLowerCol]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					rightLowerRow --
					rightLowerCol ++
				}
			}
			{
				rightUpperRow := pos.Row + 1
				rightUpperCol := pos.Col + 1
				for rightUpperRow < 8 && rightUpperCol < 8 {
					tile := board.Board[rightUpperRow][rightUpperCol]
					if tile == checkedKing {
						return true
					}
					if tile != TileEmpty {
						break
					}
					rightUpperRow ++
					rightUpperCol ++
				}
			}
		}
	}
	return false
}

func (board *Index) checkTestPawn(pawnTile Tile, checkedKing Tile) bool {
	captureRowDelta := int8(-1)
	if pawnTile == TilePawnBlack {
		captureRowDelta = int8(1)
	}

	for _, pos := range board.Index[pawnTile] {
		capturedRow := pos.Row + captureRowDelta
		if capturedRow >= 0 && capturedRow < 8 {
			if pos.Col > 0 && board.Board[capturedRow][pos.Col-1] == checkedKing {
				return true
			}
			if pos.Col < 7 && board.Board[capturedRow][pos.Col+1] == checkedKing {
				return true
			}
		}
	}
	return false
}

func (board *Index) checkTestKnight(knightTile Tile, checkedKing Tile) bool {
	for _, pos := range board.Index[knightTile] {
		deltas := []Pos{
			{-1, -2}, {-2, -1},
			{-2, 1}, {-1, 2},
			{1, 2}, {2, 1},
			{2, -1}, {1, -2}}
		for _, delta := range deltas {
			capturedRow := pos.Row + delta.Row
			capturedCol := pos.Col + delta.Col
			if capturedRow >= 0 && capturedRow < 8 &&
				capturedCol >= 0 && capturedCol < 8 &&
				board.Board[capturedRow][capturedCol] == checkedKing {
				return true
			}
		}
	}
	return false
}

func (board *Index) isWhiteUnderCheck() bool {
	return board.checkTestPawn(TilePawnBlack, TileKingWhite) ||
		board.checkTestRook([]Tile{TileRookBlack, TileQueenBlack}, TileKingWhite) ||
		board.checkTestBishop([]Tile{TileBishopBlack, TileQueenBlack}, TileKingWhite) ||
		board.checkTestKnight(TileKnightBlack, TileKingWhite)
}

func (board *Index) isBlackUnderCheck() bool {
	return board.checkTestPawn(TilePawnWhite, TileKingBlack) ||
		board.checkTestRook([]Tile{TileRookWhite, TileQueenWhite}, TileKingBlack) ||
		board.checkTestBishop([]Tile{TileBishopWhite, TileQueenWhite}, TileKingBlack) ||
		board.checkTestKnight(TileKnightWhite, TileKingBlack)
}

type In struct {
	boards []*Board
}

type Out struct {
	promotionsWithCheckPerBoard []int
}

func solve(in In) (out Out) {
	promotionsWithCheckPerBoard := make([]int, len(in.boards))
	for boardIdx, board := range in.boards {
		indexed := board.Index()
		promotionsWithCheck := 0
		for _, pos := range indexed.Index[TilePawnWhite] {
			if pos.Row == 1 && indexed.Board[0][pos.Col] == TileEmpty {
				for _, tile := range []Tile{TileQueenWhite, TileRookWhite, TileBishopWhite, TileKnightWhite} {
					afterPromotionBoard := board.moveWithPromotion(pos, Pos{0, pos.Col}, tile)
					afterPromotionIndex := afterPromotionBoard.Index()
					if afterPromotionIndex.isWhiteUnderCheck() {
						continue
					}
					if afterPromotionIndex.isBlackUnderCheck() {
						promotionsWithCheck++
					}
				}
			}
		}
		promotionsWithCheckPerBoard[boardIdx] = promotionsWithCheck

	}
	return Out{promotionsWithCheckPerBoard}
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
	boards := make([]*Board, size)
	for boardIdx := range boards {
		board := Board{}
		for row := 0; row < 8; row++ {
			rowText := []rune(ReadString(scanner))
			for col := 0; col < 8; col++ {
				board[row][col] = CodeToTile[rowText[col]]
			}
		}
		boards[boardIdx] = &board
	}
	in := In{boards}

	out := solve(in)

	for _, promotionsWithCheck := range out.promotionsWithCheckPerBoard {
		Writef(writer, "%d\n", promotionsWithCheck)
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

