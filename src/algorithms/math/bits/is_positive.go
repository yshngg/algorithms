package bits

// IsPositive detects if a number is positive.
func IsPositive(number int) bool {
	if number == 0 {
		return false
	}

	const intSize = 32<<(^uint(0)>>63) - 1
	return uint(number)>>intSize&1 == 0
}
