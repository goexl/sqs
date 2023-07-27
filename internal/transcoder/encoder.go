package transcoder

type Encoder interface {
	Encode(from any) (to *string, err error)
}
