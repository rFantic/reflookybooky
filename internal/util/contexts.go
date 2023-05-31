package util

type ContextKey struct{}

var Secretkey []byte = []byte("werq2304u1rjweiofsd")

var (
	RequestIDCK ContextKey = struct{}{}
)
