package main
import "fmt"
type Read interface {
	read()
}
type Write interface {
	write()
}
type setDoFile interface {
	Read
	Write
}
func (f *file)read()  {
	fmt.Println("read")
}
func (f *file)write()  {
	fmt.Println("write")
}
type file struct {
	fileName string
}
func  main()  {
    var f file=file{
    	fileName:"AV",
	}
	var setinterface setDoFile=&f
	setinterface.read()
	setinterface.write()
}
