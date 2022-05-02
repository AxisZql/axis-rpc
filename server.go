package axisrpc

import "axisrpc/codec"

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int        // MagicNumber mark this's a axisrpc request
	CodecType   codec.Type // client may choose different Codec to encode body (不同类型的编码格式选项)
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType, // 默认采取gob格式
}
