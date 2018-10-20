package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
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
