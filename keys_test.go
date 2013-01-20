package signature

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestKeys_RandomKey(t *testing.T) {

	var key string

	key = RandomKey(20)
	assert.Equal(t, 20, len(key))

	key = RandomKey(50)
	assert.Equal(t, 50, len(key))

	key = RandomKey(2)
	assert.Equal(t, 2, len(key))

	key = RandomKey(1)
	assert.Equal(t, 1, len(key))

	key = RandomKey(0)
	assert.Equal(t, 0, len(key))

}
