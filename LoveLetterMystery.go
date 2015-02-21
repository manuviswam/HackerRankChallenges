package main
import "fmt"

func main(){
	var (
		n int
		tempStr string
	)
	fmt.Scan(&n)
	resultArray := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&tempStr)
		count:=0
		l := len(tempStr)
		for j:=0;j<l/2;j++{
			a,b := tempStr[j],tempStr[l-j-1]
			if a>b{
				a,b = b,a
			}
			count += int(b-a)
		}
		resultArray[i] = count
	}
	for i:=0;i<n;i++{
		fmt.Println(resultArray[i])
	}
}
