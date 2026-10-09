package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func main() {

	client, err := getClient()
	if err != nil {
		log.Fatal("连接 Sepolia 失败：", err)
	}
	//printBlockInfo()

	txHash, error := sendTransaction(client)
	if error != nil {
		log.Fatal("发送交易失败：", error)
	}
	for {
		fmt.Println("等待 1 秒后查询交易状态...")
		time.Sleep(1 * time.Second)
		queryTransaction(client, txHash)
	}
}

func getClient() (*ethclient.Client, error) {
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	if rpcURL == "" {
		log.Fatal("请先设置 SEPOLIA_RPC_URL")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("连接 Sepolia 失败：", err)
	}
	return client, err
}

// 打印区块信息
func printBlockInfo() {
	if len(os.Args) != 2 {
		log.Fatal("用法：go run . <区块号>")
	}

	number, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal("区块号必须是非负整数：", err)
	}

	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	if rpcURL == "" {
		log.Fatal("请先设置 SEPOLIA_RPC_URL")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("连接 Sepolia 失败：", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(number))
	if err != nil {
		log.Fatal("查询区块失败：", err)
	}

	fmt.Println("区块号：", block.Number())
	fmt.Println("区块哈希：", block.Hash().Hex())
	fmt.Println("时间戳：", block.Time())
	fmt.Println("UTC 时间：", time.Unix(int64(block.Time()), 0).UTC())
	fmt.Println("交易数量：", len(block.Transactions()))
}

func sendTransaction(cl *ethclient.Client) (string, error) {
	SEPOLIA_PRIVATE_KEY := os.Getenv("SEPOLIA_PRIVATE_KEY")
	ADDR := "0xEeBD193Cb96D61E93D22FdAeD12313E2BeC8ef6c"
	var (
		sk = crypto.ToECDSAUnsafe(common.FromHex(SEPOLIA_PRIVATE_KEY))
		to = common.HexToAddress("0xE7018da7BD70fC8696524257550585f4166E2332")
		//0.01ETH
		value    = new(big.Int).Div(big.NewInt(params.Ether), big.NewInt(100))
		sender   = common.HexToAddress(ADDR)
		gasLimit = uint64(35000)
	)
	// Retrieve the chainid (needed for signer)
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return "", err
	}
	// Retrieve the pending nonce
	nonce, err := cl.PendingNonceAt(context.Background(), sender)
	if err != nil {
		return "", err
	}
	// Get suggested gas price
	tipCap, _ := cl.SuggestGasTipCap(context.Background())
	feeCap, _ := cl.SuggestGasPrice(context.Background())
	// Create a new transaction
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainid,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     value,
			Data:      nil,
		})
	// Sign the transaction using our keys
	signedTx, _ := types.SignTx(tx, types.NewLondonSigner(chainid), sk)
	// Send the transaction to our node
	error := cl.SendTransaction(context.Background(), signedTx)

	// 输出交易信息
	fmt.Println("=== Transaction Sent ===")
	fmt.Printf("From       : %s\n", ADDR)
	fmt.Printf("To         : %s\n", to.Hex())
	fmt.Printf("Value      : %s ETH \n", value.String())
	fmt.Printf("Gas Limit  : %d\n", gasLimit)
	fmt.Printf("Gas Tip Cap: %s Wei\n", tipCap.String())
	fmt.Printf("Gas Fee Cap: %s Wei\n", feeCap.String())
	fmt.Printf("Nonce      : %d\n", nonce)
	fmt.Printf("Tx Hash    : %s\n", signedTx.Hash().Hex())
	fmt.Println("\nTransaction is pending. Use --tx flag to query status:")
	fmt.Printf("  go run main.go --tx %s\n", signedTx.Hash().Hex())

	return signedTx.Hash().Hex(), error
}

// 查询交易
func queryTransaction(client *ethclient.Client, txHashHex string) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	txHash := common.HexToHash(txHashHex)

	tx, isPending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Fatalf("failed to get transaction: %v", err)
	}

	fmt.Println("=== Transaction ===")
	printTxBasicInfo(tx, isPending)

	// 回执可能尚不可用（pending 交易）
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		log.Printf("failed to get receipt (maybe pending): %v", err)
		return
	}

	fmt.Println("=== Receipt ===")
	printReceiptInfo(receipt)
}

func printTxBasicInfo(tx *types.Transaction, isPending bool) {
	fmt.Printf("Hash        : %s\n", tx.Hash().Hex())
	fmt.Printf("Nonce       : %d\n", tx.Nonce())
	fmt.Printf("Gas         : %d\n", tx.Gas())
	fmt.Printf("Gas Price   : %s\n", tx.GasPrice().String())
	fmt.Printf("To          : %v\n", tx.To())
	fmt.Printf("Value (Wei) : %s\n", tx.Value().String())
	fmt.Printf("Data Len    : %d bytes\n", len(tx.Data()))
	fmt.Printf("Pending     : %v\n", isPending)
}

func printReceiptInfo(r *types.Receipt) {
	fmt.Printf("Status      : %d\n", r.Status)
	fmt.Printf("BlockNumber : %d\n", r.BlockNumber.Uint64())
	fmt.Printf("BlockHash   : %s\n", r.BlockHash.Hex())
	fmt.Printf("TxIndex     : %d\n", r.TransactionIndex)
	fmt.Printf("Gas Used    : %d\n", r.GasUsed)
	fmt.Printf("Logs        : %d\n", len(r.Logs))
	if len(r.Logs) > 0 {
		fmt.Printf("First Log Address : %s\n", r.Logs[0].Address.Hex())
	}
}
