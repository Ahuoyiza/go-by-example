package main
import (
	"fmt"
	"math"
)
const s string =  "constant"

func main(){
	var x = "new var"
	fmt.Println(x)

	var a,b int = 1,2
	fmt.Println("a: ", a, "b: ", b)

	var c = true
	fmt.Println(c)

	var d string
	fmt.Println(d) // returns an empty string

	var e int
	fmt.Println(e)

	f := "scoped var"
	fmt.Println(f)


	fmt.Println(s)

	const n = 500000000

	const y = 3e20/n
	fmt.Println(y)
	fmt.Println(int64(y))

	fmt.Println(math.Sin(n))

}