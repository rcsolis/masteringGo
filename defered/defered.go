package defered

import "fmt"

func countUp(){
	fmt.Println("Counting Up.")
	for i:=0; i<3; i++{
		fmt.Println(i)
	}
}

func countDown(){
	fmt.Println("Countig Down.")
	for i:=1; i<3; i++{
		defer fmt.Print(i, " ")
	}
	for i:=3; i>0;i--{
		defer func(number int) {
			fmt.Print(number, " ")
		}(i)
	}
}

func UseDefer(){
	fmt.Println("--- Using defered funcitons")
	fmt.Println("Start surrounding function")
	defer func (){
		fmt.Println("Defer postpone the execution of a function until the surrounding function returns")
	}()
	defer func (){
		fmt.Println("Defered functions executes in LIFO")
	}()
	defer countUp()
	fmt.Println()
	countDown()
	fmt.Println()
	fmt.Println("Return using defered functions")
}
