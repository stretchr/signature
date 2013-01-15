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

	assert.Equal(t, "%3Aage=%3E20&%3Aname=%21Laurie&%3Aname=%21Mat&%3Asomething=%3E2+0&~key=ABC123", ordered)

}
