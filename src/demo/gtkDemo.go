package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)//初始化
	win:=gtk.NewWindow(gtk.WINDOW_TOPLEVEL)//带边框的顶层窗口
	win.SetTitle("go_gtk")//标题
	win.SetSizeRequest(500,400)//大小
	win.Show()//显示
	gtk.Main()
}
