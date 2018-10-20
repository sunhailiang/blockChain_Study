package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"fmt"
)

func main() {
	gtk.Init(&os.Args)

	builder:=gtk.NewBuilder()
	builder.AddFromFile("E:/go/src/gtk/UI/UI.glade")


	win:=gtk.WindowFromObject(builder.GetObject("window1"))
	win.SetSizeRequest(480,320)


	//获取垂直布局
	hbox:=gtk.HBoxFromObject(builder.GetObject("hbox1"))
	//创建按钮
	btn:=gtk.NewButtonWithLabel("点击进入")
	btn.Clicked(func(){fmt.Println("点击")})
	//将按钮添加到布局中
	hbox.Add(btn)
	win.ShowAll()
	//点击关闭按钮退出程序
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	gtk.Main()

}


func main日日日(){
	gtk.Init(&os.Args)

	builder:=gtk.NewBuilder()
	builder.AddFromFile("E:/go/src/gtk/UI/UI1.glade")

	win:=gtk.WindowFromObject(builder.GetObject("window1"))

	win.SetSizeRequest(480,320)



	//水平布局
	vbox:=gtk.VBoxFromObject(builder.GetObject("vbox1"))

	//表格布局
	//gtk.TableFromObject(builder.GetObject("table1"))
	lab:=gtk.NewLabel("XX社区")
	vbox.Add(lab)

	win.ShowAll()
	//点击关闭按钮退出程序
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	gtk.Main()
}
