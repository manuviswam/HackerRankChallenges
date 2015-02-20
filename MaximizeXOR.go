package main
import (
	"fmt"
)

func main(){
	var (
		l int
		r int
		max int
	)
	max = 0
	fmt.Scan(&l)
	fmt.Scan(&r)
	if l>r{
		l=l+r
		r=l-r
		l=l-r
	}
	for i:=l;i<=r;i++{
		for j:=i;j<=r;j++{
			temp := i ^ j
			if temp > max{
				max = temp
			}
		}
	}
	fmt.Println(max)
}
