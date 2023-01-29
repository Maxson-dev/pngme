package typecode

import (
	"bytes"
	"testing"
)

func TestFromBytes(t *testing.T) {
	in := []byte{82, 117, 83, 84}
	res, err := FromBytes(in)
	if err != nil {
		t.Fatalf("FromBytes(%v) -> %v", in, err)
	}
	if !bytes.Equal(in, res.Bytes()) {
		t.Fatalf("expected: %v, have: %v", in, res.Bytes())
	}
}

func TestFromString(t *testing.T) {
	in := []byte{82, 117, 83, 84}
	expected, err := FromBytes(in)
	if err != nil {
		t.Fatalf("FromBytes(%v) -> %v\n", in, err)
	}
	have, err := FromString("RuST")
	if err != nil {
		t.Fatalf(`FromString("RuST") -> %v`, err)
	}
	if !bytes.Equal(expected.Bytes(), have.Bytes()) || expected.String() != have.String() {
		t.Fatalf("expected: %s, have: %s", expected.String(), have.String())
	}
}

func TestIsCritical(t *testing.T) {
	str := "TeSt"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if !chunk.IsCritical() {
		t.Fatalf("chunk.IsCritical('%s') -> false", str)
	}
}

func TestIsNotCritical(t *testing.T) {
	str := "teSt"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if chunk.IsCritical() {
		t.Fatalf("chunk.IsCritical('%s') -> true", str)
	}
}

func TestIsPublic(t *testing.T) {
	str := "TESt"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if !chunk.IsPublic() {
		t.Fatalf("chunk.IsPublic('%s') -> false", str)
	}
}

func TestIsNotPublic(t *testing.T) {
	str := "RuSt"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if chunk.IsPublic() {
		t.Fatalf("chunk.IsPublic('%s') -> true", str)
	}
}

func TestIsReservedBitValid(t *testing.T) {
	str := "TeSt"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if !chunk.isReservedBitValid() {
		t.Fatalf("chunk.isReservedBitValid(%s) -> false", str)
	}
}

func TestIsReservedBitInvalid(t *testing.T) {
	str := "Test"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if chunk.isReservedBitValid() {
		t.Fatalf("chunk.isReservedBitValid('%s') -> true", str)
	}
}

func TestIsSafeToCopy(t *testing.T) {
	str := "RuSt"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if !chunk.IsSafeToCopy() {
		t.Fatalf("chunk.IsSafeToCopy('%s') -> false", str)
	}
}

func TestIsUnsafeToCopy(t *testing.T) {
	str := "RuST"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if chunk.IsSafeToCopy() {
		t.Fatalf("chunk.IsSafeToCopy('%s') -> true", str)
	}
}

func TestValidChunkIsValid(t *testing.T) {
	str := "TeST"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if !chunk.IsValid() {
		t.Fatalf("chunk.IsValid('%s') -> false", str)
	}
}

func TestInvalidChunkIsValid(t *testing.T) {
	str := "Rust"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if chunk.IsValid() {
		t.Fatalf("chunk.IsValid('%s') -> true", str)
	}
	str = "Te1t"
	chunk, err = FromString(str)
	if err == nil {
		t.Fatalf("FromString('%s') -> expected error, have nil", str)
	}
}

func TestChunkTypeString(t *testing.T) {
	str := "TeSt"
	chunk, err := FromString(str)
	if err != nil {
		t.Fatalf("FromString(%s) -> %v", str, err)
	}
	if chunk.String() != str {
		t.Fatalf("chunk.String('%s') != %s", str, str)
	}
}
