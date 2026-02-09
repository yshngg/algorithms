package bits

// GetBit gets bit at specific position.
//
// bitPosition - zero based.
func GetBit(number, bitPosition int) int {
	return (number >> bitPosition) & 1
}
