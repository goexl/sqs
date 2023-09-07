package transcoder

import (
	"encoding/json"
)

var (
	_ Encoder = (*Json)(nil)
	_ Decoder = (*Json)(nil)
)

type Json struct{}

func NewJson() *Json {
	return new(Json)
}

func (j *Json) Encode(from any) (to *string, err error) {
	if bytes, me := json.Marshal(from); nil != me {
		err = me
	} else {
		encoded := string(bytes)
		to = &encoded
	}

	return
}

func (j *Json) Decode(from *string, to any) error {
	return json.Unmarshal([]byte(*from), to)
}
