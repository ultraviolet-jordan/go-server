package cache

import "awesomeProject/io"

type ConfigDecoder interface {
	decode(*io.Packet, int32)
}

func DecodeType(buf *io.Packet, decoder ConfigDecoder) {
	for {
		code := buf.G1()
		if code == 0 {
			break
		}
		decoder.decode(buf, code)
	}
}
