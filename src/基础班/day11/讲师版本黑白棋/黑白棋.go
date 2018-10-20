package main

import (
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdk"
	"os"
	"github.com/mattn/go-gtk/glib"
	"strconv"
)

type ChessWidget struct {
	Win        *gtk.Window //窗体
	BlackImg   *gtk.Image  //黑棋图片
	WhiteImg   *gtk.Image  //白棋图片
	MinBtn     *gtk.Button //最小化按钮
	CloseBtn   *gtk.Button //关闭按钮
	BlackScore *gtk.Label  //黑棋分数
	WhiteScore *gtk.Label  //白棋分数
	TimerLab   *gtk.Label  //计时器
	LabTimerId int         //闪烁计时器id
	NumTimerId int         //时间计时器id

	w, h           int       //窗体大小
	startx, starty int       //棋盘起始坐标
	endx, endy     int       //棋盘结束坐标
	gridw, gridh   int       //棋盘格子宽高
	board          [8][8]int //存储棋盘棋子
	curRole        int       //执棋角色
	timerNum       int       //计时器时间
}

//枚举
const (
	Empty = iota //0
	Black        //1
	White        //2

)

//棋盘初始化方法
func (obj *ChessWidget) InitChess() {
	//初始化数据
	obj.startx = 210
	obj.starty = 70
	obj.endx = 620
	obj.endy = 480
	//棋盘格子大小
	obj.gridw = 52
	obj.gridh = 51

	//初始化棋盘棋子
	obj.board[3][3] = Black
	obj.board[4][4] = Black

	obj.board[3][4] = White
	obj.board[4][3] = White

	//黑方执棋
	obj.curRole = Black

}

//为棋盘对象绑定方法
func (obj *ChessWidget) CreateWidGet() {

	builder := gtk.NewBuilder()
	builder.AddFromFile("src/UI.glade")

	//加载窗体信息
	obj.Win = gtk.WindowFromObject(builder.GetObject("Win"))
	//设置窗体大小不可变
	obj.Win.SetResizable(false)
	//设置窗体无边框
	obj.Win.SetDecorated(false)
	//获取窗体大小
	obj.Win.GetSize(&obj.w, &obj.h)
	//设置标题
	obj.Win.SetTitle("黑白棋")

	//绘图事件
	obj.Win.SetAppPaintable(true)
	obj.Win.Connect("expose-event", DrawWindowImagefromFile, obj)

	//允许鼠标事件
	obj.Win.SetEvents(int(gdk.BUTTON_PRESS_MASK))
	obj.Win.Connect("button-press-event", MouseClickEvent, obj)

	//最小化按钮
	obj.MinBtn = gtk.ButtonFromObject(builder.GetObject("BtnMin"))
	//设置按钮图片
	DrawButtonImageFromFile(obj.MinBtn, "src/image/min.png")
	//设置按钮点击事件
	obj.MinBtn.Connect("clicked", BtnMinClick, obj)

	//关闭按钮
	obj.CloseBtn = gtk.ButtonFromObject(builder.GetObject("BtnClose"))
	//设置按钮图片
	DrawButtonImageFromFile(obj.CloseBtn, "src/image/close.png")
	//设置按钮点击事件
	obj.CloseBtn.Connect("clicked", BtnCloseClick, obj)

	//设置图片
	obj.BlackImg = gtk.ImageFromObject(builder.GetObject("Img1"))
	DrawImageFromFile(obj.BlackImg, "src/image/black.png")

	//设置图片
	obj.WhiteImg = gtk.ImageFromObject(builder.GetObject("Img2"))
	DrawImageFromFile(obj.WhiteImg, "src/image/white.png")

	//设置标签
	obj.TimerLab = gtk.LabelFromObject(builder.GetObject("Timer"))
	obj.TimerLab.ModifyFontSize(16)

	obj.BlackScore = gtk.LabelFromObject(builder.GetObject("BlackScore"))
	obj.BlackScore.ModifyFontSize(40)

	obj.WhiteScore = gtk.LabelFromObject(builder.GetObject("WhiteScore"))
	obj.WhiteScore.ModifyFontSize(40)

}

//游戏规则
//x y 当前位置下标  role当前角色 eatChess 是否吃子 返回值 吃子个数
func (obj *ChessWidget) ChessRule(x int, y int, role int, eatChess bool) (eatNum int) {

	// 棋盘的八个方向
	dir := [8][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	tempX, tempY := x, y // 临时保存棋盘数组坐标位置

	if obj.board[tempX][tempY] != Empty { // 如果此方格内已有棋子，返回
		return 0
	}

	// 棋盘的8个方向
	for i := 0; i < 8; i++ {
		tempX += dir[i][0]
		tempY += dir[i][1] // 准备判断相邻棋子

		// 如果没有出界，且相邻棋子是对方棋子，才有吃子的可能．
		if (tempX < 8 && tempX >= 0 && tempY < 8 && tempY >= 0) && (obj.board[tempX][tempY] != role) && (obj.board[tempX][tempY] != Empty) {
			tempX += dir[i][0]
			tempY += dir[i][1] // 继续判断下一个，向前走一步
			for tempX < 8 && tempX >= 0 && tempY < 8 && tempY >= 0 {
				if obj.board[tempX][tempY] == Empty { // 遇到空位跳出
					break
				}

				if obj.board[tempX][tempY] == role { // 找到自己的棋子，代表可以吃子
					if eatChess == true { // 确定吃子
						obj.board[x][y] = role // 开始点标志为自己的棋子
						tempX -= dir[i][0]
						tempY -= dir[i][1] // 后退一步
						for (tempX != x) || (tempY != y) {
							// 只要没有回到开始的位置就执行
							obj.board[tempX][tempY] = role // 标志为自己的棋子
							tempX -= dir[i][0]
							tempY -= dir[i][1] // 继续后退一步
							eatNum++           // 累计
						}
					} else { //不吃子，只是判断这个位置能不能吃子
						tempX -= dir[i][0]
						tempY -= dir[i][1] // 后退一步
						for (tempX != x) || (tempY != y) { // 只计算可以吃子的个数
							tempX -= dir[i][0]
							tempY -= dir[i][1] // 继续后退一步
							eatNum++
						}
					}

					break // 跳出循环
				} // 没有找到自己的棋子，就向前走一步

				tempX += dir[i][0]
				tempY += dir[i][1] // 向前走一步
			}
		} // 如果这个方向不能吃子，就换一个方向
		tempX, tempY = x, y
	}

	return // 返回能吃子的个数
}

func ShowTip(obj *ChessWidget) {
	if obj.curRole == Black {
		//隐藏白色棋子
		obj.WhiteImg.Hide()
		if obj.BlackImg.GetVisible() == true {
			obj.BlackImg.Hide()
		} else {
			obj.BlackImg.Show()
		}
	} else {
		//隐藏黑色棋子
		obj.BlackImg.Hide()
		if obj.WhiteImg.GetVisible() == true {
			obj.WhiteImg.Hide()
		} else {
			obj.WhiteImg.Show()
		}
	}
}

//控件显示
func (obj *ChessWidget) ShowChess() {
	//隐藏棋子图片
	obj.BlackImg.Hide()
	obj.WhiteImg.Hide()

	//闪烁计时器
	obj.LabTimerId = glib.TimeoutAdd(500, func() bool {
		ShowTip(obj)
		return true
	})

	//时间控制计时器
	obj.timerNum = 20
	obj.NumTimerId = glib.TimeoutAdd(1000, func() bool {
		obj.timerNum--
		//显示计时器时间
		obj.TimerLab.SetText(strconv.Itoa(obj.timerNum))
		if obj.timerNum == 0 {
			obj.ChangeRole()
		}
		return true
	})

}
func (obj *ChessWidget) ChangeRole() {
	//切换角色 重置时间
	obj.timerNum = 20
	//显示计时器时间
	obj.TimerLab.SetText(strconv.Itoa(obj.timerNum))
	obj.Win.QueueDraw()
	if obj.curRole == Black {
		obj.curRole = White
	} else {
		obj.curRole = Black
	}

	//棋子个数
	BlackCount := 0
	WhiteCount := 0
	//统计棋子个数
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if obj.board[i][j] == Black {
				BlackCount++
			} else if obj.board[i][j] == White {
				WhiteCount++
			}
		}
	}
	obj.BlackScore.SetText(strconv.Itoa(BlackCount))
	obj.WhiteScore.SetText(strconv.Itoa(WhiteCount))

}
func main() {

	gtk.Init(&os.Args)

	var obj ChessWidget
	//初始化方法
	obj.InitChess()
	//创建棋盘
	obj.CreateWidGet()
	//显示棋盘控件
	obj.ShowChess()

	//显示窗体信息
	obj.Win.ShowAll()
	gtk.Main()
}
