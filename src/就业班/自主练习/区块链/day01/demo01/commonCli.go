package main

import (
	"fmt"
	"github.com/base58"
	"bytes"
	"os"
	"time"
)

func (cli *CLI) addBlock(data string) {
	bc := GetBlockChainObj()
	//bc.AddBlock(data)
	bc.db.Close()
	fmt.Println("添加块成功", bc)
}

func (cli *CLI) PrinBlockChain() {
	bc := GetBlockChainObj()
	it := bc.NewIterator()
	for ; ; {
		//返回区块，左移
		block := it.Next()
		fmt.Printf("===========================\n\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		fmt.Printf("时间戳: %d\n", block.TimeStamp)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		//fmt.Printf("区块数据 :%s\n", block.Data)

		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		}
	}
}

//获取余额
//获取余额需要指定地址，通过遍历整个账本，从而找到这个地址可用的utxo，为此我们要做两件事：
//1, 校验地址的有效性
//   传递过来的地址有可能是无效的，无效的地址直接返回即可。
//2, 逆推出公钥哈希
//   并不是所有的地址都是本地生成的，有可能是别人的地址，所以我们需要逆推而不是打开钱包去查找。
//3, 遍历账本
//   调用FindUTXOs函数

func (cli *CLI) GetBalance(address string) {
	bc := GetBlockChainObj()
	defer bc.db.Close()
	//校验地址有效性
	CheckAddress(address)
	publicKeyHash := base58.Decode(address)
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-4] //因为是(version+publicKeyHash)+checkCode拼成的（4位）
	utxos := bc.FindUTXOs(publicKeyHash)
	var total float64
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("The balance of %s is : %f\n", address, total)
}

func IsValidAddress(address string) bool {
	//1. 逆向求出pubKeyHash数据
	//2. 得到version + 哈希1 部分，并做checksum运算
	//3. 与实际checksum比较，如果相同，则地址有效，反之则无效

	publicKeyHash := base58.Decode(address)
	if len(publicKeyHash) < 4 {
		return false
	}
	//之前加密我们用了vesion+publicKeyHash+checkNum拼接的byte
	payload := publicKeyHash[:len(publicKeyHash)-4]
	CheckCode := publicKeyHash[len(publicKeyHash)-4:]
	fmt.Printf("CheckCode : %x\n", CheckCode)
	targetCode := Checksum(payload) //返回payload两次加密后的后四位
	fmt.Printf("targetCode : %x\n", targetCode)
	return bytes.Compare(CheckCode, targetCode) == 0
}

func CheckAddress(address string) {
	if !IsValidAddress(address) {
		fmt.Println("钱包地址有误", address)
		os.Exit(1)
	}
	fmt.Println("有效的钱包地址")
}


//反向打印
func (cli *CLI) PrinBlockChainReverse() {
	bc := GetBlockChainObj()
	//创建迭代器
	it := bc.NewIterator()

	//调用迭代器，返回我们的每一个区块数据
	for {
		//返回区块，左移
		block := it.Next()

		fmt.Printf("===========================\n\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳: %s\n", timeFormat)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Signature)

		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		}
	}
}
