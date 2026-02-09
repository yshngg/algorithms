package bits_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestMultiplyByTwo(t *testing.T) {
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
			want:   2,
		},
		{
			number: -1,
			want:   -2,
		},
		{
			number: 3,
			want:   6,
		},
		{
			number: -3,
			want:   -6,
		},
		{
			number: 10,
			want:   20,
		},
		{
			number: -10,
			want:   -20,
		},
		{
			number: 17,
			want:   34,
		},
		{
			number: -17,
			want:   -34,
		},
		{
			number: 125,
			want:   250,
		},
		{
			number: -125,
			want:   -250,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.number), func(t *testing.T) {
			require.Equal(t, tt.want, bits.MultiplyByTwo(tt.number))
		})
	}
}
