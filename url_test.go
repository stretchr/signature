package signature

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestOrderParams(t *testing.T) {

	values := make(url.Values)
	values.Add("~key", "ABC123")
	values.Add(":name", "!Mat")
	values.Add(":name", "!Laurie")
	values.Add(":age", ">20")
	values.Add(":something", ">2 0")

	ordered := OrderParams(values)

	//TODO: @matryer does this make sense? We no longer encode right?
	//assert.Equal(t, "%3Aage=%3E20&%3Aname=%21Laurie&%3Aname=%21Mat&%3Asomething=%3E2+0&~key=ABC123", ordered)
	assert.Equal(t, ":age=>20&:name=!Laurie&:name=!Mat&:something=>2 0&~key=ABC123", ordered)

}
