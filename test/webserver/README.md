# Signature package - web server test

This package is a simple web server that implements the `signature` package and allows you to test clients that are trying to generate the same signature.

## Usage

  * Start the web server

In Terminal:

    cd signature/test/webserver
    go run main.go

  * Make a request to the web server like `http://localhost:8080/some/path?~sign=ABC123&`

  * Use the private key `PRIVATE`
