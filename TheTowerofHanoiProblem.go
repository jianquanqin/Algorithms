package BLC

import "fmt"

//Test example

//func main() {
//	t := 3
//	a, b, c := "A", "B", "C"
//
//	hanoi(t, a, b, c)
//
//}

func hanoi(n int, a, b, c string) {

	if n > 0 {
		hanoi(n-1, a, c, b)
		fmt.Printf("moving from %s to %s\n", a, c)
		hanoi(n-1, b, a, c)
	}
}
