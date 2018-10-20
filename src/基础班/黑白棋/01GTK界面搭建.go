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
	//将界面显示
	win.Show()

	//调用gtk main  运行程序
	gtk.Main()

}
