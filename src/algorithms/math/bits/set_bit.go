package bits

// SetBit sets bit at specific position.
//
// bitPosition - zero based.
func SetBit(number, bitPosition int) int {
	return number | (0b0001 << bitPosition)
}
