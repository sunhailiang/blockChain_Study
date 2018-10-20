package 单一功能测试

//import (
//	"bytes"
//	"encoding/gob"
//	"log"
//	"fmt"
//)
//
//type person struct {
//	Name string
//	Age  int
//}
//
//func main() {
//	var xiaoming = person{"小明", 18}
//	//编码结果容器
//	buffer := bytes.Buffer{}
//	//gob创建编码器
//	encoder := gob.NewEncoder(&buffer)
//	//编码
//	err := encoder.Encode(&xiaoming)
//	if err != nil {
//		log.Panic("编码出错",err)
//	}
//	fmt.Println("编码后:", buffer.Bytes())
//	//创建解码器
//	decoder := gob.NewDecoder(&buffer)
//	//解码
//	err = decoder.Decode(&xiaoming)
//	if err != nil {
//		log.Panic("解码出错")
//	}
//	fmt.Println(xiaoming)
//}
