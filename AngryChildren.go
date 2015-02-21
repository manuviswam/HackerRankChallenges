package main
import (
	"fmt"
	"sort"
)


func main(){
	var(
		n int32
		k int32
		i int32
		curUnfairness float64
		unfairness float64
		temp float64
	)
	fmt.Scan(&n)
	fmt.Scan(&k)
	inputArray := make([]float64,n)
	for i=0;i<n;i++{
		fmt.Scan(&temp)
		inputArray[i] = temp
	}
	unfairness = 1.797693134862315708145274237317043567981e+308 //max value of flaot64
	sort.Float64s(inputArray)
	for i=0;i<=n-k;i++{
		curUnfairness = inputArray[i+k-1] - inputArray[i]
		if curUnfairness < unfairness{
			unfairness = curUnfairness
		}
	}
	fmt.Printf("%.f",unfairness)
}
