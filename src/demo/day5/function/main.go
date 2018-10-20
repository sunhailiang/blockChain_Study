package main
import(
	"fmt"
)

type person struct{
	gender string
	name string 
	age int
	score float32
}
//声明一个函数并只想person
func (p *person) init(gender string,name string,age int,score float32){
	p.gender=gender
	p.age=age
	p.name=name
	p.score=score

	fmt.Println(p)
}

func main(){
	var stu person
	stu.init("man","harry",27,95)
}
