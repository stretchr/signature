signature
=========

URL signing package for Go.

## What does it do?

Secure web calls by generating a security hash on the client (using a private key shared with the server), to ensure that the request is geniune.  Only a client who knows the private key will be able to generate the same security hash.

## How does it work?

### Encoding

  * Generate the original request URL
  * Add public key value

To generate SignatureKey parameter:

  * Create a copy of the request URL
  * Add `PrivateKeyKey` key parameter
  * Add `BodyHashKey` value containing an SHA-1 hash of the body contents if there is a body - otherwise, skip this step (and do not add a `BodyHashKey` parameter at all)
  * Order parameters alphabetically
  * Prefix it with the HTTP method (in uppercase) followed by an ampersand (i.e. "GET&http://...")
  * Hash it (using SHA-1)
  * Add the hash as `SignatureKey` to the END of the original URL

### Decoding

  * Strip off the `SignatureKey` parameter (and keep it)
  * Lookup the account (using the public key) and get the `PrivateKeyKey` parameter, and add it to the URL
  * Hash it
  * Compare the generated hash with the `SignatureKey` value to decide if it the request is valid or not

## Validation

To validate your code is generating the correct hash, try these:

The URL

    http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20
    
with HTTP method

    GET
    
and body

    body
    
and private key

    ABC123-private
    
should be hashed as

    df073ee4086eed5848d167871c7424937027728e

leaving the final URL as

    http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=df073ee4086eed5848d167871c7424937027728e
