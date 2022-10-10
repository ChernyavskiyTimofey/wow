package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	message := NewMessage(GrandType, []byte("test"))

	bs, err := message.Encode()
	require.NoError(t, err)

	newMessage, err := ParseMessage(bs)
	require.NoError(t, err)

	assert.Equal(t, message.Header, newMessage.Header)
	assert.Equal(t, message.Body, newMessage.Body)
}
