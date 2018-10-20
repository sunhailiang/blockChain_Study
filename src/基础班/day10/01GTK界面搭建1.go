package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"github.com/mattn/go-gtk/glib"
	"fmt"
)

func main() {

	//初始化
	gtk.Init(&os.Args)

	//通过GTK创建界面
	win:=gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	//设置窗体大小
	win.SetSizeRequest(480,320)
	//设置窗体标题
	win.SetTitle("GTK窗口")



	//创建按钮
	b:=gtk.NewButton()
	//设置按钮大小
	b.SetSizeRequest(100,80)
	//设置按钮标题
	b.SetLabel("点击按钮")
	//设置按钮点击事件
	//b.Connect("clicked",BtnClick,"点击")
	//设置按钮点击事件
	//b.Clicked(BtnClick,"点击")
	//匿名函数  实现按钮点击事件处理
	//b.Clicked(func(ctx *glib.CallbackContext){
	//	data:=ctx.Data()
	//	d,ok:=data.(string)
	//	if ok{
	//		fmt.Println(d)
	//	}
	//},"点击")

	//released事件处理
	//b.Connect("released",BtnClick,"松开")
	//pressed事件处理
	b.Connect("pressed",BtnClick,"按下")

	//创建布局，将按钮放在布局上
	layout:=gtk.NewFixed()
	//layout.Add(b)
	//放在布局的指定位置
	layout.Put(b,190,120)

	//将布局放在界面中显示
	win.Add(layout)


	//将界面显示
	//win.Show()
	win.ShowAll()
	//调用gtk main  运行程序
	gtk.Main()

}

func BtnClick(ctx *glib.CallbackContext){
	//获取点击事件传递的数据
	data:=ctx.Data()
	//对数据进行断言操作
	d,ok:=data.(string)
	if ok{
		fmt.Println(d)
	}
}
