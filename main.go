//go 1.19.4

package main

import "fmt"

func main() {
	var x, y, z int
	var control string

	fmt.Print("Please enter the first number: ")
	fmt.Scan(&x)

	fmt.Print("Please enter the second numbers: ")
	fmt.Scan(&y)
	fmt.Println("Please Enter A Number. Numbers 1,2,3,4 are for first four main equations. 5 6 7 8 are what if these numbers place is switched")
	fmt.Scan(&z)

	var d, d1 int
	if x == 0 || y == 0 {
		fmt.Println("Cannot divide by 0")
	} else {
		d = x / y
		d1 = y / x
	}

	a := x + y
	b := x - y
	c := x * y

	a1 := y + x
	b1 := y - x
	c1 := y * x

	switch z {
	case 1:
		fmt.Println(a)
		fmt.Scan(&control)
	case 2:
		fmt.Println(b)
		fmt.Scan(&control)
	case 3:
		fmt.Println(c)
		fmt.Scan(&control)
	case 4:
		fmt.Println(d)
		fmt.Scan(&control)
	case 5:
		fmt.Println(a1)
		fmt.Scan(&control)
	case 6:
		fmt.Println(b1)
		fmt.Scan(&control)
	case 7:
		fmt.Println(c1)
		fmt.Scan(&control)
	case 8:
		fmt.Println(d1)
		fmt.Scan(&control)
	default:
		fmt.Println("İnvalid")
		fmt.Scan(&control)
	}
}
