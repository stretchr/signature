package signature

import (
	"crypto/sha1"
	"fmt"
)

// HashFunc represents funcs that can hash a string.
type HashFunc func(s string) string

// Hash hashes a string using the default (SHA1Hash) hasher func.
var Hash HashFunc = SHA1Hash

// SHA1Hash hashes a string using the SHA-1 algorithm.
var SHA1Hash HashFunc = func(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
