package bits

// ClearBit clears bit at specific position.
//
// bitPosition - zero based.
func ClearBit(number, bitPosition int) int {
	return number & ^(0b0001 << bitPosition)
}
