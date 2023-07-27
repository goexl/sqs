package transcoder

type Decoder interface {
	Decode(from *string, to any) (err error)
}
