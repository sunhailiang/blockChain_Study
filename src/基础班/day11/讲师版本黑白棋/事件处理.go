package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
	"unsafe"
)

func DrawWindowImagefromFile(ctx *glib.CallbackContext) {
	data := ctx.Data()

	obj, ok := data.(*ChessWidget)
	if ok == false {
		return
	}
	//设置画板
	paint := obj.Win.GetWindow().GetDrawable()
	gc := gdk.NewGC(paint)
	//加载图片到内存中
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("src/image/bg.jpg", obj.w, obj.h, false)
	blackpixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("src/image/black.png", obj.gridw-5, obj.gridh-5, false)
	whitepixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("src/image/white.png", obj.gridw-5, obj.gridh-5, false)
	//将图片绘制在窗体中
	paint.DrawPixbuf(gc, pixbuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	for i := 0; i < len(obj.board); i++ {
		for j := 0; j < len(obj.board[i]); j++ {
			//绘制黑棋和白棋
			if obj.board[i][j] == Black {
				paint.DrawPixbuf(gc, blackpixbuf, 0, 0, obj.startx+i*obj.gridw+3, obj.starty+j*obj.gridh+4,
					-1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			} else if obj.board[i][j] == White {
				paint.DrawPixbuf(gc, whitepixbuf, 0, 0, obj.startx+i*obj.gridw+3, obj.starty+j*obj.gridh+4,
					-1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			}
		}
	}

	//释放资源
	pixbuf.Unref()
	blackpixbuf.Unref()
	whitepixbuf.Unref()
}

func DrawButtonImageFromFile(btn *gtk.Button, name string) {
	//将纹理加载到内存中
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(name, 37, 37, false)

	//设置图片对象
	img := gtk.NewImage()
	img.SetFromPixbuf(pixbuf)
	//将纹理转成图片
	btn.SetImage(img)
	//释放资源
	pixbuf.Unref()

}

func BtnMinClick(ctx *glib.CallbackContext) {
	data := ctx.Data()
	obj, ok := data.(*ChessWidget)
	if ok == false {
		return
	}
	//最小化
	obj.Win.Iconify()
}
func BtnCloseClick(ctx *glib.CallbackContext) {
	data := ctx.Data()
	obj, ok := data.(*ChessWidget)
	if ok == false {
		return
	}
	//关闭计时器
	glib.TimeoutRemove(obj.NumTimerId)
	glib.TimeoutRemove(obj.LabTimerId)

	//关闭
	gtk.MainQuit()
}

//为图片添加纹理
func DrawImageFromFile(img *gtk.Image, name string) {
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(name, 50, 50, false)
	img.SetFromPixbuf(pixbuf)
	pixbuf.Unref()
}

func MouseClickEvent(ctx *glib.CallbackContext) {
	arg := ctx.Args(0)

	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
	//起始坐标是210 70 棋盘格子大小52 52
	//fmt.Println("X:",int(event.X),"Y:",int(event.Y))

	data := ctx.Data()
	obj, ok := data.(*ChessWidget)
	if ok == false {
		return
	}
	x, y := int(event.X), int(event.Y)

	//paint := obj.Win.GetWindow().GetDrawable()
	//gc := gdk.NewGC(paint)

	//blackpixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("src/image/black.png", obj.gridw-5, obj.gridh-5, false)
	//whitepixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("src/image/white.png", obj.gridw-5, obj.gridh-5, false)

	//判断是否是可绘制区域
	if x >= obj.startx && y >= obj.starty && x <= obj.endx && y <= obj.endy {
		////计算点击位置和起始位置的差值
		//offsetx:=x-obj.startx
		//offsety:=y-obj.starty
		////用差值除以格子的宽高
		//i:=offsetx/obj.gridw
		//j:=offsety/obj.gridh
		i := (x - obj.startx) / obj.gridw
		j := (y - obj.starty) / obj.gridh
		//判断棋盘如果为空可以绘制
		if obj.board[i][j] == Empty {

			//根据当前角色绘制图片
			//if obj.curRole == Black {
			//	paint.DrawPixbuf(gc, blackpixbuf, 0, 0, obj.startx+i*obj.gridw+3, obj.starty+j*obj.gridh+4,
			//		-1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			//
			//	//将棋子放在棋盘中
			//	//obj.board[i][j] = Black
			//} else if obj.curRole == White {
			//	paint.DrawPixbuf(gc, whitepixbuf, 0, 0, obj.startx+i*obj.gridw+3, obj.starty+j*obj.gridh+4,
			//		-1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			//	//将棋子放在棋盘中
			//	//obj.board[i][j] = White
			//}

			if obj.ChessRule(i, j, obj.curRole, true) > 0 {

				//更新界面信息
				obj.Win.QueueDraw()
				//切换当前角色
				obj.ChangeRole()
			}

		}

		//fmt.Println(j,i)
	}
	//blackpixbuf.Unref()
	//whitepixbuf.Unref()

}
