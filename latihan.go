package main
import "fmt"

func faktorBil(n int)int{
	var i, num int
	num = 0
	i = 1
	for i <= n{
		if n%i == 0{
			num++
		}
		i++
	}
	return num
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Print(faktorBil(n))
}