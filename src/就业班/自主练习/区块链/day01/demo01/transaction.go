package main

import (
	"bytes"
	"encoding/gob"
	"crypto/sha256"
	"fmt"
	"os"
	"github.com/base58"
	"crypto/ecdsa"
	"log"
	"crypto/rand"
	"math/big"
	"crypto/elliptic"
)

const reward = 12.5 //当前挖矿奖励
//交易输入
//包含交易发起人可支付资金的来源
type TXInput struct {
	//引用output所在的交易ID
	TXID []byte
	//output索引值
	VoutIndex int64
	//解锁脚本
	//ScriptSig string
	//签名
	Signature []byte
	//公钥
	PublicKey []byte
}

//输出交易
//包含资金接收方的相关信息
type TXOutput struct {
	//接收的金额
	Value float64
	//锁定脚本
	//ScriptPublic string
	//公钥哈希
	PublicKeyHash []byte
}

//交易结构
type Transaction struct {
	//交易ID
	TXID []byte
	//输入交易----可能是多个
	TXInputs []TXInput
	//输出交易----可能是多个（跟输入数目均衡）
	TXOutputs []TXOutput
}

//设置交易哈希ID
func (tx *Transaction) SetTransactionHash() {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(tx)
	data := buf.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

//挖矿交易
//coinbase总是新区块的第一条交易，这条交易中只有一个输出，即对矿工的奖励，没有输入。
func NewCoinbaseTX(address string, data string) *Transaction {
	//挖矿交易的特点：
	//1. 只有一个input
	//2. 无需引用交易id
	//3. 无需引用index
	//矿工由于挖矿时无需指定签名，所以这个sig字段可以由矿工自由填写数据，一般是填写矿池的名字
	if data == "" {
		data = fmt.Sprintf("reward %s %f\n", address, reward)
	}
	input := TXInput{nil, -1, nil, []byte(data)}
	output := NewUTXOutput(reward, address)

	txTemp := Transaction{nil, []TXInput{input}, []TXOutput{*output}}
	txTemp.SetTransactionHash()
	return &txTemp
}

//解锁脚本
//存储在output中的其实是公钥的哈希，我们去锁定的时候也是输入地址，然后内部逆向算出哈希存储的
func (input *TXInput) CanUnlockUTXOWith(publicKeyHash []byte /*收款人的公钥哈希*/) bool {
	hash := HashPublicKey(input.PublicKey)
	return bytes.Compare(hash, publicKeyHash) == 0
}

//锁定脚本
func (output *TXOutput) CanBeUnlockedWith(publicKeyHash []byte /*收款人的公钥哈希*/) bool {
	return bytes.Compare(output.PublicKeyHash, publicKeyHash) == 0
}

//=========产生交易======
//from：付款人
//to：收款人
//amount：转账金额
//bc：区块链本身
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	validUTXOs := make(map[string][]int64)
	var total float64

	//第一步：打开钱包，根据创建人的address找到对应的钱包（银行卡）
	//返回当前的钱包容器，加载到内存，注意，不会创建新的秘钥对
	ws := NewWallets()
	//获取钱包地址  //获取到当前地址的钱包(注意要加*)
	wallet := *ws.WalletsMap[from]
	publicKeyHash := HashPublicKey(wallet.PublicKey)
	//第二步找到所需要的UTXO的集合
	validUTXOs /*本次支付所需要的utxo的集合*/ , total /*返回utxos所包含的金额*/ = bc.FindSuitableUTXOs(publicKeyHash, amount)

	//钱不够的状态
	if total < amount {
		fmt.Println("Money is not enough!!!")
		os.Exit(1)
	}

	//钱够了====此处相等
	//将返回的utxo转换成input
	var inputs []TXInput
	var outputs []TXOutput
	//一个交易要包含两个信息
	//交易id
	//utxo索引
	//解锁脚本
	for txid, indexs := range validUTXOs {
		for _, index := range indexs {
			input := TXInput{[]byte(txid), index, nil, wallet.PublicKey /*这里千万别错误的填哈希值*/ }
			//input集合
			inputs = append(inputs, input)
		}
	}

	//第四步：创建输出（付款，找零）
	//调用函数生成新的output
	//创建output接收金额，锁定脚本
	output := NewUTXOutput(amount, from)
	outputs = append(outputs, *output)

	//utoxs的总数大于要交易的数-----找零
	if total > amount {
		output := NewUTXOutput(total-amount, from)
		outputs = append(outputs, *output)
	}
	txTmp := Transaction{nil, inputs, outputs}
	txTmp.SetTransactionHash()

	//交易签名
	bc.SignTransaction(&txTmp, wallet.PrivateKey)
	return &txTmp
}

//锁定脚本
//需要逆向求出来pubKeyHash
func (output *TXOutput) Lock(address string /*付款人地址*/) {
	//先解码base58
	bts := base58.Decode(address)
	//这个byte包含版本，公钥哈希，和4位校验码
	publicKeyHash := bts[1 : len(bts)-4]
	//使用公钥哈希锁定该output
	output.PublicKeyHash = publicKeyHash
}

//单独创建输出交易
func NewUTXOutput(value float64, address string) *TXOutput {
	txoutput := TXOutput{value, nil}
	txoutput.Lock(address) //锁定
	return &txoutput
}

//======签名=====

//交易创建完成之后，在写入区块之前，我们要对其进行签名，这样对端才能够校验，从而保证系统的安全性。

//======签名材料=====
//想要签名的数据
//私钥
func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	//挖矿属于系统奖励---不需要签名
	if tx.IsCoinBase() {
		return
	}
	//确保交易有效
	for _, input := range tx.TXInputs {
		//判断交易是否为空
		if prevTXs[string(input.TXID)].TXID == nil {
			log.Panic("Previous txs are not valid!")
		}
	}

	//获取要签名的信息
	//我们对把当前的交易签名，然后存放到Signature中。
	//将当前的交易复制一份，然后做签名处理
	txCopy := tx.TrimmedCopy()

	//- 欲使用utxo中的pubKeyHash（这描述了付款人）
	//- 新生成utxo中的pubKeyHash（这描述了收款人） (output中)
	//- 转账金额 (output中)

	//注意，遍历的是txCopy，而不是tx本身
	for index, input := range txCopy.TXInputs {
		prevTX := prevTXs[string(input.TXID)]
		//使用input的PublicKey字段暂时存储一下这个想要解锁的utxo的公钥哈希
		//这里有个坑！！！
		//我一直使用input.PublicKey来赋值，但是其实range会复制一个新的变量，而不会修改原有的txCopy
		//这里一定要小心！！！
		//input.PublicKey = prevTX.TXOutputs[input.VoutIndex].PublicKeyHash
		txCopy.TXInputs[index].PublicKey = prevTX.TXOutputs[input.VoutIndex].PublicKeyHash
		//对哈希进行签名
		txCopy.SetTransactionHash() //交易的哈希值存放的TXID里面
		//赋值nil为了防止污染其他交易，生成要签名的数据就销毁
		txCopy.TXInputs[index].PublicKey = nil
		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, txCopy.TXID)
		if err != nil {
			fmt.Println("ecdsa签名失败")
			log.Panic(err)
		}
		//拼接r,s放在签名变量中
		signature := append(r.Bytes(), s.Bytes()...)
		tx.TXInputs[index].Signature = signature //输入签名成功
	}
}
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput
	//找到所有输入信息
	for _, input := range tx.TXInputs {
		inputs = append(inputs, TXInput{input.TXID, input.VoutIndex, nil, nil})
	}
	//找到所有输出信息
	for _, output := range tx.TXOutputs {
		outputs = append(outputs, TXOutput{output.Value, output.PublicKeyHash})
	}
	txCopy := Transaction{tx.TXID, inputs, outputs}
	return txCopy
}

//在挖矿时校验
//当一笔交易发送到对端时，接收方在打包到自己的区块前，需要先对交易进行校验，从而保证
//交易确实是由私钥的持有者发起的
//持有者花费的确实是自己的钱
//- 分析
//校验函数(Verify)依然在Transaction结构中实现，需要三个数据：
//想要签名的数据
//数字签名
//公钥
//代码如下：
//校验函数原型如下，由于交易中已经存储了数字签名和公钥，所以只需要将引用的交易传递进来即可（为了获取引用输出的公钥哈希）
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.IsCoinBase() {
		return true
	}
	for _, input := range tx.TXInputs {
		if prevTXs[string(input.TXID)].TXID == nil {
			log.Panic("前区块交易信息有误")
		}
	}

	//以下是有效交易数据的的校验
	txCopy := tx.TrimmedCopy()
	//获取签名的原始数据(一个哈希值)，与签名时步骤一样
	//签名时对每一个引用的input都签名，校验也一定是对每一个input都校验
	//这里要注意for循环的对象，应该是原始的inputs结构，而不应该是copy过来的那个
	//与Sign时遍历的不同
	for index, input := range tx.TXInputs {
		prevTX := prevTXs[string(input.TXID)]
		//这个可以为了确保为空再设置一次nil
		txCopy.TXInputs[index].Signature = nil
		txCopy.TXInputs[index].PublicKey = prevTX.TXOutputs[input.VoutIndex].PublicKeyHash
		txCopy.SetTransactionHash()
		//这个也可以再次置空，不过不需要，因为每一个index对应的input只使用一次(这是错误的理解)
		//之所尚未出错，因为我们现在都只引用一个交易就完成了，没校验多个input
		//一定要设置，否则会影响其他交易的校验
		txCopy.TXInputs[index].PublicKey = nil
		//签名处理
		sigLen := len(input.Signature)
		r := big.Int{}
		s := big.Int{}

		r.SetBytes(input.Signature[:(sigLen/2)])
		s.SetBytes(input.Signature[(sigLen/2):])
		//处理公钥
		x := big.Int{}
		y := big.Int{}
		keyLen := len(input.PublicKey)
		x.SetBytes(input.PublicKey[:(keyLen/2)])
		y.SetBytes(input.PublicKey[(keyLen/2):])
		cutve := elliptic.P256()
		rawPubKey := ecdsa.PublicKey{cutve, &x, &y}
		if !ecdsa.Verify(&rawPubKey, txCopy.TXID, &r, &s) {
			return false
		}
	}
	return true
}
