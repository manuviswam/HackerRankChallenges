package main
import (
	"fmt"
	"math"
)

func main(){
	var (
		n int
		temp int
	)
	fmt.Scan(&n)
	array := make([]uint,n)
	for i:=0;i<n;i++{
		fmt.Scan(&temp)
		if temp%2==1 {
			array[i] =uint(math.Pow(2,float64(((temp+1)/2)+1)) - 2)
		}else{
			array[i] = uint(math.Pow(2,float64(((temp)/2)+1)) -1)
		}
	}
	for i:=0;i<n;i++{
		fmt.Println(array[i])
	}
}
