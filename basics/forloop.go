package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
	numbers:= []int{10,20,30,40,50}
	for index,value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i:= range 10 {
		fmt.Println(i)  // this feature is very similar to python
	}

	// no while loop in go
	j := 1
	for j<=5 {
		fmt.Println("While Loop Iteration:", j)
		j++
	}
	for {
		fmt.Println("Infinite Loop")
	}
}