package judy

/*
#cgo LDFLAGS: -lJudy
#include <Judy.h>
*/
import "C"

import (
	"unsafe"
)

type JudyHS struct {
	array unsafe.Pointer
}

// Insert an Index and Value into the JudyHS array. If the Index is successfully inserted, the Value is
// initialized as well. If the Index was already present, the current Value is replaced with the provided Value.
func (j *JudyHS) Insert(data []byte, value uint64) {
	len := len(data)
	pval := unsafe.Pointer(C.JudyHSIns(C.PPvoid_t(&j.array), unsafe.Pointer(&data[0]), C.Word_t(len), nil))
	*((C.PWord_t)(pval)) = C.Word_t(value)
}

// Delete the Index/Value pair from the JudyHS array.
// Returns true if successful. Returns false if Index was not present.
func (j *JudyHS) Delete(data []byte) bool {
	len := len(data)
	return C.JudyHSDel(C.PPvoid_t(&j.array), unsafe.Pointer(&data[0]), C.Word_t(len), nil) != 0
}

// Get the Value associated with Index in the Judy array
//   returns (value, true) if the index was found
//   returns (_, false) if the index was not found
func (j *JudyHS) Get(data []byte) (uint64, bool) {
	len := len(data)
	pval := unsafe.Pointer(C.JudyHSGet(C.Pcvoid_t(j.array), unsafe.Pointer(&data[0]), C.Word_t(len)))
	if pval == nil {
		return 0, false
	} else {
		return uint64(*((*C.Word_t)(pval))), true
	}
}

// Free the entire JudyL array.
// Return the number of bytes freed.
//
// NOTE: The Judy array allocates memory directly from the operating system and is NOT garbage collected by the
// Go runtime. It is very important that you call Free() on a Judy array after using it to prevent memory leaks.
func (j *JudyHS) Free() uint64 {
	return uint64(C.JudyLFreeArray(C.PPvoid_t(&j.array), nil))
}
