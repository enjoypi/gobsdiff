package wrapper

// #include "../bsdiff/bsdiff.h"
import "C"
import (
	"bytes"
	"unsafe"
)

//export cgo_write
func cgo_write(stream *C.struct_bsdiff_stream, buffer unsafe.Pointer, size int) int {
	writer := (*bytes.Buffer)(stream.opaque)
	if _, err := writer.Write(C.GoBytes(buffer, C.int(size))); err != nil {
		return -1
	}
	return 0
}
