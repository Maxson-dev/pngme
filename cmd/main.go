package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

func main() {
	str := "hello"
	bb := new(bytes.Buffer)
	binary.Write(bb, binary.BigEndian, []byte(str))

	buf1 := []byte(str)
	buf2 := bb.Bytes()

	fmt.Printf("buf1 == buf2? - %t\n", bytes.Equal(buf1, buf2))

	c1 := crc32.ChecksumIEEE(buf1)
	c2 := crc32.ChecksumIEEE(buf2)

	fmt.Printf("buf1 crc == %d\nbuf2 crc == %d\n", c1, c2)
}
