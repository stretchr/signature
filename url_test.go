package signature

import (
	"github.com/stretchrcom/testify/assert"
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

	assert.Equal(t, ":age=>20&:name=!Laurie&:name=!Mat&:something=>2 0&~key=ABC123", ordered)

}
