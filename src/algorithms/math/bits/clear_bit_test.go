package bits_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestClearBit(t *testing.T) {
	// 1 = 0b0001
	t.Run("0b0001", func(t *testing.T) {
		require.Equal(t, 0b0000, bits.ClearBit(0b0001, 0))
		require.Equal(t, 0b0001, bits.ClearBit(0b0001, 1))
		require.Equal(t, 0b0001, bits.ClearBit(0b0001, 2))
		require.Equal(t, 0b0001, bits.ClearBit(0b0001, 3))
	})

	// 10 = 0b1010
	t.Run("0b1010", func(t *testing.T) {
		require.Equal(t, 0b1010, bits.ClearBit(0b1010, 0))
		require.Equal(t, 0b1000, bits.ClearBit(0b1010, 1))
		require.Equal(t, 0b1010, bits.ClearBit(0b1010, 2))
		require.Equal(t, 0b0010, bits.ClearBit(0b1010, 3))
	})
}
