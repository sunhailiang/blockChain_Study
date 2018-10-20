package main

import (
	"github.com/bolt-master"
	"log"
)

type BlockChainIterator struct {
	db                 *bolt.DB
	currentHashPointer []byte
}

//迭代器作用：遍历容器，将数据逐个返回，防止一次性加载到内存，所以一点一点读取。
//类比：for循环里面range
//我们的区块链迭代器图示,最初指向最后一个区块，返回区块，指针前移，直至第一个区块。

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		//最初指向区块链的最后一个区块，随着Next的调用，不断变化
		bc.tail,
	}
}

//next方法属于当前区块，执行这个方法，指针前移，知道指针指向创世块
func (it *BlockChainIterator) Next() *Block {
	var block *Block
	it.db.View(func(tx *bolt.Tx) error {
		buckct := tx.Bucket([]byte(blockBucket))
		if buckct == nil {
			log.Panic("迭代器遍历时bucket不应该是空的，请检查")
		}
		blockTemp := buckct.Get(it.currentHashPointer)
		//将取出来的数据反编译
		block = block.DeSerialize(blockTemp)
		//游标向左移动
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return block
}
