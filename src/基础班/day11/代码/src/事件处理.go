// handleenent
package main

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//最小化事件处理
func BtnMinEvent(ctx *glib.CallbackContext) {
	data := ctx.Data()
	obj, ok := data.(*ChessBoard)
	if ok == false {
		fmt.Println("无效操作")
		return
	}

	//最小化
	obj.Win.Iconify()

}

//关闭事件处理
func BtnCloseEvent(ctx *glib.CallbackContext) {
	data := ctx.Data()
	obj, ok := data.(*ChessBoard)
	if ok == false {
		fmt.Println("无效操作")
		return
	}
	//关闭
	gtk.MainQuit()
	//关闭计时器
	glib.TimeoutRemove(obj.TipTimerId)
	glib.TimeoutRemove(obj.RoleTimerId)
}

//主界面背景事件处理

func PaintBg(ctx *glib.CallbackContext) {
	data := ctx.Data()
	obj, ok := data.(*ChessBoard)
	if ok == false {
		fmt.Println("无效操作")
		return
	}

	//设置绘图
	paint := obj.Win.GetWindow().GetDrawable()
	gc := gdk.NewGC(paint)

	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/bg.jpg", obj.w, obj.h, false)
	//在窗体绘制纹理
	paint.DrawPixbuf(gc, pixbuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	//初始化黑子和白子
	blackpixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/black.png", obj.GridW-4, obj.GridH-4, false)
	whitepixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/white.png", obj.GridW-4, obj.GridH-4, false)

	//找到具体位置 进行绘制
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if obj.chess[i][j] == Black {
				paint.DrawPixbuf(gc, blackpixbuf, 0, 0, obj.StartX+i*obj.GridW+2, obj.StartY+j*obj.GridH+2, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			} else if obj.chess[i][j] == White {
				paint.DrawPixbuf(gc, whitepixbuf, 0, 0, obj.StartX+i*obj.GridW+2, obj.StartY+j*obj.GridH+2, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			}
		}
	}
	//销毁纹理
	pixbuf.Unref()
	blackpixbuf.Unref()
	whitepixbuf.Unref()

}

//窗体点击事件处理
func MousePressEvent(ctx *glib.CallbackContext) {

	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
	//event.x 点击横坐标 event.y 纵坐标
	//fmt.Println(event.X, event.Y)

	data := ctx.Data()
	obj, ok := data.(*ChessBoard)
	if ok == false {
		fmt.Println("无效操作")
		return
	}
	obj.x, obj.y = int(event.X), int(event.Y)

	//fmt.Println(obj.x, obj.y)
	//计算点击位置
	//obj.x - obj.StartX  计算点击位置和起始位置的横坐标差值
	i := (obj.x - obj.StartX) / obj.GridW
	j := (obj.y - obj.StartY) / obj.GridH

	if i >= 0 && i <= 7 && j >= 0 && j <= 7 {
		//二维数组下标
		//fmt.Println(j, i)

		if obj.ChessRule(i, j, obj.CurRole, true) > 0 {
			//按钮点击后设置棋子信息
			//obj.chess[i][j] = obj.CurRole

			//改变执子角色
			ChangeRole(obj)

			obj.ChessResult()
			//更新界面
			obj.Win.QueueDraw()
		}
		//如果没有落子位置 自动切换下一个用户
	}

}

func ChangeRole(obj *ChessBoard) {
	//重置时间
	obj.TimeNum = 20
	//切换角色
	if obj.CurRole == Black {
		obj.CurRole = White
	} else {
		obj.CurRole = Black
	}
}

//定义棋盘的规则  实现落子
func (obj *ChessBoard) ChessRule(x, y int, role int, eatChess bool) (eatNum int) {
	//如果该位置有棋子
	if obj.chess[x][y] != Empty {
		return 0
	}
	//定义棋盘八个方向
	dir := [8][2]int{{0, 1}, {1, 0}, {1, -1}, {-1, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}}
	tempx, tempy := x, y

	for i := 0; i < 8; i++ {
		tempx += dir[i][0]
		tempy += dir[i][1]
		//是否越界 是否是空的 是否是自己的 如果是对方棋子继续
		if tempx >= 0 && tempx <= 7 && tempy >= 0 && tempy <= 7 && obj.chess[tempx][tempy] != role && obj.chess[tempx][tempy] != Empty {
			//往下走
			tempx += dir[i][0]
			tempy += dir[i][1]
			for tempx >= 0 && tempx <= 7 && tempy >= 0 && tempy <= 7 {
				//如果为空遍历其他路线
				if obj.chess[tempx][tempy] == Empty {
					break
				}
				//判断是否可以吃掉
				if obj.chess[tempx][tempy] == role {
					if eatChess == true {
						//原始位置可以使用
						obj.chess[x][y] = role
						//回滚到原始坐标
						tempx -= dir[i][0]
						tempy -= dir[i][1]
						for (tempx != x) || (tempy != y) {
							obj.chess[tempx][tempy] = role
							tempx -= dir[i][0]
							tempy -= dir[i][1]
							eatNum++
						}
					} else {
						tempx -= dir[i][0]
						tempy -= dir[i][1]
						for (tempx != x) || (tempy != y) {
							tempx -= dir[i][0]
							tempy -= dir[i][1]
							eatNum++
						}
					}
					break
				}
				//继续向下遍历
				tempx += dir[i][0]
				tempy += dir[i][1]
			}

		}
		//改变其他路线
		tempx = x
		tempy = y
	}

	return
}
func (obj *ChessBoard) ChessResult() {
	var result string
	isOver := true
	//设置当前个数
	BlackNum, WhiteNum := 0, 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if obj.chess[i][j] == Black {
				BlackNum++
			} else if obj.chess[i][j] == White {
				WhiteNum++
			}

			if obj.ChessRule(i, j, Black, false) > 0 || obj.ChessRule(i, j, White, false) > 0 {
				isOver = false
			}
		}
	}
	obj.LabBlack.SetText(strconv.Itoa(BlackNum))
	obj.LabWhite.SetText(strconv.Itoa(WhiteNum))
	if isOver {
		if BlackNum > WhiteNum {
			result = "黑棋玩家赢"
		} else {
			result = "白棋玩家赢"
		}

		//新建消息对话框，选择对话框
		dialog := gtk.NewMessageDialog(
			obj.Win.GetTopLevelAsWindow(), //指定父窗口
			gtk.DIALOG_MODAL,              //模态对话框
			gtk.MESSAGE_QUESTION,          //指定对话框类型
			gtk.BUTTONS_YES_NO,            //默认按钮
			result)                        //设置内容

		dialog.SetTitle("结果") //对话框设置标题

		flag := dialog.Run() //运行对话框，返回值为按下的按钮类型
		if flag == gtk.RESPONSE_YES {
			//fmt.Println("继续")
			glib.TimeoutRemove(obj.TipTimerId)
			glib.TimeoutRemove(obj.RoleTimerId)
			//重置界面
			obj.InitChess()
		} else if flag == gtk.RESPONSE_NO {
			glib.TimeoutRemove(obj.TipTimerId)
			glib.TimeoutRemove(obj.RoleTimerId)
			gtk.MainQuit()
		} else {
			glib.TimeoutRemove(obj.TipTimerId)
			glib.TimeoutRemove(obj.RoleTimerId)
			//重置界面
			obj.InitChess()
		}

		dialog.Destroy() //销毁对话框

	}

}
