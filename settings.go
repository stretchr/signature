package signature

var (
	// SignPrivateKey is the key (URL field) for the private key.
	SignPrivateKey string = "~private"

	// SignPublicKey is the key (URL field) for the public key.
	SignPublicKey string = "~key"

	// BodyHashKey is the key (URL field) for the body hash used for signing requests.
	BodyHashKey string = "~bodyhash"

	// SignatureKey is the key (URL field) for the signature of requests.
	SignatureKey string = "~sign"
)
