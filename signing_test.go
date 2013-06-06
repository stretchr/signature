package signature

import (
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestGetSignature(t *testing.T) {

	var signed string

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "6c3dc03b3f85c9eb80ed9e4bd21e82f1bbda5b8d", signed)

	signed, _ = GetSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "6c3dc03b3f85c9eb80ed9e4bd21e82f1bbda5b8d", signed, "Lower case method shouldn't affect GetSignature")

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "2d6ad8d46cd8d08d5dfeb91a30dd4cd50f30eb36", signed)

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?:name=!Laurie&~key=ABC123&:age=>20&:name=!Mat", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "2d6ad8d46cd8d08d5dfeb91a30dd4cd50f30eb36", signed, "Different order of args shouldn't matter")

}

func TestGetSignedURL(t *testing.T) {

	var signed string

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=6c3dc03b3f85c9eb80ed9e4bd21e82f1bbda5b8d", signed)

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=2d6ad8d46cd8d08d5dfeb91a30dd4cd50f30eb36", signed)

}

func TestValidateSignature(t *testing.T) {

	var valid bool

	signed, _ := GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")

	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&%%7Esign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "Signature with %7E should also work")

	valid, _ = ValidateSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=qJWro1ZxLeToLjNr5Znfi2ZbD+o=", "ABC123", "ABC123-private-wrong")
	assert.Equal(t, false, valid, "2")

	signed, _ = GetSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "3")

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("get", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "4")

	valid, _ = ValidateSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&", "ABC123", "ABC123-private")
	assert.Equal(t, false, valid, "5")

}

func TestNoBodyHashWhenNoBody(t *testing.T) {

	signed, _ := GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=a895aba356712de4e82e336be599df8b665b0fea", signed)

}

func TestSigning_BodyInURL(t *testing.T) {

	valid, _ := ValidateSignature("GET", `http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~body={"question":"Is this OK & working?"}&~sign=934c2494dd617dfeeae63a3a3341f6f4db0adadb`, "", "ABC123-Private")
	assert.Equal(t, true, valid, "1")

	// The tests below represent real requests via JSONP
	valid, _ = ValidateSignature("GET", `http://test.stretchr.com/api/v1/test?~always200=1&~body=%7B%22question%22%3A%22Is%20this%20OK%20%26%20working%3F%22%7D&~callback=Stretchr.callback&~context=1&~key=PjPQMRsam7ewtQbboRLiEC7n88kICT5d&~method=POST&~sign=7a8be95b929b4496a0a68f2a00a01f38cadef872`, "", "HHyLNu5sSt3tYdrUNVukG57tidfo89W1")
	assert.Equal(t, true, valid, "2")

}
