package main
import "fmt"
//1、接口的实现
type USBer interface {
	Read()
	Write()
}
//2、创建对象
type USBDev struct {
	id int
	name string
	rspeed int
	wspeed int
}
type Mobile struct {
	USBDev
	call string
}
type Upan struct {
	USBDev
}
//3、实现方法
func (m * Mobile)Read(){
	fmt.Printf("%s正在读取数据速度为：%d\n",m.name,m.rspeed)
}
func (m *Mobile)Write(){
	fmt.Printf("%s正在写入数据速度为：%d\n",m.name,m.wspeed)
}
func (u *Upan)Read(){
	fmt.Printf("%s正在读取数据速度为：%d\n",u.name,u.rspeed)
}
func (u *Upan)Write(){
	fmt.Printf("%s正在写入数据速度为：%d\n",u.name,u.wspeed)
}
//4、多态实现  将接口作为函数参数
func UseDev(usb USBer){
	usb.Read()
	usb.Write()
}
func main() {
	//接口类型
	var usb USBer
	usb=&Mobile{USBDev{101,"手机",5,10},"隔壁老王"}
	UseDev(usb)
	usb=&Upan{USBDev{102,"U盘",20,30}}
	UseDev(usb)

}
