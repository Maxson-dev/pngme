package utils

import (
	"reflect"
	"unsafe"
)

// B2S zero allocation []byte to string
func B2S(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

// S2B zero allocation string to []byte
func S2B(str string) []byte {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	var buf []byte
	bfHdr := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	bfHdr.Len = strHdr.Len
	bfHdr.Cap = strHdr.Len
	bfHdr.Data = strHdr.Data
	return buf
}
