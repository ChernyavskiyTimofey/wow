package quotes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection(t *testing.T) {
	coll := NewCollection()
	quote1 := coll.GetRandQuote()
	assert.NotEmpty(t, quote1)

	quote2 := coll.GetRandQuote()
	assert.NotEqualValues(t, quote1, quote2)
}
