package main
import (
	"fmt"
)

func main(){
	var(
		n int
		i int
		temp uint
	)
	fmt.Scan(&n)
	var array = make([]uint,n)
	for i=0;i<n; i++{
		fmt.Scan(&temp)
		array[i] = (4294967295 - temp)
	}
	for i=0; i<n; i++{
		fmt.Println(array[i])
	}
}
