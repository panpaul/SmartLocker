package face

//#cgo LDFLAGS: -L${SRCDIR}/service/face -LD:/source/go/SmartLocker/backend/service/face -lFaceRecognizer
//
//#include "lib.h"
import "C"
import (
	"fmt"
	"syscall"
	"unsafe"
)

func Setup() {
	C.init()
}

func InsertImage(name string, filepath string) bool {
	fp := C.CString(filepath)
	n := C.CString(name)
	s := C.addImage(n, fp)
	return s == 1
}

func Recognize(filepath string) string {
	fp := C.CString(filepath)
	n := C.recognizer(fp)
	name := C.GoString(n)
	fmt.Println(name)

	return name
}

func strPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}
