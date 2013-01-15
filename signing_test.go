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

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=df073ee4086eed5848d167871c7424937027728e", signed)

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=34f55c3a086c260098e75066b38ac42e33e8faab", signed)

}

func TestValidateSignature(t *testing.T) {

	var valid bool

	signed, _ := GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

	valid, _ = ValidateSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=qJWro1ZxLeToLjNr5Znfi2ZbD+o=", "ABC123", "ABC123-private-wrong")
	assert.Equal(t, false, valid, "2")

}

func TestNoBodyHashWhenNoBody(t *testing.T) {

	signed, _ := GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=bdf49047abf3c8e56de21e244bc24b1c2a6086a2", signed)

}
