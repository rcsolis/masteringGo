package clib
// #cgo CFLAGS: -I${SRCDIR}/myclib
// #cgo LDFLAGS: ${SRCDIR}/myclib/ccode.a
//#include <stdlib.h>
//#include <ccode.h>
import "C"
import (
	"fmt"
	"unsafe"
)
//CallingCcodeSeparetely For calling a C code in separete files, we must compile it before like gcc -c Ccode.c
//and then create a library like ar rs Ccode.a *.o . At the end, in Go we need to include it with FLAGS and then
//we can use the .h header file and call the code.
func CallingCcodeSeparetely() {
	fmt.Println("--- Calling C code using separete files from C and GO")
	fmt.Println("You need to compile c file with gcc -c and then create and archive file (.a) with ar rs [FILE].a *.o")
	fmt.Println("Calling C function from Go:")
	C.helloFromC()
	fmt.Println("Calling C function from Go with arguments")
	// Define the argument type
	myGoMessage := C.CString(" Hello from GO!!")
	// FREE MEMORY!!! Its important to free memory after ends the function
	defer C.free(unsafe.Pointer(myGoMessage))
	C.passingArgumentsToC(myGoMessage)
	fmt.Println("All perfectly done!")
}
