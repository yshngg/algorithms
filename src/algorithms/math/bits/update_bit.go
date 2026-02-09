package bits

// UpdateBit updates bit at specific position.
//
// bitPosition - zero based.
// bitValue - 0 or 1.
// func UpdateBit(number, bitPosition, bitValue int) int {
// 	if number>>bitPosition&0b0001 == bitValue {
// 		return number
// 	}

// 	if number>>bitPosition&0b0001 > 0 {
// 		return number & ^(0b0001 << bitPosition)
// 	}

// 	return number | 0b0001<<bitPosition
// }

// UpdateBit updates bit at specific position.

// bitPosition - zero based.
// bitValue - 0 or 1.
func UpdateBit(number, bitPosition, bitValue int) int {
	return number&^(1<<bitPosition) | bitValue<<bitPosition
}
