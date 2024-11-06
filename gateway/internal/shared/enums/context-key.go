package enums

type ContextKey int

const (
	ContextKeyRequestId ContextKey = iota
	ContextKeyClaims
	ContextKeyTelegram
)
