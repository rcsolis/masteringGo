package internals

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func printStatistics( memory runtime.MemStats){
	runtime.ReadMemStats(&memory)
	fmt.Println("----------")
	fmt.Println("Mem Allocation:", memory.Alloc)
	fmt.Println("Mem Total Allocation:", memory.TotalAlloc)
	fmt.Println("Mem Heap Allocation:", memory.HeapAlloc)
	fmt.Println( "Mem Num.GCs:", memory.NumGC )
	fmt.Println("----------")
}
//unsafeCode The unsafe code (package unsafe) bypass the type safety and the memory secutiry of Go
// mostly related to pointers, but It can be dangerous, if you dont be sure tu needed, don't use it
func unsafeCode(){
	var value int64 = 10
	var pointer1 = &value
	// Create a int32 pointer to the int64 variable that the pointer 1 points to
	var pointer2 = (*int32)(unsafe.Pointer(pointer1))

	fmt.Printf("Original Variable: %v %T  \n", value, value)
	fmt.Printf("Pointer 1: %v %T  Value=%v \n", pointer1, pointer1, *pointer1)
	fmt.Printf("Pointer 2: %v %T  Value=%v \n", pointer2, pointer2, *pointer2)
	*pointer1 = 5434123412312431212
	fmt.Printf("Change 1 Variable: %v %T  \n", value, value)
	fmt.Printf("Pointer 1: %v %T  Value=%v \n", pointer1, pointer1, *pointer1)
	fmt.Printf("Pointer 2: %v %T  Value=%v \n", pointer2, pointer2, *pointer2)
	*pointer1 = 10
	fmt.Printf("Changes 2 Variable: %v %T  \n", value, value)
	fmt.Printf("Pointer 1: %v %T  Value=%v \n", pointer1, pointer1, *pointer1)
	fmt.Printf("Pointer 2: %v %T  Value=%v \n", pointer2, pointer2, *pointer2)
	fmt.Println()
	arr := [...]int{9, 1, -2, 3, 4}
	var arr_pointer *int = &arr[0]
	fmt.Printf("Arr Pointer %v %T value=%v \n",arr_pointer, arr_pointer, *arr_pointer )
	memoryAddress := uintptr(unsafe.Pointer(arr_pointer)) + unsafe.Sizeof(arr[0])
	fmt.Printf("mem addr %v %T \n\n",memoryAddress, memoryAddress)
	for i:=0; i<len(arr); i++ {
		arr_pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Printf("Arr Pointer %v %T value=%v \n",arr_pointer, arr_pointer, *arr_pointer )
		memoryAddress = uintptr(unsafe.Pointer(arr_pointer)) + unsafe.Sizeof(arr[0])
		fmt.Printf("mem addr %v %T \n",memoryAddress, memoryAddress)
	}
	fmt.Println("One more")
	memoryAddress = uintptr(unsafe.Pointer(arr_pointer)) + unsafe.Sizeof(arr[0])
	arr_pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Printf("Arr Pointer %v %T value=%v \n",arr_pointer, arr_pointer, *arr_pointer )
	fmt.Printf("mem addr %v %T \n",memoryAddress, memoryAddress)
	fmt.Println()

}

func GoInternals(){
	fmt.Println("--- Garbage collections statistics")
	var mem runtime.MemStats
	printStatistics(mem)
	fmt.Println("Starting a simulation of the memory allocation...")
	// Simulate memory allocation for trigger the garbage collector
	for i:= 0; i<10; i++ {
		s:= make( []byte, 50000000)
		if s == nil {
			fmt.Println("Allocation: Operation failed!!")
		}
	}
	fmt.Println("Ends simulation of the memory allocation...")
	printStatistics(mem)
	// Simulate more memory allocation
	fmt.Println("Starting a second simulation of the memory allocation...")
	for i:= 0; i< 10; i++ {
		s := make( []byte, 10000000)
		if s == nil {
			fmt.Println("Second Allocation: Operation failed!!")
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Ends second simulation of the memory allocation...")
	printStatistics(mem)

	fmt.Println("--- Using unsage code (unsafe package)")
	unsafeCode()
}