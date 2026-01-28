package bits

// IsEven detects if a number is even.
func IsEven(number int) bool {
	return number&0b0001 == 0
}
