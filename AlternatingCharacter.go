package main
import (
	"fmt"
)

func main(){
	var (
		n int
		temp string
	)
	fmt.Scan(&n)
	countArray := make([]int,n)
	for i:=0; i<n; i++{
		fmt.Scan(&temp)
		count := 0
		for j:=0;j<len(temp) -1;j++{
			if (temp[j] == temp[j+1]){
				count++
			}
		}
		countArray[i] = count
	}
	for i:=0;i<n;i++{
		fmt.Println(countArray[i])
	}
}
