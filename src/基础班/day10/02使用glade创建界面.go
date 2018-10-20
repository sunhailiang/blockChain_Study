package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"fmt"
	"github.com/mattn/go-gtk/gdk"
)

func main() {

	//gtk初始化
	gtk.Init(&os.Args)
	//创建builder  通过builder加载glade文件
	builder:=gtk.NewBuilder()
	builder.AddFromFile("D:/GoCode/src/UI.glade")

	win:=gtk.WindowFromObject(builder.GetObject("Win1"))
	win.SetTitle("glade界面")

	//通过glade创建label
	lab:=gtk.LabelFromObject(builder.GetObject("Lab1"))
	lab.SetText("澳门在线赌场开业了")
	//为标签设置字号和颜色
	lab.ModifyFontSize(20)
	lab.ModifyFG(gtk.STATE_NORMAL,gdk.NewColor("red"))


	//通过glade创建button
	btn:=gtk.ButtonFromObject(builder.GetObject("Btn1"))
	btn.SetLabel("点击进入")
	btn.Clicked(func(){
		fmt.Println("点击")
	})



	win.Show()

	gtk.Main()


}
