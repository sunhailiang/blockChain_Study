package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"github.com/mattn/go-gtk/gdk"
	"fmt"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
)
func main() {
	gtk.Init(&os.Args) //初始化
	//创建builder加载glade
	builder := gtk.NewBuilder()
	//加载控件
	builder.AddFromFile("E:/go/src/gtk/UI/01demo.glade")
	//获取主窗体
	win := gtk.WindowFromObject(builder.GetObject("window1"))
	win.SetTitle("gdk第一个界面")
	//获取窗体控件
	lab := gtk.LabelFromObject(builder.GetObject("lab1"))
	lab.SetText("第一个操蛋的界面")
	lab.ModifyFontSize(20)
	lab.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("red"))
	//获取按钮
	btn := gtk.ButtonFromObject(builder.GetObject("btn1"))
	btn.Clicked(func() {
		fmt.Println("fuck")
	})
	//设置窗体固定大小
	win.SetResizable(false)
	//获取img控件
	img := gtk.ImageFromObject(builder.GetObject("img1"))
	//获取图片的款高
	imgw, imgh := 0, 0
	img.GetSizeRequest(&imgw, &imgh)
	//将图片加载到内存中
	imgbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("E:/avatar.png", imgw, imgh, false)
	img.SetFromPixbuf(imgbuf)
	//释放内存
	imgbuf.Unref()

	//真正关闭窗体
	win.Connect("destroy", func(ctx *glib.CallbackArg) {
		//主函数退出
		gtk.MainQuit()
	})
	win.Show()
	gtk.Main()
}
