package main
import "fmt"
func main(){
	var size int
	fmt.Println()("Enter the size of the array")
	fmt.Scanf(&size)
  arr:=make([]int,size)
  for i:=0,i<size;i++ {
	fmt.Println("Enter element %d: ",i+1)
	fmt.Scan(&arr[i])
  }

}
