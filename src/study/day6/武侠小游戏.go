package main

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

//自定义角色名称{角色结构体{姓名，性别，随机位置，血量，内力，所选技能，状态}}
//自选技能{技能结构体{技能名称}}
//技能结构体{伤害,招式名称，耗蓝，攻击范围}
//擂台{长，宽范围，打出擂台之后即输}
//对话框{随机}
//按键出招，根基随机范围，技能攻击范围判断是否躲过，耗蓝躲避，根据血量判断输赢

//角色
type role struct {
	name     string `json:"userName"`
	gender   string `json:"userGender"`
	position int    `json:"userCurrentPosition"`
	life     int    `json:"userCurrentlife"`
	skill    int    `json:"userChoosedSkill"`
	warning  int    `json:"userCurrentLifeStuta"`
}

//技能
type skills struct {
	skillName string `json:"skill_name"`
}

//招数
type action struct {
	hurt        int    `json:"actionHurt"`
	actionName  string `json:"actionName"`
	useBlue     int    `json:"use_blue"`
	attankRange int    `json:"attank_range"`
}

//擂台
type arena struct {
	width  int
	height int
}

//随机对话框
type talkInfo struct {
	talkCotent string
}

func main() {
	//初始化操作键盘的包
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlW:
				fmt.Println("ddd")
			}
		default:
			break
		}
	}

}
