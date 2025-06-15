package axongo

type Payloadable interface {
	// GetType indicates what payload/response type applies.
	GetType() string
}
