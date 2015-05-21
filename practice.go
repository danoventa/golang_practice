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

var temporary = 100

func doSomething(bleh bool) (x, y string) {
	if (bleh == true) {
		x = "test"
		y = "test"
	} else {
		x = "nonono"
		y = "nonono"
	}
	return
} 

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
	var x, t int = 1, 2
	var pyts, jaba, s = true, false, "no!"
	fmt.Println(x, t, s, pyts, jaba)

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("this", "that")
	fmt.Println(a, b)
	fmt.Println(split(17)) // naked returns! 
	
	fmt.Println(doSomething(false))
}