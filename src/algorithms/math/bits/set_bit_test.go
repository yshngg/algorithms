package bits_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestSetBit(t *testing.T) {
	// 1 = 0b0001
	t.Run("0b0001", func(t *testing.T) {
		require.Equal(t, 0b0001, bits.SetBit(0b0001, 0))
		require.Equal(t, 0b0011, bits.SetBit(0b0001, 1))
		require.Equal(t, 0b0101, bits.SetBit(0b0001, 2))
		require.Equal(t, 0b1001, bits.SetBit(0b0001, 3))
	})

	// 10 = 0b1010
	t.Run("0b1010", func(t *testing.T) {
		require.Equal(t, 0b1011, bits.SetBit(0b1010, 0))
		require.Equal(t, 0b1010, bits.SetBit(0b1010, 1))
		require.Equal(t, 0b1110, bits.SetBit(0b1010, 2))
		require.Equal(t, 0b1010, bits.SetBit(0b1010, 3))
	})
}
