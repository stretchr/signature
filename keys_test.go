package signature

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestKeys_RandomKey(t *testing.T) {

	key := RandomKey()
	assert.Equal(t, 32, len(key))

	key = RandomKey(20)
	assert.Equal(t, 20, len(key))

}
