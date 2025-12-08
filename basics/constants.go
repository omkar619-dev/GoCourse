package main
import "fmt"

const Pi = 3.14
const GRAVITY = 9.8

func const() {
	const days int = 7
	const(
		monday = 1
		tuesday = 2
	)
	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Value of Gravity:", GRAVITY)
}