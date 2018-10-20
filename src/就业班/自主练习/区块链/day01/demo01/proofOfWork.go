package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

//工作量证明（挖矿）
type proofOfWork struct {
	//区块数据
	block Block
	//目标值（暂时写死）
	target big.Int
}

func NewProofOfWork(block Block) *proofOfWork {
	pow := proofOfWork{
		block: block,
	}
	tarGetString := "0000100000000000000000000000000000000000000000000000000000000000"
	bigIntTemp := big.Int{}
	bigIntTemp.SetString(tarGetString, 16) //将tarGetString转成8进制并赋值bigIntTemp
	pow.target = bigIntTemp
	return &pow
}
//挖矿函数
func (pow *proofOfWork) Run() ([]byte, uint64) {
	var Nonce uint64
	var hash [32]byte
	for ; ; {
		hash = sha256.Sum256(pow.prepareData(Nonce))
		tTemp := big.Int{}
		tTemp.SetBytes(hash[:]) //将bytes
		//   X          Y    x<y时返回-1；x>y时返回+1；否则返回0。
		if tTemp.Cmp(&pow.target) == -1 {
			fmt.Printf("挖矿成功: %x, %d,%x \n", hash, Nonce, pow.block.PrevHash)
			break
		} else {
			//继续循环
			Nonce++
		}
	}
	return hash[:], Nonce
}

//准备数据类似setHash
func (pow *proofOfWork) prepareData(num uint64) []byte {
	block := pow.block
	//将所有信息拼接容易hash
	obj := [][]byte{
		Uint64ToByte(block.Version),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(num),
		Uint64ToByte(block.TimeStamp),
		//block.Data,
		block.MerkelRoot,
		block.PrevHash,
	}
	data := bytes.Join(obj, []byte{})
	return data
}

//校验函数
func (pow *proofOfWork) IsValid() bool {
	hash := sha256.Sum256(pow.prepareData(pow.block.Nonce))
	tTmp := big.Int{}
	tTmp.SetBytes(hash[:])
	if tTmp.Cmp(&pow.target) == -1 {
		return true
	}
	return false
}
