package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// generate n metric names
func genMetricNames(prefix string, id, n int) []string {
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = fmt.Sprintf("\x05agent\x06%6d\x07metrics\x06%6d", id, i)
	}

	return names
}

// actually write the data in carbon line format
func carbonate(w io.ReadWriteCloser, name string, value int, epoch int64) error {
        buf := new(bytes.Buffer)
        binary.Write(buf, binary.BigEndian, uint8(5))
        binary.Write(buf, binary.BigEndian, uint64(epoch))
        binary.Write(buf, binary.BigEndian, uint16(len(name)))
        w.Write(buf.Bytes())
        w.Write([]byte(name))
        buf = new(bytes.Buffer)
        binary.Write(buf, binary.BigEndian, uint32(8))
        binary.Write(buf, binary.BigEndian, uint8(1))
        binary.Write(buf, binary.BigEndian, uint8(0))
        binary.Write(buf, binary.BigEndian, uint8(0))
        binary.Write(buf, binary.BigEndian, uint8(0))
        binary.Write(buf, binary.BigEndian, uint32(value))
        w.Write(buf.Bytes())
	return nil
}
