package main

import (
	"fmt"
	"io"
	"masteringGo/defered"
	goEnv "masteringGo/environmentInfo"
	"masteringGo/errorHandling"
	goInternals "masteringGo/internals"
	goAnExternalC "masteringGo/internals/clib"
	"masteringGo/logs"
	"masteringGo/panicRecover"
	"os"
)

// printStandardIO using standard output to print a message
func printStandardIO(message string) {
	io.WriteString(os.Stdout, message)
}

func main() {
	fmt.Println("Welcome to mastering go")
	var command_args []string

	if len(os.Args) < 2 {
		panic("PANIC: No arguments")
	} else {
		switch os.Args[1] {
		case "logs":
			logs.UsingSystemLogServer("Send this message to the Log!!!")
			logs.CustomLog("This is a custom error")
		case "errors":
			command_args = os.Args[2:]
			printStandardIO("** Error Handling Samples **\n")
			errorHandling.ErrorHandlingMain(command_args)
		case "memstats":
			printStandardIO("** Garbage Collector and Unsafe package **\n")
			goInternals.GoInternals()
		case "callingC":
			printStandardIO("** Calling C code **\n")
			goInternals.CallingCCodeSameFile()
			goAnExternalC.CallingCcodeSeparetely()
		case "panicRecover":
			printStandardIO("** Using panic() and recover() **\n")
			panicRecover.PanicAndRecover()
		case "defer":
			printStandardIO("** Calling C code **\n")
			defered.UseDefer()
		case "enviroment":
			printStandardIO("** Getting runtime information **\n")
			goEnv.GetInfo()
		default:
			fmt.Println("DEFAULT")
		}
	}

}
