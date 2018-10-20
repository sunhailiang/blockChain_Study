// loadsrc
package main

import (
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

//加载按钮图片信息
func SetButtonFromFile(btn *gtk.Button, name string) {
	w, h := 0, 0
	btn.GetSizeRequest(&w, &h)
	//fmt.Println(w, h)
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(name, w-10, h-10, false)
	img := gtk.NewImageFromPixbuf(pixbuf)
	//按钮设置纹理需要图片
	btn.SetImage(img)
	//释放资源
	pixbuf.Unref()
	//取消按钮焦点
	btn.SetCanFocus(false)
}

//加载图片信息
func SetImageFromFile(img *gtk.Image, name string) {
	w, h := 0, 0
	img.GetSizeRequest(&w, &h)
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(name, w, h, false)
	img.SetFromPixbuf(pixbuf)
	pixbuf.Unref()
}
