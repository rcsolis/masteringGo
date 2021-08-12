package errorHandling;

// definition of custom errors
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// customErrorHandling function to create a custom error, uses named return values and naked return
func customErrorHandling( a, b int) (res int, err error){
	if a == b{
		err = errors.New("Error in custom error handling function")
		return
	}
	res = a * b
	return
}

func getError(msg string) error{
	return errors.New(msg)
}

func sumArgs(args []string) float64{
	var total float64 = 0

	for i:= 0; i< len(args); i++{
		val, err := strconv.ParseFloat(args[i], 64)
		if err == nil{
			total += val
		}
	}

	return total
}

func averageArgs(args []string) (average float64, total_nums int){
	total_nums = 0
	average = 0
	for i:=0; i < len(args); i++ {
		val, err := strconv.ParseFloat(args[i], 64)
		if err == nil {
			average += val
			total_nums += 1
		}
	}
	if total_nums > 0{
		average = average/float64(total_nums)
	}
	return
}

func captureInts() {
	// Reference to a standard input
	var file *os.File = os.Stdin
	// Close after execution ends
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "STOP" {
			break
		}else if val, err:= strconv.ParseInt(input, 10, 64); err == nil {
			fmt.Println("Data >", val)
		}else{
			fmt.Println("Not an integer")
		}
	}
}

//ErrorHandlingMain The main function to execute all the exercises related to the errors handling
func ErrorHandlingMain(args []string){
	// custom error handling
	testA, testB := 10, 10
	res, err := customErrorHandling(testA,testB)
	fmt.Printf("Response of custom error handling: a=%v b=%d res=%v(%T) err=%v(%T) \n", testA, testB, res,res, err, err )
	// Reuse variables, do not use :=
	testA, testB = 2, 4
	res, err = customErrorHandling(testA,testB)
	fmt.Printf("Response of custom error handling: a=%v b=%d res=%v(%T) err=%v(%T) \n", testA, testB, res,res, err, err )
	// Parse error as string
	fmt.Println("error.Error() parse error type to a string")
	var err2 = getError("This is a custom error")
	if err2.Error() == "This is a custom error" {
		fmt.Println("To compare a error variable to a string")
	}
	fmt.Println("Exercise1: Finds de sum of all of its numeric command line args")
	total := sumArgs(args)
	fmt.Println("ARGS:", args, "Result:", total)
	fmt.Println("Exercise2: Finds de average value of all of its numeric command line args")
	average, total_numbers := averageArgs(args)
	fmt.Println("ARGS:", args, "Average:", average, "Total numbers:", total_numbers)
	fmt.Println("Exercise3: Capture ints until STOP")
	captureInts()
}

