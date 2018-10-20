package main

import (
	"github.com/bolt-master"
	"fmt"
	"os"
	"log"
	"crypto/ecdsa"
	"bytes"
	"errors"
)

//中本聪在创世块中保存的信息
const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type BlockChain struct {
	//操作数据库句柄
	db *bolt.DB
	//最后一个哈希
	tail []byte
}

const blockChainDB = "blockChain.db"
const blockBucket = "blockBucket"

//定义区块链
func NewBlockChain(address string) *BlockChain {

	var lastHash []byte
	db, err := bolt.Open(blockChainDB, 0600, nil)
	if err != nil {
		fmt.Println("bolt.Open:", err)
		os.Exit(1)
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		//空的创建
		if bucket == nil {
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				fmt.Println("CreateBucket", err)
				os.Exit(1)
			}

			//发生交易
			coinbaseTx := NewCoinbaseTX(address, genesisInfo)
			//区块链没有数据，创建数据，添加创世区块
			genesisBlock := NewBlock([]*Transaction{coinbaseTx}, []byte{})
			//添加数据
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("lastHash"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("lastHash"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

//获取一个区块链实例
func GetBlockChainObj() *BlockChain {
	if !dbExit() {
		fmt.Println("blockchain not exist, pls create first!")
		os.Exit(1)
	}
	var lashHash []byte
	db, err := bolt.Open(blockChainDB, 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败", err)
	}
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket should not be nill")
		}
		lashHash = bucket.Get([]byte("lastHash"))
		return nil
	})
	return &BlockChain{db, lashHash}
}

//检查数据库是否存在
func dbExit() bool {
	//Stat返回一个描述name指定的文件对象的FileInfo。
	//如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，
	//本函数会尝试跳转该链接。如果出错，返回的错误值为*PathError类型。
	if _, err := os.Stat(blockChainDB); os.IsNotExist(err) {
		return false
	}
	return true
}

//添加区块链
//在挖矿时校验
//当一笔交易发送到对端时，接收方在打包到自己的区块前，需要先对交易进行校验，从而保证
//交易确实是由私钥的持有者发起的
//持有者花费的确实是自己的钱
func (bc *BlockChain) AddBlock(data []*Transaction) {
	//获取最后一个哈希
	lastBlockHash := bc.tail
	db := bc.db
	//创建新的区块链
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("不该为空，请检查") //因为执行到此处区块链的库已经存在，直接使用即可
		}
		//添加新的块
		block := NewBlock(data, lastBlockHash)
		for _, tx := range data {
			if !bc.VerifyTransaction(tx) {
				log.Panic("挖矿前数据校验失败")
			}
		}
		//将新的块添加到链中
		bucket.Put(block.Hash, block.Serialize())
		//更新数据库中的lastHash
		bucket.Put([]byte("lastHash"), block.Hash)
		//将内存中的tail更新
		bc.tail = block.Hash
		return nil
	})
}

//挖矿前数据校验
func (bc *BlockChain) VerifyTransaction(tx *Transaction) bool {
	//必须加上这句校验，否则在FindTranction中就会报错，因为coinbase所引用的input的id为nil
	if tx.IsCoinBase() {
		return true
	}
	prevTXs := make(map[string]Transaction)
	for _, input := range tx.TXInputs {
		prevTX, err := bc.FindTransaction(input.TXID)
		if err != nil {
			log.Panic("挖矿前校验获取查找交易", err)
		}
		prevTXs[string(prevTX.TXID)] = prevTX //根据交易id找到交易
	}

	return tx.Verify(prevTXs)
}

func (bc *BlockChain) FindUTXOTransactions(publicKeyHash []byte) []Transaction {
	var transactions []Transaction
	//支付UTXOS的集合
	spentUTXOs := make(map[string][]int64)
	//创建链迭代器
	it := bc.NewIterator()
	//首先遍历链上的区块
	for {

		block := it.Next()
		//遍历区块中的交易信息
		for _, tx := range block.Transactions {
			//遍历所有output信息
		OUTPUTS:
			for currentIndex, output := range tx.TXOutputs {
				//在这里做一个过滤，将所有消耗过的outputs和当前的所即将添加output对比一下
				//如果相同，则跳过，否则添加
				//如果当前的交易id存在于我们已经表示的map，那么说明这个交易里面有消耗过的output
				if spentUTXOs[string(tx.TXID)] != nil {
					indexs := spentUTXOs[string(tx.TXID)]
					for _, index := range indexs {
						if int64(currentIndex) == index {
							continue OUTPUTS
						}
					}
				}
				if output.CanBeUnlockedWith(publicKeyHash) {
					transactions = append(transactions, *tx)
				}
			}
			//遍历所有input信息
			if !tx.IsCoinBase() {
				for _, input := range tx.TXInputs {
					if input.CanUnlockUTXOWith(publicKeyHash) {
						spentUTXOs[string(input.TXID)] = append(spentUTXOs[string(input.TXID)], input.VoutIndex)
					}
				}
			}
		}
		if len(block.PrevHash) == 0 {
			break
		}
	}
	return transactions
}

//判断是否是挖矿操作
func (tx *Transaction) IsCoinBase() bool {
	if len(tx.TXInputs) == 1 {
		if tx.TXInputs[0].TXID == nil && tx.TXInputs[0].VoutIndex == -1 {
			return true
		}
	}
	return false
}

//返回指定地址能够支配的utxo的集合
func (bc *BlockChain) FindUTXOs(publicKeyHash []byte) []TXOutput {

	txs := bc.FindUTXOTransactions(publicKeyHash)
	var utxos []TXOutput

	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			//找到能够解锁匹配的--也就是找到自己能够拿出钱的utxo
			if output.CanBeUnlockedWith(output.PublicKeyHash) {
				utxos = append(utxos, output)
			}
		}
	}
	return utxos
}

//遍历所有交易找到合适的utxo或者累加足够数目用于交易
func (bc *BlockChain) FindSuitableUTXOs(publicKeyHash []byte, amount float64) (map[string][]int64, float64) {
	//获取所有交易

	txs := bc.FindUTXOTransactions(publicKeyHash)
	validUTXOS := make(map[string][]int64)
	var total float64 = 0

CALCULATE:
	for _, currentTx := range txs {
		for currentIndex, output := range currentTx.TXOutputs {
			if output.CanBeUnlockedWith(output.PublicKeyHash) { //遍历当前我可以解锁的output
				if total < amount {
					total += output.Value
					//本次utxo消费的集合
					validUTXOS[string(currentTx.TXID)] = append(validUTXOS[string(currentTx.TXID)], int64(currentIndex))
				}
			} else {
				break CALCULATE
			}
		}
	}
	return validUTXOS, total
}

//签名
//在NewTransaction的最后，使用私钥对其进行签名。
//签名函数由Transaction提供，但是签名动作由blockChain来实现，因为我们要遍历账本，
//在SignTransaction内部再调用Sign函数。

func (bc *BlockChain) SignTransaction(tx *Transaction, privateKey ecdsa.PrivateKey) {
	//寻找所引用的交易数组
	//调用tx.Sign进行签名
	//遍历账本对所有交易进行签名
	prevTXs := make(map[string]Transaction)
	for _, input := range tx.TXInputs {
		prevTX, err := bc.FindTransaction(input.TXID)
		if err != nil {
			log.Panic(err)
		}
		//找到了要签的交易信息数据
		prevTXs[string(prevTX.TXID)] = prevTX

	}
	tx.Sign(privateKey, prevTXs) //完事直接签名
}

//通过ID返回交易结构
func (bc *BlockChain) FindTransaction(txid []byte) (Transaction, error) {
	it := bc.NewIterator()
	for ; ; {
		block := it.Next()
		for _, tx := range block.Transactions {
			if bytes.Compare(txid, tx.TXID) == 0 {
				return *tx, nil
			}
		}
		if len(block.PrevHash) == 0 {
			break
		}
	}
	return Transaction{}, errors.New("不存在此交易")
}
