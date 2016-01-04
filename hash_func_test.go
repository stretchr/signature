package signature

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func TestHash(t *testing.T) {
	assert.Equal(t, SHA1Hash, Hash)
}
*/

func TestSHA1Hash(t *testing.T) {

	assert.Equal(t, SHA1Hash("abc"), "a9993e364706816aba3e25717850c26c9cd0d89d")

}

func TestMD5Hash(t *testing.T) {

	assert.Equal(t, MD5Hash("abc"), "900150983cd24fb0d6963f7d28e17f72")

}

func TestHashWithPrivateKey(t *testing.T) {

	assert.Equal(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), "4c9ac9d6594e19c0bcbfcc261f82e190cf222526")
	assert.NotEmpty(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("different private key")))
	assert.NotEmpty(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), HashWithKeys([]byte("some bytes to hash"), []byte("different public key"), []byte("private key")))
	assert.NotEmpty(t, HashWithKeys([]byte("some bytes to hash"), []byte("public key"), []byte("private key")), HashWithKeys([]byte("different bytes to hash"), []byte("public key"), []byte("private key")))

}
