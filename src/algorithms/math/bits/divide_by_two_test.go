package bits_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestDivideByTwo(t *testing.T) {
	tests := []struct {
		number int
		want   int
	}{
		{
			number: 0,
			want:   0,
		},
		{
			number: 1,
			want:   0,
		},
		{
			number: -1,
			want:   -1,
		},
		{
			number: 3,
			want:   1,
		},
		{
			number: -3,
			want:   -2,
		},
		{
			number: 10,
			want:   5,
		},
		{
			number: -10,
			want:   -5,
		},
		{
			number: 17,
			want:   8,
		},
		{
			number: -17,
			want:   -9,
		},
		{
			number: 125,
			want:   62,
		},
		{
			number: -125,
			want:   -63,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.number), func(t *testing.T) {
			require.Equal(t, tt.want, bits.DivideByTwo(tt.number))
		})
	}
}
