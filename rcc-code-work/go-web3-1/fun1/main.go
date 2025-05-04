package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 1. 创建随机钱包
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("生成私钥失败: ", err)
	}

	// 获取钱包地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换到ECDSA公钥")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 2. 准备消息
	message := "Hello"

	// 3. 计算带前缀的消息哈希
	messageBytes := []byte(message)
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(messageBytes))
	prefixedMessage := []byte(prefix + message)
	hash := crypto.Keccak256Hash(prefixedMessage)

	// 4. 签名消息
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal("签名失败: ", err)
	}

	// 5. 分解签名
	if len(signature) != 65 {
		log.Fatal("签名长度不正确")
	}

	r := hexutil.Encode(signature[:32])
	s := hexutil.Encode(signature[32:64])
	v := uint8(signature[64]) + 27 // 转换为以太坊的V值

	// 输出结果
	fmt.Println("Wallet Address:", address.Hex())
	fmt.Println("Message Hash:", hash.Hex())
	fmt.Printf("Signature Values:\n  v: %d\n  r: %s\n  s: %s\n", v, r, s)
}
