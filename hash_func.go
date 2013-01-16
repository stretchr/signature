package signature

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

// HashFunc represents funcs that can hash a string.
type HashFunc func(s string) string

// Hash hashes a string using the current HashFunc.
//
// To tell Signature to use a different hashing algorithm, you
// just need to assign a different HashFunc to the Hash variable.
//
// To use the MD5 hash:
//
//     signature.Hash = signature.MD5Hash
//
// Or you can write your own hashing function:
//
//     signature.Hash = func(s string) string {
//	     // TODO: do your own hashing here
//     }
var Hash HashFunc = SHA1Hash

// SHA1Hash hashes a string using the SHA-1 hash algorithm as defined in RFC 3174.
var SHA1Hash HashFunc = func(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// MD5Hash hashes a string using the MD5 hash algorithm as defined in RFC 1321.
var MD5Hash HashFunc = func(s string) string {
	md5 := md5.New()
	io.WriteString(md5, s)
	return fmt.Sprintf("%x", md5.Sum(nil))
}
