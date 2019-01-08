package grains

import "errors"

// Square takes a number representing the a square on teh chessboard from 1-64
// and returns the number of grains on it.
func Square(square int) (uint64, error) {
	if square < 1 {
		return 0, errors.New("The squares begin at 1; any number below that is invalid")
	}
	if square > 64 {
		return 0, errors.New("The squares end at 64; any number above that is invalid")
	}
	// Bit shifting is nice, but won't get us the zero-th power of two
	if square == 1 {
		return 1, nil
	}

	// In this case, we're only dealing with powers of two, which can be handled by shifting
	// the bit over one.
	return uint64(2 << uint(square-2)), nil
}

// Total returns the total count of all squares on the board added together
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		// In this case, we know these units won't fail, so we can skip the error check
		squareAtI, _ := Square(i)
		total += squareAtI
	}

	return total
}