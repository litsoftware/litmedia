package errors

// NotImplemented If a feature is not implemented
type NotImplemented struct {
	Message string
}

func (e NotImplemented) Error() string {
	return e.Message
}

// UnsupportedMetadata - unsupported metadata
type UnsupportedMetadata struct{}

func (e UnsupportedMetadata) Error() string {
	return "Unsupported headers in Metadata"
}
