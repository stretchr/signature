package signature

import (
	"errors"
	"fmt"
	stewstrings "github.com/stretchrcom/stew/strings"
	"github.com/stretchrcom/tracer"
	"net/url"
	"strings"
)

// FailedSignature is the string that will be used if signing fails.
const FailedSignature string = ":-("

// ErrNoSignatureFound is the error that is thrown when no signature could be found.
var ErrNoSignatureFound = errors.New("No signature was found.")

// trace writes some trace (if there is a Tracer set).
func trace(format string, args ...interface{}) {
	if Tracer != nil {

		// add the 'signature' prefix to trace
		if len(format) > 0 {
			format = stewstrings.MergeStrings("signature: ", format)
		}

		// trace this
		Tracer.Trace(tracer.LevelDebug, format, args...)
	}
}

// GetSignature gets the signature of a request based on the given parameters.
func GetSignature(method, requestUrl, body, privateKey string) (string, error) {

	trace("GetSignature: method=%s", method)
	trace("GetSignature: requestUrl=%s", requestUrl)
	trace("GetSignature: body=%s", body)
	trace("GetSignature: privateKey=%s", privateKey)

	// parse the URL
	u, parseErr := url.ParseRequestURI(requestUrl)

	if parseErr != nil {
		trace("GetSignature: FAILED to parse the URL: %s", parseErr)
		return FailedSignature, parseErr
	}

	trace("GetSignature: Parsed the URL as: %s", u.String())

	// get the query values
	values := u.Query()

	// add the private key parameter
	values.Set(PrivateKeyKey, privateKey)

	trace("GetSignature: Set the private key (%s): %s", PrivateKeyKey, privateKey)

	if len(body) > 0 {
		bodyHash := Hash(body)
		trace("GetSignature: Set the body hash (%s): %s", BodyHashKey, bodyHash)
		values.Set(BodyHashKey, bodyHash)
	} else {
		trace("GetSignature: Skipping body hash as there's no body (%s).", BodyHashKey)
	}

	// get the ordered params
	orderedParams := OrderParams(values)

	trace("GetSignature: Ordered parameters: %s", orderedParams)

	base := strings.Split(u.String(), "?")[0]
	combined := stewstrings.MergeStrings(strings.ToUpper(method), "&", base, "?", orderedParams)

	trace("GetSignature: Base    : %s", base)
	trace("GetSignature: Combined: %s", combined)

	theHash := Hash(combined)

	trace("GetSignature: Output: %s", theHash)

	return theHash, nil

}

// GetSignedURL gets the URL with the sign parameter added based on the given parameters.
func GetSignedURL(method, requestUrl, body, privateKey string) (string, error) {

	trace("GetSignedURL: method=%s", method)
	trace("GetSignedURL: requestUrl=%s", requestUrl)
	trace("GetSignedURL: body=%s", body)
	trace("GetSignedURL: privateKey=%s", privateKey)

	hash, hashErr := GetSignature(method, requestUrl, body, privateKey)

	if hashErr != nil {
		trace("GetSignedURL: FAILED to get the signature: %s", hashErr)
		return FailedSignature, hashErr
	}

	var signedUrl string
	if strings.Contains(requestUrl, "?") {
		signedUrl = stewstrings.MergeStrings(requestUrl, "&", url.QueryEscape(SignatureKey), "=", url.QueryEscape(hash))
	} else {
		signedUrl = stewstrings.MergeStrings(requestUrl, "?", url.QueryEscape(SignatureKey), "=", url.QueryEscape(hash))
	}

	trace("GetSignedURL: Output: %s", signedUrl)

	return signedUrl, nil

}

// ValidateSignature validates the signature in a URL to ensure it is correct based on
// the specified parameters.
func ValidateSignature(method, requestUrl, body, privateKey string) (bool, error) {

	trace("ValidateSignature: method=%s", method)
	trace("ValidateSignature: requestUrl=%s", requestUrl)
	trace("ValidateSignature: body=%s", body)
	trace("ValidateSignature: privateKey=%s", privateKey)

	if !strings.Contains(requestUrl, "?") {
		trace("ValidateSignature: FAILED because there was no signature found.")
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

	trace("ValidateSignature: Modified URL (without signature): %s", modifiedURL)

	expectedSignature, signErr := GetSignature(method, modifiedURL, body, privateKey)

	if signErr != nil {
		trace("ValidateSignature: FAILED to GetSignature: %s", signErr)
		return false, signErr
	}

	if signature != expectedSignature {
		err := errors.New(fmt.Sprintf("Signature \"%s\" is incorrect when \"%s\" is expected.", signature, expectedSignature))
		trace("ValidateSignature: Signatures do not match: %s", err)
		return false, err
	}

	trace("ValidateSignature: Happy because the signatures match: %s", signature)

	return true, nil

}
