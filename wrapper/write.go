package wrapper


// #include "../bsdiff/bsdiff.h"
import "C"
import (
	"io"
	"unsafe"
)

//export cgo_write
func cgo_write(stream *C.struct_bsdiff_stream, buffer unsafe.Pointer, size int) int {
	writer := stream.opaque.(io.Writer)
	if _, err := writer.Write(C.GoBytes(buffer, size)); err != nil {
		return -1
	}
	return 0
}


