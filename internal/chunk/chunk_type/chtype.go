package chunk

import (
	"fmt"
	"pgnme/pkg/utils"
	"unicode"
)

type chunkType [4]rune

func (c *chunkType) Len() int {
	return len(c)
}

func (c *chunkType) Bytes() []byte {
	res := make([]byte, 4)
	for i := range c {
		res[i] = byte(c[i])
	}
	return res
}

func (c *chunkType) String() string {
	return string(c[:])
}

func (c *chunkType) IsCritical() bool {
	return unicode.IsUpper(c[0])
}

func (c *chunkType) IsPublic() bool {
	return unicode.IsUpper(c[1])
}

func (c *chunkType) IsValid() bool {
	if !c.isReservedBitValid() {
		return false
	}
	if c.IsCritical() && c.IsSafeToCopy() {
		return false
	}
	return true
}

func (c *chunkType) isReservedBitValid() bool {
	return unicode.IsUpper(c[2])
}

func (c *chunkType) IsSafeToCopy() bool {
	return unicode.IsLower(c[3])
}

func FromBytes(bf []byte) (*chunkType, error) {
	if len(bf) != 4 {
		return nil, fmt.Errorf("invalid chunk type legnth: expected - 4, have - %d", len(bf))
	}
	tp := new(chunkType)
	for i := range bf {
		r := rune(bf[i])
		if !unicode.IsLetter(r) {
			return nil, fmt.Errorf("invalid symbol: %v", r)
		}
		tp[i] = r
	}
	return tp, nil
}

func FromString(str string) (*chunkType, error) {
	return FromBytes(utils.S2B(str))
}
