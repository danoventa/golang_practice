package main

// factored import statements. 
import (
	"fmt"
	"math/rand"
	"math"
)

func add(x, y int) int { // you can do this instead of (x int, y int)
	return x + y
}

func swap(x, y string) (string, string) { // can return tuples! Just like python and swift ! :D
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17)) // naked returns! 
}