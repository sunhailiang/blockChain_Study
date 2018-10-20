package main
import(
	"fmt"
)
func main(){
   arr:=[...]int{22,3,4,55,6,77,34,5,1};
   selected(arr[:])
   fmt.Println(arr)
}
func selected(arr []int){
    
	  for i:=0;i<len(arr);i++{
		  min:=i
		  for j:=i+1;j<len(arr);j++{
			  if arr[min]>arr[j]{
				  min=j;
			  }
		  }
		  arr[i],arr[min]=arr[min],arr[i]
	  }

}
