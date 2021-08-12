package panicRecover
//Panic ends the current flow and starts panicking
//Recover allows take back the control of a GORUTINE that just panicking using panic
import "fmt"

func myA(){
	fmt.Println("Inside myA() function")
	// Anonymous function
	defer func() {
		if c:= recover(); c!= nil {
			fmt.Println("Recover inside myA() function!")
		}
	}()
	fmt.Println("Calling myB() function")
	myB()
	fmt.Println("myB() functions ends")
	fmt.Println("myA() function ends")
}

func myB(){
	fmt.Println("Inside myB() function")
	panic("Panic in myB() function!!!")
	fmt.Println("Exiting myB()")
}


func PanicAndRecover(){
	fmt.Println("--- Starting Panic and Recover sample")
	fmt.Println("Calling myA function")
	myA()
	fmt.Println("Ends Panic and Recover")
}