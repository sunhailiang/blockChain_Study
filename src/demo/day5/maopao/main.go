package main
 import(
	 "fmt"
 )
 func main(){
	arr:=[...]int{2,33,1,44,666,3,6,8};
	mp(arr[:])
	fmt.Println(arr);
 }

 func mp(arr[] int){
      for i:=0;i<len(arr);i++{
		  for j:=1;j<len(arr)-i;j++{
			  if arr[j]<arr[j-1]{
                  arr[j],arr[j-1]=arr[j-1],arr[j]
			  }
		  }
	  }
 }