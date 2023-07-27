package transcoder

var (
	_ Encoder = (*Json)(nil)
	_ Decoder = (*Json)(nil)
)

type Json struct{}

func NewJson() *Json {
	return new(Json)
}

func (j *Json) Encode(from any) (to *string, err error) {
	return
}

func (j *Json) Decode(from *string, to any) (err error) {
	return
}
