package main

import (
	"fmt"
	"myproject/sMaps"
)

type name struct {
	Fname string
	Lname string
}

func main() {

	// NOV 3
	sMaps.MapFunc()

	//defer fmt.Println("ADD ->> ", simplecalc.Add(3, 5))

	//defer statement(5) // defer

	//STRUCT
	n := name{
		Fname: "Dinesh",
		Lname: "Kolla",
	}
	fmt.Print("Hello ", n.Fname, " ", n.Lname, "\n")

	//POINTER
	X := 10
	var p = &X
	fmt.Println(p, "-->", X)
	*p = 9
	fmt.Println("Changed value -->", X)

	//Array
	var arr [5]int = [5]int{11, 12, 13, 14, 15}
	fmt.Println("Index", " ", "value")
	for index, val := range arr {
		fmt.Println(index, " --->", val)
	}

}
func statement(x int) {
	// FOR LOOP
	for i := 1; i <= 10; i++ {
		fmt.Print(x*i, " ")
	}
	fmt.Println()
	//SWITH
	switch x {
	case 1:
		fmt.Println("Hi")
	case 2:
		fmt.Println("Hello")
	case 3:
		fmt.Println("Bye")
	default:
		fmt.Println("Sorry")
	}

	//IF- ELSE
	if x/2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

}
