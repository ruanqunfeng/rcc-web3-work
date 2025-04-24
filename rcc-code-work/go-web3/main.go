package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	rpcURL        = "wss://sepolia.gateway.tenderly.co"          //"https://ethereum-sepolia-rpc.publicnode.com"                      // 替换为你的节点URL
	privateKeyHex = ""                                           // 替换为你的私钥
	contractAddr  = common.HexToAddress("0xCONTRACT_ADDRESS")    // 合约地址
	ethAddress    = "0x4aFc74781822F4B1eCa779312458b7EE75D7B69a" // 替换为你的以太坊地址
)

func main() {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 创建以太坊地址
	// createAddress()

	// 查询余额
	// queryBalance(client, ethAddress)

	// 示例调用各个功能
	// 1. 发起转账交易
	// sendTransaction(client)

	// 2. 查询合约数据
	// queryContractData(client)

	// 3. 查询区块数据
	queryBlockByNumber(client, 1234567)
	// queryBlockByHash(client, common.HexToHash("0x52612df6e27a87b5e305c2d34058b933ebfc0e33521822a4468f2b5a48364fe0"))

	// 4. 查询交易
	// queryTransaction(client, common.HexToHash("0x3d90668565b1ffb7a7c0a0c78a3ceafe1e6351cbc0cc1de1b974fb937c43890f"))

	// 5. 查询收据
	// queryReceipt(client, common.HexToHash("0x..."))

	// 6. 查询合约事件
	// queryContractEvents(client)

	// 7. 订阅合约事件
	// subscribeContractEvents(client)

	// 8. 订阅链事件
	// subscribeNewBlocks(client)
}

func createAddress() {
	// 1. 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 2. 获取私钥字节（32字节）
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// 3. 导出公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// 4. 获取公钥字节（64字节，未压缩格式）
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// 5. 生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 格式化输出
	fmt.Println("Private Key (hex):", hex.EncodeToString(privateKeyBytes))
	fmt.Println("Public Key (hex):", hex.EncodeToString(publicKeyBytes))
	fmt.Println("Ethereum Address:", address.Hex())

	// 可选：使用更安全的格式显示私钥
	fmt.Println("\n--- Secure Formats ---")
	fmt.Println("Private Key (hex with 0x):", hexutil.Encode(privateKeyBytes))
	fmt.Println("Public Key (hex with 0x):", hexutil.Encode(publicKeyBytes))
}

// 查询余额
func queryBalance(client *ethclient.Client, addressStr string) {
	// 要查询的地址（示例地址，可替换）
	address := common.HexToAddress(addressStr)

	// 获取最新区块的余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("查询余额失败: ", err)
	}

	// 将余额从 wei 转换为 ether
	ethBalance := new(big.Float)
	ethBalance.SetString(balance.String())
	ethBalance = ethBalance.Quo(ethBalance, big.NewFloat(1e18))

	// 输出结果
	fmt.Printf("地址: %s\n", address.Hex())
	fmt.Printf("余额 (wei): %s\n", balance.String())
	fmt.Printf("余额 (ETH): %s\n", ethBalance.String())
}

// 1. 发起交易（转账示例）
func sendTransaction(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(5000000000000000) // 0.005 ETH
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	to := common.HexToAddress("0xEEf0285Ce273d4460B23b8E37889ad81beC80cee")
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}

// 2. 查询合约数据（需要ABI）
// func queryContractData(client *ethclient.Client) {
// 	// 示例：查询ERC20余额
// 	contractABI, err := abi.JSON(strings.NewReader(erc20ABI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	data, err := contractABI.Pack("balanceOf", common.HexToAddress("0x..."))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg := ethereum.CallMsg{
// 		To:   &contractAddr,
// 		Data: data,
// 	}

// 	result, err := client.CallContract(context.Background(), msg, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var balance *big.Int
// 	err = contractABI.UnpackIntoInterface(&balance, "balanceOf", result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Balance:", balance)
// }

// 3. 查询区块数据（通过区块高度）
func queryBlockByNumber(client *ethclient.Client, blockNumber int64) {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Block #%d: %s\n", block.Number().Int64(), block.Hash().Hex())
}

// 3. 查询区块数据（通过区块哈希）
func queryBlockByHash(client *ethclient.Client, hash common.Hash) {
	block, err := client.BlockByHash(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Block #%d: %s\n", block.Number().Int64(), block.Hash().Hex())
}

// 转换wei为ETH
func weiToEth(wei *big.Int) string {
	fbalance := new(big.Float)
	fbalance.SetString(wei.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	return ethValue.Text('f', 18) // 保留18位小数
}

// 交易状态显示
func statusString(isPending bool) string {
	if isPending {
		return "待确认"
	}
	return "已确认"
}

// 4. 查询交易（通过交易哈希）
func queryTransaction(client *ethclient.Client, txHash common.Hash) {
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	singer := types.NewEIP155Signer(tx.ChainId())
	from, err := types.Sender(singer, tx)
	if err != nil {
		log.Fatal("获取发送方失败:", err)
	}

	var toAddress common.Address
	if tx.To() != nil {
		toAddress = *tx.To()
	} else {
		toAddress = common.HexToAddress("0x") // 合约创建交易
	}

	// 获取交易的 gasPrice 和 gasLimit
	gasPrice := tx.GasPrice()
	gasLimit := tx.Gas()

	// 获取交易的收据
	receipt := queryReceipt(client, txHash)

	// 实际使用的 gas
	actualGasUsed := receipt.GasUsed

	// 将 uint64 转换为 *big.Int
	gasLimitBigInt := new(big.Int).SetUint64(gasLimit)
	actualGasUsedBigInt := new(big.Int).SetUint64(actualGasUsed)

	var fee *big.Int

	// 计算手续费
	if receipt.Status == types.ReceiptStatusFailed {
		// 如果交易失败，矿工收取 gasLimit * gasPrice
		fee = new(big.Int).Mul(gasLimitBigInt, gasPrice)
		fmt.Printf("Transaction failed. Fee: %s wei\n", fee.String())
	} else {
		// 如果交易成功，矿工收取实际使用的 gas * gasPrice
		fee = new(big.Int).Mul(actualGasUsedBigInt, gasPrice)
		fmt.Printf("Transaction succeeded. Fee: %s wei\n", fee.String())
	}

	// 输出结果
	fmt.Printf("交易状态: %s\n", statusString(isPending))
	fmt.Printf("发送方地址: %s\n", from.Hex())
	fmt.Printf("接收方地址: %s\n", toAddress.Hex())
	fmt.Printf("金额 (ETH): %s\n", weiToEth(tx.Value()))
	fmt.Printf("手续费 (ETH): %s\n", weiToEth(fee))
}

// 5. 查询收据（通过交易哈希）
func queryReceipt(client *ethclient.Client, txHash common.Hash) *types.Receipt {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Receipt Status: %d\n", receipt.Status)
	fmt.Printf("Gas Used: %d\n", receipt.GasUsed)

	return receipt
}

// 6. 查询合约事件
// func queryContractEvents(client *ethclient.Client) {
// 	contractABI, err := abi.JSON(strings.NewReader(erc20ABI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 创建查询过滤器
// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(1000000),
// 		ToBlock:   big.NewInt(1001000),
// 		Addresses: []common.Address{contractAddr},
// 		Topics: [][]common.Hash{{
// 			contractABI.Events["Transfer"].ID, // 事件签名
// 		}},
// 	}

// 	logs, err := client.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, log := range logs {
// 		fmt.Printf("Log: %+v\n", log)
// 	}
// }

// 7. 订阅合约事件
// func subscribeContractEvents(client *ethclient.Client) {
// 	contractABI, err := abi.JSON(strings.NewReader(erc20ABI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{contractAddr},
// 	}

// 	logsCh := make(chan types.Log)
// 	sub, err := client.SubscribeFilterLogs(context.Background(), query, logsCh)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer sub.Unsubscribe()

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case log := <-logsCh:
// 			if log.Topics[0] == contractABI.Events["Transfer"].ID {
// 				fmt.Println("New Transfer event detected")
// 			}
// 		}
// 	}
// }

// 8. 订阅链事件（新区块）
func subscribeNewBlocks(client *ethclient.Client) {
	headersCh := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headersCh)
	if err != nil {
		log.Fatal(err)
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headersCh:
			fmt.Printf("New block: %d (%s)\n",
				header.Number.Int64(),
				header.Hash().Hex())
		}
	}
}
