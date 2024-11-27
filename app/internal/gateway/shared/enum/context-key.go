package enum

type ContextKey int

const (
	ContextKeyRequestId ContextKey = iota
	ContextKeyClaims
	ContextKeyTelegram
)
