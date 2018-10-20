package main

import "fmt"

func main() {
	var  nums =[5]int{1,2,3,4,5}
	var max int=nums[0]
	var min int=nums[0]
	var sum int
	for i:=0;i< len(nums);i++{
        if max<nums[i]{
        	max=nums[i]
		}
		if min>nums[i] {
			min=nums[i]
		}

		sum+=nums[i]
	}

	fmt.Printf("最大：%d,最小：%d,平均：%d,和：%d",max,min,sum/len(nums),sum)
}
