package wrapper

//#include "../bsdiff/bsdiff.c"
// #include <stdio.h>
// #include <stdlib.h>
// extern int cgo_write(struct bsdiff_stream* stream, const void* buffer, int size);
// int cgo_bsdiff(const void* old, int64_t oldsize, const void* new, int64_t newsize, void *opaque) {
//     printf("%p %lld %p %lld %p\n", old, oldsize, new, newsize, opaque);
//     struct bsdiff_stream stream;
//     stream.opaque = opaque;
//     stream.malloc = malloc;
//     stream.free = free;
//     stream.write = cgo_write;
// 	   int ret = bsdiff(old, oldsize, new, newsize, &stream);
//	   return ret;
// }
//
import "C"

import (
	"bytes"
	"io"
	"io/ioutil"
	"unsafe"

	"go.uber.org/zap"
)

func Diff(oldReader, newReader io.Reader, patch *bytes.Buffer) (err error) {
	oldB, err := ioutil.ReadAll(oldReader)
	if err != nil {
		return err
	}

	newB, err := ioutil.ReadAll(newReader)
	if err != nil {
		return err
	}

	zap.L().Debug("file compare",
		zap.Int("result", bytes.Compare(oldB, newB)))
	oldP := C.CBytes(oldB)
	newP := C.CBytes(newB)
	ret := C.cgo_bsdiff(oldP, C.longlong(len(oldB)),
		newP, C.longlong(len(newB)),
		unsafe.Pointer(patch))
	C.free(oldP)
	C.free(newP)
	if ret == 0 {
		return nil
	}

	return nil
}
