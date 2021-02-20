package wrapper

//#include "../bsdiff/bsdiff.h"
//extern int bsdiff(const uint8_t* old, int64_t oldsize, const uint8_t* new, int64_t newsize, struct bsdiff_stream* stream);
import "C"

import (
	"io"
)

func Diff(oldReader, newReader io.Reader, patchWriter io.Writer) (err error) {
	//fmt.Printf("%p\n", unsafe.Pointer(C.bsdiff))
	return nil
}
