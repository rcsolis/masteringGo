package environmentInfo

import (
	"fmt"
	"runtime"
	"strings"
	"strconv"
)

func GetInfo(){
	fmt.Println("--- Getting enviroment information")
	fmt.Println("You are using: ", runtime.Compiler, " on a machine:",runtime.GOARCH )
	fmt.Println("Your go version is:", runtime.Version())
	fmt.Println("You have CPUS:", runtime.NumCPU())
	fmt.Println("You are running Gorutines:", runtime.NumGoroutine())
	fmt.Println("You are running C Go Call:", runtime.NumCgoCall())

	goVersion := runtime.Version()
	major := strings.Split(goVersion, ".")[0][2]
	minor := strings.Split(goVersion, ".")[1]

	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 12 {
		fmt.Println("Need Go version 1.12 or higher!")
	}else{
		fmt.Println("You have version 1.12 or higher")
	}

	fmt.Println(" ---- Ends enviroment information")
}
