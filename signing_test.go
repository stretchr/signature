package signature

import (
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestGetSignature(t *testing.T) {

	var signed string

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "df073ee4086eed5848d167871c7424937027728e", signed)

	signed, _ = GetSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "df073ee4086eed5848d167871c7424937027728e", signed, "Lower case method shouldn't affect GetSignature")

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "34f55c3a086c260098e75066b38ac42e33e8faab", signed)

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?:name=!Laurie&~key=ABC123&:age=>20&:name=!Mat", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "34f55c3a086c260098e75066b38ac42e33e8faab", signed, "Different order of args shouldn't matter")

}

func TestGetSignedURL(t *testing.T) {

	var signed string

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20", "body", "DEF-public", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20&~key=DEF-public&~sign=2c96a0ed4fc0400ee781a120a4ea018a80ec9cf2", signed)

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-public", "DIFFERENT-PRIVATE")
	assert.Equal(t, "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20&~key=DIFFERENT-public&~sign=3e4618a4662fb69dd2853f0e52b4bb190c1c733d", signed)
}

func TestValidateSignature(t *testing.T) {

	var valid bool

	// Validate with signature at the end
	signed, _ := GetSignature("GET", "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20&~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

	// Validate with signature at the beginning
	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~sign=%s&:name=!Mat&:name=!Laurie&:age=>20", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

	// Validate with signature in the middle
	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?:name=!Mat&~sign=%s&:name=!Laurie&:age=>20", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

	// Validate with signature only
	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

	valid, _ = ValidateSignature("GET", "http://test.stretchr.com/api/v1:name=!Mat&:name=!Laurie&:age=>20&~sign=qJWro1ZxLeToLjNr5Znfi2ZbD+o=", "ABC123", "ABC123-private-wrong")
	assert.Equal(t, false, valid, "2")

	signed, _ = GetSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "3")

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("get", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "4")

}

func TestNoBodyHashWhenNoBody(t *testing.T) {

	signed, _ := GetSignedURL("GET", "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20", "", "DEF-public", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?:name=!Mat&:name=!Laurie&:age=>20&~key=DEF-public&~sign=19bf8c8fd1af5ba314c1b7a29ccc6c79c48d7bec", signed)

}
