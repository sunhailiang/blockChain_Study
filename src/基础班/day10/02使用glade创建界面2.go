package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gdkpixbuf"
)

func main() {

	//gtk初始化
	gtk.Init(&os.Args)
	//创建builder  通过builder加载glade文件
	builder:=gtk.NewBuilder()
	builder.AddFromFile("D:/GoCode/src/UI.glade")

	win:=gtk.WindowFromObject(builder.GetObject("Win1"))
	win.SetTitle("glade界面")
	//设置窗口图标
	win.SetIconFromFile("D:/GoCode/src/image/cls.jpg")

	//窗体大小
	w,h:=0,0
	//获取窗体大小
	win.GetSize(&w,&h)

	fmt.Println("窗体宽 高",w,h)

	//设置窗体固定大小
	win.SetResizable(false)

	//通过glade创建label
	lab:=gtk.LabelFromObject(builder.GetObject("Lab1"))
	lab.SetText("澳门在线赌场开业了")
	//为标签设置字号和颜色
	lab.ModifyFontSize(15)
	lab.ModifyFG(gtk.STATE_NORMAL,gdk.NewColor("red"))
	//获取标签文本信息
	str:=lab.GetText()
	fmt.Println(str)

	//通过glade创建button
	btn:=gtk.ButtonFromObject(builder.GetObject("Btn1"))
	btn.SetLabel("点击进入")
	btn.Clicked(func(){
		fmt.Println("点击")
	})

	imgw,imgh:=0,0
	//获取Image控件
	img:=gtk.ImageFromObject(builder.GetObject("Img1"))
	//获取图片控件大小
	img.GetSizeRequest(&imgw,&imgh)

	//将图片加载到内存中
	pixbuf,_:=gdkpixbuf.NewPixbufFromFileAtScale("D:/GoCode/src/image/cls.jpg",imgw,imgh,false)

	img.SetFromPixbuf(pixbuf)

	//将内存中图片释放
	pixbuf.Unref()

	win.Show()

	//关闭窗体  事件处理
	win.Connect("destroy",func(ctx *glib.CallbackContext){
		gtk.MainQuit()
	})

	gtk.Main()


}
