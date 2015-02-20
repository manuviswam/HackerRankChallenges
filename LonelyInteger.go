package main
import (
	"fmt"
	"sort"
)

func main(){
	var (
		n int
		i int
		x int
	)
	fmt.Scan(&n)
	var array = make([]int,n)
	for i=0; i<n; i++{
		fmt.Scan(&x)
		array[i]=x
	}
	sort.Ints(array)
	for i=0; i<n; i+=2{
		if i== n-1 || array[i] != array[i+1] {
			fmt.Println(array[i])
			break;
		}
	}
}
