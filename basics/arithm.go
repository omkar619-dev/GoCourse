package main
import "fmt"
import "math"

func main(){
	var a,b int = 10,5
	var result int

	result = a + b
	fmt.Println("Addition:", result)
	result = a - b
	fmt.Println("Subtraction:", result)
	result = a * b
	fmt.Println("Multiplication:", result)
	result = a / b
	fmt.Println("Division:", result)
	result = a % b
	fmt.Println("Modulus:", result)
	var smallF float64 = 1.0-323
	smallF = smallF/math.MaxFloat64
	fmt.Println("Small Float:", smallF)
}