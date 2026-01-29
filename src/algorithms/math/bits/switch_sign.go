package bits

// SwitchSign switch the sign of the number using twos complement approach.
func SwitchSign(number int) int {
	return ^number + 1
}
