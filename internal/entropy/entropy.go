package entropy

/*
cgo LDFLAGS: -lm
#include "entropy.h"
*/
import "C"
import "unsafe"

// GetEntropy invoca el cálculo en C pasando un puntero al buffer de Go.
func GetEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0.0
	}

	// Obtenemos el puntero a la base del slice.
	// C.unsignedchar equivale a uint8 en Go.
	ptr := (*C.uchar)(unsafe.Pointer(&data[0]))
	size := C.int(len(data))

	return float64(C.calculate_entropy(ptr, size))
}
