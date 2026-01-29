package bits_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestSwitchSign(t *testing.T) {
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
			want:   -1,
		},
		{
			number: -1,
			want:   1,
		},
		{
			number: 32,
			want:   -32,
		},
		{
			number: 23,
			want:   -23,
		},
	}
	for _, tt := range tests {
		t.Logf("%b", -23)
		t.Logf("%b", 23)
		t.Run(fmt.Sprintf("%d", tt.number), func(t *testing.T) {
			require.Equal(t, tt.want, bits.SwitchSign(tt.number))
		})
	}
}
