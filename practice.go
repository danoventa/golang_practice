package main

// factored import statements. 
import (
	"fmt"
	"math/rand"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
}