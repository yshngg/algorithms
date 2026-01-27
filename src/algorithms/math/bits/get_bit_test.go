package bits_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestGetBit(t *testing.T) {
	// 1 = 0b0001
	t.Run("0b0001", func(t *testing.T) {
		require.Equal(t, 1, bits.GetBit(0b0001, 0))
		require.Equal(t, 0, bits.GetBit(0b0001, 1))
	})

	// 2 = 0b0010
	t.Run("0b0010", func(t *testing.T) {
		require.Equal(t, 0, bits.GetBit(0b0010, 0))
		require.Equal(t, 1, bits.GetBit(0b0010, 1))
	})

	// 3 = 0b0011
	t.Run("0b0011", func(t *testing.T) {
		require.Equal(t, 1, bits.GetBit(0b0011, 0))
		require.Equal(t, 1, bits.GetBit(0b0011, 1))
	})

	// 10 = 0b1010
	t.Run("0b1010", func(t *testing.T) {
		require.Equal(t, 0, bits.GetBit(0b1010, 0))
		require.Equal(t, 1, bits.GetBit(0b1010, 1))
		require.Equal(t, 0, bits.GetBit(0b1010, 2))
		require.Equal(t, 1, bits.GetBit(0b1010, 3))
	})
}
