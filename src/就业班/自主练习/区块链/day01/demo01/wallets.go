package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"io/ioutil"
	"crypto/elliptic"
	"os"
)

//定义一个容纳所有钱包的容器
type Wallets struct {
	//key:地址  value：钱包
	WalletsMap map[string]*Wallet
}

//对钱包的操作分为三步：
//
//从本地加载已有的钱包到内存
//添加新的钱包到内存
//将内存中的钱包保存到本地

//一、从本地加载已有的钱包到内存
func NewWallets() (*Wallets) {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)
	//加载本地数据
	ws.LoadFromFile()
	return &ws
}

//二、//添加新的钱包到内存
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.getAddress()
	ws.WalletsMap[address] = wallet
	return address
}

//三、将内存中的钱包保存到本地

func (ws *Wallets) SaveToFile() {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	gob.Register(elliptic.P256()) //<<-----加上这句，注册一个interface对象
	//注册原因：如果Encode/Decode类型是interface或者struct中某些字段是interface{}的时候，
	// 需要在gob中注册interface可能的所有实现或者可能类型，否则会上面的错误
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic("wallets gob encode err:", err)
	}
	walletFileName := "walletsCollect.dat" //存储钱包信息
	err = ioutil.WriteFile(walletFileName, buf.Bytes(), 0644)
	if err != nil {
		log.Panic("wallets ioutil.WriteFile err:", err)
	}
}

//本地加载钱包数据
//添加函数LoadFromFile，注意要使用地址传递func (ws *Wallets)
func (ws *Wallets) LoadFromFile() {
	walletFileName := "walletsCollect.dat"
	_, err := os.Stat(walletFileName)
	if os.IsNotExist(err) {
		return
	}
	content, err := ioutil.ReadFile(walletFileName)
	if err != nil {
		log.Panic("wallets:ioutil.ReadFile:", err)
	}
	var localWallet Wallets
	gob.Register(elliptic.P256())

	//解密
	decoder := gob.NewDecoder(bytes.NewReader(content))
	err = decoder.Decode(&localWallet)
	if err != nil {
		log.Panic("wallets:gob.NewDecoder:", err)
	}
	ws.WalletsMap = localWallet.WalletsMap
}

//获取所有地址
func (ws *Wallets) GetAllAddresses() []string {
	var addressContainer []string
	for address := range ws.WalletsMap {
		addressContainer = append(addressContainer, address)
	}
	return addressContainer
}
