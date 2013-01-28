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

func TestHashWithPrivateKey(t *testing.T) {

	assert.Equal(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), "f09b5012ffea295d3c199c1db9beaeaff3755501")
	assert.NotEmpty(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("different private key")))
	assert.NotEmpty(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), HashWithKeys([]byte("some bytes to hash"), []byte("different public key"), []byte("private key")))
	assert.NotEmpty(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), HashWithKeys([]byte("different bytes to hash"), []byte("public key"), []byte("private key")))

}
