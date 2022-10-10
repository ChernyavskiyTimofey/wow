package hashcash

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	resource      string
	bits          uint
	presetCounter uint64
	iters         uint64
	err           error
}

var (
	date     = time.Date(2022, 10, 8, 12, 30, 0, 0, time.UTC)
	rnd      = "qG3zoQ/2wpw="
	maxIters = uint64(100_000_000)

	cases = []testCase{
		{"test1", 1, 0, 12, nil},
		{"test2", 2, 0, 5, nil},
		{"test3", 3, 0, 570, nil},
		{"test4", 4, 94000, 94733, nil},
		{"test5", 5, 500000, 506310, nil},
		{"test6", 6, 36000000, 36001844, nil},
		{"test7", 7, maxIters - 1, maxIters, ErrMaxIterations},
	}
)

func TestHash(t *testing.T) {
	t.Parallel()

	for _, c := range cases {
		t.Run(fmt.Sprintf("test for %s", c.resource), func(t *testing.T) {
			hash, err := NewHash(c.resource, c.bits)
			require.NoError(t, err)

			hash.Date = date
			hash.Rand = rnd
			hash.Counter = c.presetCounter

			_, err = hash.Compute(maxIters)
			if c.err == nil {
				require.NoError(t, err)
			} else {
				assert.EqualValues(t, c.err, err)
			}
			assert.EqualValues(t, c.iters, hash.Counter)
		})
	}
}
