package main

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	totalOccupiedSquaresInFile := 0
	occupiedSquares, exists := cb[file]
	if !exists {
		return totalOccupiedSquaresInFile
	}

	for _, occupiedSquare := range occupiedSquares {
		if occupiedSquare {
			totalOccupiedSquaresInFile += 1
		}
	}

	return totalOccupiedSquaresInFile
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	totalOccupiedSquaresInRank := 0

	for _, occupiedSquaresInFile := range cb {
		if occupiedSquaresInFile[rank-1] {
			totalOccupiedSquaresInRank += 1
		}
	}

	return totalOccupiedSquaresInRank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0
	for range cb {
		totalSquares += 8
	}

	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalOccupied := 0
	for file := range cb {
		totalOccupied += CountInFile(cb, file)
	}

	return totalOccupied
}

func main() {

}
