package bits_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yshngg/algorithms/src/algorithms/math/bits"
)

func TestIsEven(t *testing.T) {
	testCases := []struct {
		number int
		isEven bool
	}{
		{
			number: 0,
			isEven: true,
		},
		{
			number: 1,
			isEven: false,
		},
		{
			number: -1,
			isEven: false,
		},
		{
			number: 2,
			isEven: true,
		},
		{
			number: -2,
			isEven: true,
		},
		{
			number: 3,
			isEven: false,
		},
		{
			number: -3,
			isEven: false,
		},
		{
			number: 8,
			isEven: true,
		},
		{
			number: 9,
			isEven: false,
		},
		{
			number: 121,
			isEven: false,
		},
		{
			number: 122,
			isEven: true,
		},
		{
			number: 1201,
			isEven: false,
		},
		{
			number: 1202,
			isEven: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.number), func(t *testing.T) {
			require.Equal(t, tc.isEven, bits.IsEven(tc.number))
		})
	}
}
