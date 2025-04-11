package main
import "fmt"

func hitungRataDigit(a, b *int){
	var i, hasil int
	var c float64
	*b = 0
	for *a> 0{
		i = *a%10
		hasil += i
		*b++
		*a = *a/10
	}
	c = float64(hasil)/float64(*b)
	fmt.Print(c)
}
func main(){
	var a, b int
	fmt.Scan(&a)
	hitungRataDigit(&a, &b)
}
