package chunk

import (
	"fmt"
	"hash/crc32"
	"pgnme/pkg/utils"
)

const crcPoly = 0xedb88320

var (
	crcTable *crc32.Table
)

type chunk struct {
	typeCode chunkType
	data     []byte
	crc      uint32
}

func New(tp chunkType, data []byte) (*chunk, error) {
	if !tp.IsValid() {
		return nil, fmt.Errorf("ivalid chunk type: %s", tp.String())
	}
	if crcTable == nil {
		crcTable = crc32.MakeTable(crcPoly)
	}
	crc := crc32.Checksum(append(data, tp.Bytes()...), crcTable)
	return &chunk{
		typeCode: tp,
		data:     data,
		crc:      crc,
	}, nil
}

func (c *chunk) Len() int {
	return len(c.data)
}

func (c *chunk) Data() []byte {
	return c.data
}

func (c *chunk) String() string {
	return utils.B2S(c.data)
}

func (c *chunk) Type() chunkType {
	return c.typeCode
}
