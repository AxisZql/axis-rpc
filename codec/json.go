package codec

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
)

/*
@author: axiszql
@date: 2022-5-2
@desc: the json decode and encode way implement
*/

type JsonCodec struct {
	conn io.ReadWriteCloser // the Uinx or Tcp socket example
	buf  *bufio.Writer
	dec  *json.Decoder
	enc  *json.Encoder
}

var _ Codec = (*JsonCodec)(nil)

func NewJsonCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &JsonCodec{
		conn: conn,
		buf:  buf,
		dec:  json.NewDecoder(conn),
		enc:  json.NewEncoder(buf),
	}
}

//=========implement the method of Codec interface

func (c *JsonCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *JsonCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *JsonCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.conn.Close()
		}
	}()

	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec:json err encode", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec:json err encode", err)
		return err
	}
	return nil
}

func (c *JsonCodec) Close() error {
	return c.conn.Close()
}
