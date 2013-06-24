package signature

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSignature(t *testing.T) {

	var signed string

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "0bd60ce074a9a3ecda66a438f04a6cf779ab60d3", signed)

	signed, _ = GetSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "0bd60ce074a9a3ecda66a438f04a6cf779ab60d3", signed, "Lower case method shouldn't affect GetSignature")

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "be7348dd329e0791b3c082a9044e55bc16779587", signed)

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?:name=!Laurie&~key=ABC123&:age=>20&:name=!Mat", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "be7348dd329e0791b3c082a9044e55bc16779587", signed, "Different order of args shouldn't matter")

}

func TestGetSignedURL(t *testing.T) {

	var signed string

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=0bd60ce074a9a3ecda66a438f04a6cf779ab60d3", signed)

	signed, _ = GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=be7348dd329e0791b3c082a9044e55bc16779587", signed)

}

func TestValidateSignature(t *testing.T) {

	var valid bool

	signed, _ := GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")

	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

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

func TestValidateSignature_NoTilde(t *testing.T) {

	var valid bool

	signed, _ := GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")

	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "1")

	valid, _ = ValidateSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=qJWro1ZxLeToLjNr5Znfi2ZbD+o=", "ABC123", "ABC123-private-wrong")
	assert.Equal(t, false, valid, "2")

	signed, _ = GetSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("GET", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "3")

	signed, _ = GetSignature("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "ABC123", "ABC123-private")
	valid, _ = ValidateSignature("get", fmt.Sprintf("http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=%s", signed), "ABC123", "ABC123-private")
	assert.Equal(t, true, valid, "4")

	valid, _ = ValidateSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&", "ABC123", "ABC123-private")
	assert.Equal(t, false, valid, "5")

}

func TestNoBodyHashWhenNoBody(t *testing.T) {

	signed, _ := GetSignedURL("GET", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "", "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&sign=0e1e85ebdf8c4bcebdfc9033b6590c6c4a13a78f", signed)

}

func TestSigning_BodyInURL(t *testing.T) {

	valid, _ := ValidateSignature("GET", `http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~body={"question":"Is this OK & working?"}&sign=a1f72a10e882ac64236c43dca381981c77aa8a48`, "", "ABC123-Private")

	assert.Equal(t, true, valid, "1")

	// The tests below represent real requests via JSONP
	valid, _ = ValidateSignature("GET", `http://test.stretchr.com/api/v1/test?~always200=1&~body=%7B%22question%22%3A%22Is%20this%20OK%20%26%20working%3F%22%7D&~callback=Stretchr.callback&~context=1&~key=PjPQMRsam7ewtQbboRLiEC7n88kICT5d&~method=POST&sign=63c309c72bb3e8626187cbbf344b5a2c50fcf450`, "", "HHyLNu5sSt3tYdrUNVukG57tidfo89W1")

	assert.Equal(t, true, valid, "2")

}
