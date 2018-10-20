package main

import (
	"os"
	"fmt"
	"strconv"
)

type CLI struct {
}

//用法提示信息
const Usage = `
    send --from FROM --to TO --amount AMOUNT --miner MINER  "send money from FROM to TO"
    printChain  "print all blockchain data"
    addblockchain --address  ADDRESS  "create a blockchian"
    getbanlance --address ADDRESS "get banlance by address"
    createwallet "create a new key pair wallet and save into wallet.dat"
    listaddress "list all addresses in wallet.dat"
`

//接收并判断接收参数，执行相应的代码
func (cli *CLI) Run() {
	args := os.Args
	//判断命令行长度，不合规的直接不处理
	if len(args) < 2 {
		fmt.Println("参数不对001：", Usage)
		return
	}
	//分析命令
	cmd := args[1] //具体判断
	switch cmd {
	case "send":
		//添加区块
		if len(args) < 10 {
			fmt.Printf("添加区块参数使用不当，请检查")
			fmt.Printf(Usage)
			os.Exit(1)
		} else { //不然就是输入有误
			from := os.Args[3]
			to := os.Args[5]
			amount, _ := strconv.ParseFloat(os.Args[7], 64)
			miner := os.Args[9]
			cli.Send(from, to, amount, miner)
		}
	case "printchain":
		fmt.Printf("打印区块\n")
		cli.PrinBlockChain()
	case "printchainr":
		cli.PrinBlockChainReverse()
	case "createblockchain":
		if len(os.Args) > 3 && os.Args[2] == "--address" {
			address := args[3]
			if address == "" {
				fmt.Println("address should not be empty!")
				os.Exit(1)
			}
			cli.CreateBlockChain(address)
		} else {
			println(Usage)
		}
	case "getbalance":
		if len(os.Args) > 3 && os.Args[2] == "--address" {
			address := args[3]
			if address == "" {
				fmt.Println("getbanlance:address should not be empty")
				os.Exit(1)
			}
			cli.GetBalance(address)
		} else {
			fmt.Println(Usage)
		}
	case "createwallet":
		cli.CreateWallet()
	case "listaddress":
		cli.ListAddresses()

	default:
		fmt.Printf("无效的命令，请检查!\n")
		fmt.Printf(Usage)
	}
}

//创建blockChain
func (cli *CLI) CreateBlockChain(address string) {
	CheckAddress(address)
	bc := NewBlockChain(address)
	err := bc.db.Close()
	if err != nil {
		if dbExit() {
			err = os.Remove(blockChainDB)
			if err != nil {
				fmt.Println("warning!db not be removed")
			}
		}
		fmt.Println("create blockchain failed!")
	}
	fmt.Println("create blockchain successfully!")
}

//转账功能
func (cli *CLI) Send(from, to string, amount float64, miner string) {
	bc := GetBlockChainObj()
	tx := NewTransaction(from, to, amount, bc)
	defer bc.db.Close()

	//指定挖矿
	coinBase := NewCoinbaseTX(miner, "这是一个矿~矿~矿~矿~矿~矿~矿~矿~矿~")
	bc.AddBlock([]*Transaction{coinBase, tx})
	fmt.Println("转账成功")
}

//创建钱包
func (cli *CLI) CreateWallet() {
	ws := NewWallets()
	address := ws.CreateWallet()
	//保存地址
	ws.SaveToFile()
	fmt.Printf("address  : %s\n", address)
}

//获取所有钱包地址
func (cli *CLI) ListAddresses() {
	ws := NewWallets()
	addresses := ws.GetAllAddresses()
	for i, address := range addresses {
		fmt.Printf("address[%d] : %s\n", i, address)
	}
}



