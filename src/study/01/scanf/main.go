package main

func  main()  {
	//整形
	//var num1,num2 int
	//fmt.Scanf("%d %d",&num1,&num2)

	//浮点型
	//var num3,num4 float64
	//fmt.Scanf("%f %f",&num3,&num4)
	//fmt.Printf("sum: %.2f",calcf(&num3,&num4))

	//字符串

	//var str1 string
	//var str2 string
	//fmt.Scanf("%s,%s",)

}

func calc(num1 *int,num2 *int)int{
	return  *num1+*num2
}

func  calcf(num1 *float64,num2 *float64)float64{
	return  *num1+*num2
}