package anonymizer

import (
	"github.com/golang/protobuf/proto"
)

// Anonymizer is an interface which have Anonymize()
type Anonymizer interface {
	Anonymize()
}

// Anonymize filter messages by calling Anonymize() if the message
// implementes Anonymizer
func Anonymize(msg proto.Message) {
	if msg == nil {
		return
	}
	if a, ok := msg.(Anonymizer); ok {
		a.Anonymize()
	}
}
