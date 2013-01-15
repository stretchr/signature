signature
=========

URL signing package for Go.

## Encoding

  * Generate the original request URL
  * Add public key value

## To generate SignatureKey parameter

  * Create a copy of the request URL
  * Add PrivateKeyKey key parameter
  * Add BodyHashKey value containing an SHA1 hash of the body contents if there is a body - otherwise, skip this step
  * Order parameters alphabetically
  * Prefix it with the HTTP method (in uppercase) followed by an ampersand (i.e. "GET&http://...")
  * Hash it (using SHA-1)
  * Add the hash as SignatureKey to the END of the original URL

## Decoding

  * Strip off the SignatureKey parameter (and keep it)
  * Lookup the account (using the public key) and get the ~private parameter, and add it to the URL
  * Hash it
  * Compare the generated hash with the SignatureKey value to decide if it the request is valid or not