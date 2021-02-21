package wrapper

// #include "../bsdiff/bsdiff.c"
// int cgo_bsdiff(const uint8_t* old, int64_t oldsize, const uint8_t* new, int64_t newsize, const void *opaque) {
//     struct bsdiff_stream stream;
//     stream.opaque = opaque;
//     stream.malloc = malloc;
//     stream.free = free;
//     stream.write = cgo_write;
// 	   return bsdiff(old, oldsize, new, newsize, &stream);
// }
//
import "C"

import (
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

func Diff(oldReader, newReader io.Reader, patchWriter io.Writer) (err error) {
	oldB, err := ioutil.ReadAll(oldReader)
	if err != nil {
		return err
	}

	newB, err := ioutil.ReadAll(newReader)
	if err != nil {
		return err
	}

	oldP := C.CBytes(oldB)
	newP := C.CBytes(newB)
	ret := C.cgo_bsdiff(oldP, len(oldB), newB, len(newP), unsafe.Pointer(patchWriter))
	fmt.Println(C.bsdiff)
	fmt.Printf("%p\n", C.bsdiff)
	stream := new(C.struct_bsdiff_stream)
	//stream.malloc = C.malloc
	//stream.free = C.free
	//stream.write = write
	fmt.Printf("%+v\n", stream)

	return nil
}
