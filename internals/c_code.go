package internals
// //C code must be followed by the import of C (with NO SPACES OR NEWLINES between both)
//#include<stdio.h>
//void callingCcode(){
// printf("This is C code called from Go code! \n");
//}
import "C"
// Import Go Packages
import "fmt"

func CallingCCodeSameFile(){
	fmt.Println("-- Calling C from the same go file")
	fmt.Println("Calling C code, whe use C package, all other packages must be import separately")
	fmt.Println("Between import of C package and C code, there not be any newline, last line of C code must be followed by import C ")
	fmt.Println("A simple GO statement")
	C.callingCcode()
	fmt.Println("Another GO statement after calling C code")
}
