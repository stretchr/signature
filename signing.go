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

// ErrNoSignatureFound is the error that is thrown when no signature could be found.
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
	values.Set(SignPrivateKey, privateKey)

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
func GetSignedURL(method, requestUrl, body, publicKey, privateKey string) (string, error) {

	hash, hashErr := GetSignature(method, requestUrl, body, privateKey)

	if hashErr != nil {
		return FailedSignature, hashErr
	}

	if strings.Contains(requestUrl, "?") {
		return fmt.Sprintf("%s&%s=%s&%s=%s", requestUrl, url.QueryEscape(SignPublicKey), publicKey, url.QueryEscape(SignatureKey), url.QueryEscape(hash)), nil
	}

	return fmt.Sprintf("%s?%s=%s&%s=%s", requestUrl, url.QueryEscape(SignPublicKey), publicKey, url.QueryEscape(SignatureKey), url.QueryEscape(hash)), nil

}

// ValidateSignature validates the signature in a URL to ensure it is correct based on
// the specified parameters.
func ValidateSignature(method, requestUrl, body, privateKey string) (bool, error) {

	if !strings.Contains(requestUrl, "?") {
		return false, ErrNoSignatureFound
	}

	// First, get the query string alone
	segs := strings.Split(requestUrl, "?")

	bareURL := segs[0]

	// segs[1] now contains all the parameters. We need to extract the signature
	// and reconstruct the url without it
	paramSegs := strings.Split(segs[1], "&")

	var cleanParams []string
	var signature string
	for _, param := range paramSegs {
		if strings.Contains(param, SignatureKey) {
			sigParts := strings.Split(param, "=")
			signature = sigParts[1]
		} else {
			cleanParams = append(cleanParams, param)
		}
	}

	modifiedURL := stewstrings.MergeStrings(bareURL, "?", stewstrings.JoinStrings("&", cleanParams...))

	expectedSignature, signErr := GetSignature(method, modifiedURL, body, privateKey)

	if signErr != nil {
		return false, signErr
	}

	if signature != expectedSignature {
		return false, errors.New(fmt.Sprintf("Signature \"%s\" is incorrect when \"%s\" is expected.", signature, expectedSignature))
	}
	return true, nil

}
