package signature

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	assert.Equal(t, SHA1Hash, Hash)
}

func TestSHA1Hash(t *testing.T) {

	assert.Equal(t, SHA1Hash("abc"), "a9993e364706816aba3e25717850c26c9cd0d89d")

}

func TestMD5Hash(t *testing.T) {

	assert.Equal(t, MD5Hash("abc"), "900150983cd24fb0d6963f7d28e17f72")

}
