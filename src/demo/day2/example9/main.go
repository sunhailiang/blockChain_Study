package main
import "fmt"
func main(){
	str:="fuck";
	res:=reverse(str);
    fmt.Println("reverse",res);
}

//字符串反转
func reverse(str string) string{
   var result string;
	for i:=0;i<len(str);i++{
        result+=fmt.Sprintf("%c",str[len(str)-i-1]);
	}
	return result
}