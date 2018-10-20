package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"github.com/mattn/go-gtk/glib"
	"unsafe"
	"github.com/mattn/go-gtk/gdk"
	"fmt"
	"github.com/mattn/go-gtk/gdkpixbuf"
)

func main0201() {
	gtk.Init(&os.Args)

	win:=gtk.NewWindow(gtk.WINDOW_TOPLEVEL)

	win.SetSizeRequest(480,320)
	win.SetTitle("常用事件")

	//设置键盘事件
	//key-press-event   键盘按下事件
	//key-release-event  键盘松开事件
	win.Connect("key-release-event", func(ctx *glib.CallbackContext) {
		arg:=ctx.Args(0)
		//获取键盘按下key
		event:=*(**gdk.EventKey)(unsafe.Pointer(&arg))
		key:=event.Keyval

		switch key {
		case gdk.KEY_a:
			fmt.Println("a键按下")
		case gdk.KEY_d:
			fmt.Println("d键按下")
		}

	})

	win.ShowAll()

	gtk.Main()

}
func main0202(){
	gtk.Init(&os.Args)

	win:=gtk.NewWindow(gtk.WINDOW_TOPLEVEL)

	win.SetSizeRequest(480,320)
	win.SetTitle("常用事件")

	//绘图事件
	//允许在屏幕上进行绘制
	win.SetAppPaintable(true)


	win.Connect("expose-event",func(){
		//设置绘制区域
		paint:=win.GetWindow().GetDrawable()
		gc:=gdk.NewGC(paint)
		//加载一张图片到内存中
		pixbuf,_:=gdkpixbuf.NewPixbufFromFileAtScale("D:/GoCode/src/image/cls.jpg",
			480,320,false)

		//将图片绘制在win上
		paint.DrawPixbuf(gc,pixbuf,0,0,0,0,-1,-1,gdk.RGB_DITHER_NONE,0,0)

		//将内存中图片释放
		pixbuf.Unref()
	})


	win.ShowAll()

	gtk.Main()
}

func main(){
	gtk.Init(&os.Args)

	win:=gtk.NewWindow(gtk.WINDOW_TOPLEVEL)

	win.SetSizeRequest(480,320)
	win.SetTitle("常用事件")

	//鼠标按下事件
	win.SetEvents(int(gdk.BUTTON_PRESS_MASK))
	win.Connect("button-press-event",func(ctx *glib.CallbackContext){
		arg:=ctx.Args(0)

		event:=*(**gdk.EventButton)(unsafe.Pointer(&arg))

		//判断是单机还是双击事件
		//if event.Type == int(gdk.BUTTON_PRESS){
		//	fmt.Println("鼠标单机")
		//}else if event.Type == int(gdk.BUTTON2_PRESS){
		//	fmt.Println("鼠标双击")
		//}
		//获取鼠标点击位置

		fmt.Println("X:",int(event.X),"Y:",int(event.Y))
	})
	win.ShowAll()

	gtk.Main()
}