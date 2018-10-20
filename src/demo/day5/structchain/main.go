package main
import(
	"fmt"
)
type detailInfo struct{
	 fileType int
	 fileName string
	 filePath string
	 fileNum int

	 //实现结构体链表
	 next *detailInfo
}
//从尾部继续添加
func insertTail(p *detailInfo){
	var tail=p;
	for i:=0;i<10;i++{
		stu:=detailInfo{
			fileType:i,
			fileName:"AAAAA",
			filePath:"d://dd/dd/dd/dd",
			fileNum:333,
		}
		tail.next=&stu
		tail=&stu
	}
}

//从头部添加
func insertfront(p **detailInfo){
	for i:=0;i<10;i++{
		stu:=detailInfo{
			fileType:i,
			fileName:"AAAAA",
			filePath:"d://dd/dd/dd/dd",
			fileNum:333,
		}
		stu.next=*p
		*p=&stu
	}
}
func main(){
	 var Head *detailInfo=new(detailInfo)
	//  var pngfile detailInfo
	
	 Head.fileType=1
	 Head.fileName="项目文档"
	 Head.filePath="d://documents/doc"
	 Head.fileNum=33
    //尾部插入
	//   insertTail(&Head);
	//从头部插入
	insertfront(&Head)
	  var  p *detailInfo=Head;
	  for p!=nil{
		  fmt.Println(*p)
		  //找到下一个阶段
		  p=p.next
	  }

}