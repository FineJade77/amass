package utils

import (
	"reflect"
	"unsafe"
)

/*
At the bottom of Golang, string and slice are actually structs

type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}

type StringHeader struct {
    Data uintptr
    Len  int
}
*/

func StringToBytes(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytesHeader := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bytesHeader))
}

func BytesToString(b []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{
		Data: bytesHeader.Data,
		Len:  bytesHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&strHeader))
}

func StringToBytesSimple(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func BytesToStringSimple(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
