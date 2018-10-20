// Reversi.go
package main

import (
	"os"
	"strconv"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//创建结构体存储控件
type ChessWidget struct {
	Win      *gtk.Window //窗口
	BtnMin   *gtk.Button //最小化按钮
	BtnClose *gtk.Button //最大化按钮
	ImgBlack *gtk.Image  //黑色闪烁棋子
	ImgWhite *gtk.Image  //白色闪烁棋子
	LabBlack *gtk.Label  //黑色棋子个数
	LabWhite *gtk.Label  //白色棋子个数
	LabTime  *gtk.Label  //计时器时间
}

//创建结构体存储信息
type ChessInfo struct {
	x      int //获取点击位置x
	y      int //获取点击位置y
	w      int //窗体大小w
	h      int //窗体大小h
	StartX int //棋盘起始位置x 210
	StartY int //棋盘起始位置y 70
	GridW  int //棋盘格子大小w 52
	GridH  int //棋盘格子大小h 52
}

//枚举  表示棋盘存储棋子信息
const (
	Empty = iota //0
	Black = iota //1
	White = iota //2
)

//匿名字段 实现继承关系
type ChessBoard struct {
	ChessWidget
	ChessInfo

	chess       [8][8]int //存储棋盘信息 3种
	CurRole     int       //当前执子角色
	TipTimerId  int       //执子角色图标计时器
	RoleTimerId int       //执子角色时间计时器
	TimeNum     int       //当前剩余时间
}

//控件创建和加载
func (obj *ChessBoard) CreateWidget() {
	builder := gtk.NewBuilder()
	builder.AddFromFile("ui.glade")

	obj.Win = gtk.WindowFromObject(builder.GetObject("window1"))
	//不允许修改大小
	obj.Win.SetResizable(false)
	//获取窗体大小
	obj.Win.GetSizeRequest(&obj.w, &obj.h)
	//设置窗口无标题
	obj.Win.SetDecorated(false)
	//设置窗体允许绘制
	obj.Win.SetAppPaintable(true)
	//设置窗体的鼠标点击事件和移动事件
	obj.Win.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
	obj.Win.SetTitle("黑白棋")

	//加载按钮信息
	obj.BtnMin = gtk.ButtonFromObject(builder.GetObject("BtnMin"))
	obj.BtnClose = gtk.ButtonFromObject(builder.GetObject("BtnClose"))
	SetButtonFromFile(obj.BtnMin, "./image/min.png")
	SetButtonFromFile(obj.BtnClose, "./image/close.png")

	//加载标签信息
	obj.LabBlack = gtk.LabelFromObject(builder.GetObject("LabBlack"))
	obj.LabWhite = gtk.LabelFromObject(builder.GetObject("LabWhite"))
	obj.LabTime = gtk.LabelFromObject(builder.GetObject("LabTime"))
	obj.LabBlack.ModifyFontSize(40)
	obj.LabWhite.ModifyFontSize(40)
	obj.LabTime.ModifyFontSize(20)

	//加载图片信息
	obj.ImgBlack = gtk.ImageFromObject(builder.GetObject("ImgBlack"))
	obj.ImgWhite = gtk.ImageFromObject(builder.GetObject("ImgWhite"))
	SetImageFromFile(obj.ImgBlack, "./image/black.png")
	SetImageFromFile(obj.ImgWhite, "./image/white.png")
	//初始化设为隐藏
	obj.ImgBlack.Hide()
	obj.ImgWhite.Hide()

}

func (obj *ChessBoard) Handle() {

	//加载背景界面
	obj.Win.Connect("expose-event", PaintBg, obj)
	//处理鼠标点击事件
	obj.Win.Connect("button-press-event", MousePressEvent, obj)
	//处理按钮事件
	obj.BtnMin.Connect("clicked", BtnMinEvent, obj)
	obj.BtnClose.Connect("clicked", BtnCloseEvent, obj)
}

//棋盘初始化操作
func (obj *ChessBoard) InitChess() {
	//设置起始位置和格子大小
	obj.StartX = 210
	obj.StartY = 70
	obj.GridW = 52
	obj.GridH = 52

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			obj.chess[i][j] = Empty
		}
	}
	//初始化棋盘棋子
	obj.chess[3][3] = Black
	obj.chess[4][4] = Black

	obj.chess[3][4] = White
	obj.chess[4][3] = White
	//默认黑子 执棋
	obj.CurRole = Black
	obj.TimeNum = 20

	//默认棋子个数
	obj.LabBlack.SetText("2")
	obj.LabWhite.SetText("2")
	obj.Win.QueueDraw()

	//定义计时器
	obj.TipTimerId = glib.TimeoutAdd(500, func() bool {
		if obj.CurRole == Black {
			obj.ImgWhite.Hide()
			if obj.ImgBlack.GetVisible() == true {
				obj.ImgBlack.Hide()
			} else {
				obj.ImgBlack.Show()
			}
		} else {
			obj.ImgBlack.Hide()
			if obj.ImgWhite.GetVisible() == true {
				obj.ImgWhite.Hide()
			} else {
				obj.ImgWhite.Show()
			}
		}
		return true
	})

	obj.RoleTimerId = glib.TimeoutAdd(1000, func() bool {
		obj.TimeNum--

		//切换角色
		if obj.TimeNum == 0 {

			ChangeRole(obj)
		}

		//为标签设置值  整型数据转成string类型
		obj.LabTime.SetText(strconv.Itoa(obj.TimeNum))
		return true
	})
}

func main() {
	gtk.Init(&os.Args)
	var obj ChessBoard
	//加载控件信息
	obj.CreateWidget()
	//处理控件事件
	obj.Handle()
	//初始化棋盘
	obj.InitChess()
	//显示所有控件
	obj.Win.ShowAll()
	gtk.Main()
}
