package codec

import (
	"io"
)

type Header struct {
	ServiceMethod string // format 'Service.Method'
	Seq           uint64 // sequence number chosen by client
	Error         string // Err Msg
}

// Codec Codec is an interface used to encode and
// decode the message body
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodeFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // has not yet implemented
)

var NewCodeFuncMap map[Type]NewCodeFunc

func init() {
	NewCodeFuncMap = make(map[Type]NewCodeFunc)
	NewCodeFuncMap[GobType] = NewGobCodec
}
