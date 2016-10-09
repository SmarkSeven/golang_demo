package main_test

import (
	"reflect"
	"testing"
	"unsafe"
)

type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(""))

// func MyStructToBytes(s *MyStruct) []byte {
// 	var x reflect.SliceHeader
// 	x.Len = sizeOfMyStruct
// 	x.Cap = sizeOfMyStruct
// 	x.Data = uintptr(unsafe.Pointer(s))
// 	return *(*[]byte)(unsafe.Pointer(&x))
// }
// func BytesToMyStruct(b []byte) *MyStruct {
// 	return (*MyStruct)(unsafe.Pointer(
// 		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
// 	))
// }

func MyStructToBytes(s string) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&x))
}
func BytesToMyStruct(b []byte) string {
	return *(*string)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}

func Benchmark_MyStructToBytes(b *testing.B) {
	var x string = "&MyStruct{}"
	for i := 0; i < b.N; i++ {
		_ = MyStructToBytes(x)
	}
}

func Benchmark_BytesToMyStruct(b *testing.B) {
	var x = []byte("hello world")
	for i := 0; i < b.N; i++ {
		_ = BytesToMyStruct(x)
	}
}

func Benchmark_StringToByte(b *testing.B) {
	var x = "hello world"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}

func TestStringToByte(t *testing.T) {
	var x = "hello world"
	bs := MyStructToBytes(x)
	s := BytesToMyStruct(bs)
	t.Log(s)
	if s != x {
		t.Fail()
	}

}

func byteString(b []byte) string {
	// reflect.StringHeader和reflect.SliceHeader的结构体相似
	return *(*string)(unsafe.Pointer(&b))
}

func Benchmark_ByteToString(b *testing.B) {
	var x = []byte("good")
	for i := 0; i < b.N; i++ {
		_ = byteString(x)
	}
}
