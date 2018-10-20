package main

import (
	"time"
	"bytes"
	"encoding/binary"
	"log"
	"crypto/sha256"
	"encoding/gob"
)

//1,定义块结构
type Block struct {
	//====区块头=====
	//比特币网络的版本号
	Version uint64
	//前区块的哈希值，用于连接链条
	PrevHash []byte
	//梅克尔根，用于快速校验区块，校验区块的完整性。
	MerkelRoot []byte
	//时间戳，表示区块创建的时间
	TimeStamp uint64
	//难度值，调整挖矿难度
	Difficulty uint64
	//随机值，挖矿需要求的数字
	Nonce uint64
	//当前区块的哈希，注意，这是为了方便写代码而加入的。比特币不在区块中存储当前区块的哈希值
	Hash []byte

	//======区块体=======
	Transactions []*Transaction
}

//2,创建区块
func NewBlock(txs []*Transaction, prevBlockHash []byte) Block {
	block := Block{
		Version:      00,
		PrevHash:     prevBlockHash,
		MerkelRoot:   []byte{},
		TimeStamp:    uint64(time.Now().UnixNano()),
		Difficulty:   0,
		Nonce:        0,
		Transactions: txs,
	}
	//block.setHash()
	pow := NewProofOfWork(block) //准备数据
	Hash, Nonce := pow.Run()     //挖矿匹配
	block.Hash = Hash
	block.Nonce = Nonce
	return block //返回矿信息

}

//3，生成哈希

func (block *Block) setHash() {
	temp := [][]byte{
		Uint64ToByte(block.Version),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		Uint64ToByte(block.TimeStamp),
		//block.Transaction,
		block.Hash,
		block.MerkelRoot,
		block.PrevHash,
	}
	//拼接所有信息然后hash加密
	blockInfo := bytes.Join(temp, []byte{})
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//Uint64To[]byte
func Uint64ToByte(param uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, param)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//Serialize序列化
func (block *Block) Serialize() []byte {
	//buf
	buf := bytes.Buffer{}
	//创建一个编码器
	encoder := gob.NewEncoder(&buf)
	//编码
	err := encoder.Encode(block)
	if err != nil {
		log.Panic("序列化编码错误:", err)
	}
	return buf.Bytes()
}

func (block *Block) DeSerialize(data []byte) *Block {
	var buf bytes.Buffer
	//将data写入buf
	_, err := buf.Write(data)
	if err != nil {
		log.Panic("反序列化buf写入错误", err)
	}
	//创建解码器
	decoder := gob.NewDecoder(&buf)
	//解码
	err = decoder.Decode(&block)
	if err != nil {
		log.Panic("反序列化解码错误", err)
	}
	return block
}

/*
这个函数是为了生成Merkel Tree Root哈希值，正常的生成过程是使用所有交易的
哈希值生成一个平衡二叉树，此处，为了简化代码，
我们目前直接将区块中交易的哈希值进行拼接后进行哈希操作即可。*/

func (block *Block) HashTransactions() []byte {
	var temp [][]byte
	for _, tx := range block.Transactions {
		//交易的ID就是交易的哈希值，我们在Transaction里面提供了方法。
		temp = append(temp, tx.TXID)
	}
	data := bytes.Join(temp, []byte{})
	hash := sha256.Sum256(data)

	return hash[:]
}
