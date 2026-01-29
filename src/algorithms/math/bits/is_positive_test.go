package bits_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestIsPositive(t *testing.T) {
	tests := []struct {
		number int
		isEven bool
	}{
		{
			number: 1,
			isEven: true,
		},
		{
			number: 2,
			isEven: true,
		},
		{
			number: 3,
			isEven: true,
		},
		{
			number: 5665,
			isEven: true,
		},
		{
			number: 56644325,
			isEven: true,
		},
		{
			number: 0,
			isEven: false,
		},
		{
			number: -0,
			isEven: false,
		},
		{
			number: -1,
			isEven: false,
		},
		{
			number: -2,
			isEven: false,
		},
		{
			number: -126,
			isEven: false,
		},
		{
			number: -5665,
			isEven: false,
		},
		{
			number: -56644325,
			isEven: false,
		},
		{
			number: 1202,
			isEven: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.number), func(t *testing.T) {
			require.Equal(t, tt.isEven, bits.IsPositive(tt.number))
		})
	}
}
