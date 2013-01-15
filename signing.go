package signature

import (
	"errors"
	"fmt"
	stewstrings "github.com/stretchrcom/stew/strings"
	"net/url"
	"strings"
)

// FailedSignature is the string that will be used if signing fails.
const FailedSignature string = ":-("

var ErrNoSignatureFound = errors.New("No signature was found.")

// GetSignature gets the signature of a request based on the given parameters.
func GetSignature(method, requestUrl, body, privateKey string) (string, error) {

	// parse the URL
	u, parseErr := url.ParseRequestURI(requestUrl)

	if parseErr != nil {
		return FailedSignature, parseErr
	}

	// get the query values
	values := u.Query()

	// add the private key parameter
	values.Set(PrivateKeyKey, privateKey)

	if len(body) > 0 {
		values.Set(BodyHashKey, Hash(body))
	}

	// get the ordered params
	orderedParams := OrderParams(values)

	base := strings.Split(u.String(), "?")[0]
	combined := stewstrings.MergeStrings(strings.ToUpper(method), "&", base, "?", orderedParams)

	return Hash(combined), nil

}

// GetSignedURL gets the URL with the sign parameter added based on the given parameters.
func GetSignedURL(method, requestUrl, body, privateKey string) (string, error) {

	hash, hashErr := GetSignature(method, requestUrl, body, privateKey)

	if hashErr != nil {
		return FailedSignature, hashErr
	}

	signed := stewstrings.MergeStrings(requestUrl, "&", url.QueryEscape(SignatureKey), "=", url.QueryEscape(hash))

	return signed, nil

}

// ValidateSignature validates the signature in a URL to ensure it is correct based on
// the specified parameters.
func ValidateSignature(method, requestUrl, body, privateKey string) (bool, error) {

	// trim off the sign parameter
	segs := strings.Split(requestUrl, stewstrings.MergeStrings("&", SignatureKey, "="))

	if len(segs) < 2 {
		return false, ErrNoSignatureFound
	}

	modifiedUrl := segs[0]
	signature := segs[1]

	expectedSignature, signErr := GetSignature(method, modifiedUrl, body, privateKey)

	if signErr != nil {
		return false, signErr
	}

	if signature != expectedSignature {
		return false, errors.New(fmt.Sprintf("Signature \"%s\" is incorrect when \"%s\" is expected.", signature, expectedSignature))
	}
	return true, nil

}
