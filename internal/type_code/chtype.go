package typecode

import (
	"fmt"
	"pngme/pkg/utils"
	"unicode"
)

type ChunkType [4]rune

func (c *ChunkType) Len() int {
	return len(c)
}

func (c *ChunkType) Bytes() []byte {
	res := make([]byte, 4)
	for i := range c {
		res[i] = byte(c[i])
	}
	return res
}

func (c *ChunkType) Marshal() []byte {
	return c.Bytes()
}

func (c *ChunkType) String() string {
	return string(c[:])
}

func (c *ChunkType) IsCritical() bool {
	return unicode.IsUpper(c[0])
}

func (c *ChunkType) IsPublic() bool {
	return unicode.IsUpper(c[1])
}

func (c *ChunkType) IsValid() bool {
	if !c.isReservedBitValid() {
		return false
	}
	if c.IsCritical() && c.IsSafeToCopy() {
		return false
	}
	return true
}

func (c *ChunkType) isReservedBitValid() bool {
	return unicode.IsUpper(c[2])
}

func (c *ChunkType) IsSafeToCopy() bool {
	return unicode.IsLower(c[3])
}

func FromBytes(bf []byte) (*ChunkType, error) {
	if len(bf) != 4 {
		return nil, fmt.Errorf("invalid chunk type legnth: expected - 4, have - %d", len(bf))
	}
	tp := new(ChunkType)
	for i := range bf {
		r := rune(bf[i])
		if !unicode.IsLetter(r) {
			return nil, fmt.Errorf("invalid symbol: %v", r)
		}
		tp[i] = r
	}
	return tp, nil
}

func FromString(str string) (*ChunkType, error) {
	return FromBytes(utils.S2B(str))
}
