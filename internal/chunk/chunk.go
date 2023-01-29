package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	typecode "pngme/internal/type_code"
	"pngme/pkg/utils"
)

const crcPoly = 0xedb88320
const IEND = "IEND"

var (
	crcTable *crc32.Table
)

type Chunk struct {
	crc  uint32
	tp   typecode.ChunkType
	data []byte
}

func New(chTp []byte, data []byte) (*Chunk, error) {
	tp, err := typecode.FromBytes(chTp)
	if err != nil {
		return nil, err
	}

	if !tp.IsValid() {
		return nil, fmt.Errorf("ivalid chunk type: %s", tp.String())
	}

	if crcTable == nil {
		crcTable = crc32.MakeTable(crcPoly)
	}

	bytesMBS := new(bytes.Buffer)

	if err = binary.Write(bytesMBS, binary.BigEndian, tp); err != nil {
		return nil, err
	}

	if err = binary.Write(bytesMBS, binary.BigEndian, data); err != nil {
		return nil, err
	}

	crc := crc32.Checksum(bytesMBS.Bytes(), crcTable)

	return &Chunk{
		tp:   *tp,
		data: data,
		crc:  crc,
	}, nil
}

func MakeIEND() *Chunk {
	chnk, err := New(utils.S2B(IEND), []byte{})
	if err != nil {
		panic("cannot create IEND!")
	}
	return chnk
}

func (c *Chunk) Size() int {
	return len(c.data)
}

func (c *Chunk) Data() []byte {
	return c.data
}

func (c *Chunk) String() string {
	return utils.B2S(c.data)
}

func (c *Chunk) Type() *typecode.ChunkType {
	return &c.tp
}

func (c *Chunk) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Grow(c.Size() + 10)

	if err := binary.Write(buf, binary.BigEndian, uint32(c.Size())); err != nil {
		return nil, fmt.Errorf("marshal chunk: %v\n", err)
	}

	if _, err := buf.Write(c.tp.Marshal()); err != nil {
		return nil, fmt.Errorf("marshal chunk: %v\n", err)
	}

	if _, err := buf.Write(c.data); err != nil {
		return nil, fmt.Errorf("marshal chunk: %v\n", err)
	}

	if err := binary.Write(buf, binary.BigEndian, c.crc); err != nil {
		return nil, fmt.Errorf("marshal chunk: %v\n", err)
	}

	return buf.Bytes(), nil
}

func (c *Chunk) Print(i int) {
	var data string
	if c.Size() >= 30 {
		data = utils.B2S(c.Data()[:30])
	} else {
		data = utils.B2S(c.Data())
	}
	fmt.Printf("---------------------------\n")
	fmt.Printf("Chunk: # %d\n", i)
	fmt.Printf("---------------------------\n")
	fmt.Printf("Chunk length: # %d\n", c.Size())
	fmt.Printf("---------------------------\n")
	fmt.Printf("Chunk type: %s\n", c.Type().String())
	fmt.Printf("---------------------------\n")
	fmt.Printf("Chunk data: (%d bytes) %v\n", len(data), data)
	fmt.Printf("---------------------------\n")
	fmt.Printf("\n")
}
