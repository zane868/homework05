package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"homework05/bindings"
)

const sepoliaChainID = 11155111
const donationContractHex = "0x061D631016a3850877074e2727640065f7eEd7a4"

func main() {
	if err := run(os.Args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}
		fmt.Fprintln(os.Stderr, "错误：", safeError(err))
		os.Exit(1)
	}
}

func getDonation(
	ctx context.Context,
	client *ethclient.Client,
	contractAddress common.Address,
	donor common.Address,
) (*big.Int, error) {
	contract := bindings.NewBeggingContract()
	data := contract.PackGetDonation(donor)
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil) // nil 表示查询最新区块
	if err != nil {
		return nil, err
	}
	return contract.UnpackGetDonation(result)
}

func (donationGetCommand) Execute(args []string) error {
	fs := flag.NewFlagSet("donation get", flag.ContinueOnError)
	donorText := fs.String("donor", "", "捐赠者地址")
	if err := parseFlags(fs, args); err != nil {
		return err
	}
	if !common.IsHexAddress(*donorText) {
		return errors.New("请提供有效的 -donor 地址")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := dialSepolia(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	contractAddress := common.HexToAddress(donationContractHex)
	donor := common.HexToAddress(*donorText)
	amount, err := getDonation(ctx, client, contractAddress, donor)
	if err != nil {
		return fmt.Errorf("查询捐赠额失败：%w", err)
	}
	fmt.Println("合约地址：", contractAddress.Hex())
	fmt.Println("捐赠者：", donor.Hex())
	fmt.Println("累计捐赠：", formatWei(amount), "ETH")
	fmt.Println("累计捐赠（wei）：", amount)
	return nil
}

func (donationSendCommand) Execute(args []string) error {
	fs := flag.NewFlagSet("donation send", flag.ContinueOnError)
	amountText := fs.String("amount", "", "捐赠金额（ETH）")
	if err := parseFlags(fs, args); err != nil {
		return err
	}
	amount, err := parseEther(*amountText)
	if err != nil {
		return fmt.Errorf("无效的 -amount：%w", err)
	}
	privateKey, err := loadPrivateKey()
	if err != nil {
		return err
	}
	contractAddress := common.HexToAddress(donationContractHex)
	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	data, err := bindings.NewBeggingContract().TryPackDonate()
	if err != nil {
		return fmt.Errorf("编码 Donate 调用失败：%w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := dialSepolia(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	hash, err := sendValueTransaction(ctx, client, privateKey, from, contractAddress, amount, data)
	if err != nil {
		return err
	}
	fmt.Println("捐赠者：", from.Hex())
	fmt.Println("合约地址：", contractAddress.Hex())
	fmt.Println("捐赠金额：", formatWei(amount), "ETH")
	fmt.Println("交易哈希：", hash.Hex())
	fmt.Println("查询状态：go run . tx get -hash", hash.Hex())
	fmt.Println("查看事件：go run . donation events -tx", hash.Hex())
	return nil
}

func (donationEventsCommand) Execute(args []string) error {
	fs := flag.NewFlagSet("donation events", flag.ContinueOnError)
	hashText := fs.String("tx", "", "捐赠交易哈希")
	if err := parseFlags(fs, args); err != nil {
		return err
	}
	if !isHexHash(*hashText) {
		return errors.New("请提供有效的 -tx 交易哈希")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := dialSepolia(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(*hashText))
	if err != nil {
		return fmt.Errorf("查询交易回执失败：%w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("交易执行失败，没有成功的捐赠事件")
	}
	contractAddress := common.HexToAddress(donationContractHex)
	contract := bindings.NewBeggingContract()
	eventID := contract.GetABI().Events["Donated"].ID
	found := false
	for _, entry := range receipt.Logs {
		if entry.Address != contractAddress || len(entry.Topics) == 0 || entry.Topics[0] != eventID {
			continue
		}
		event, err := contract.UnpackDonatedEvent(entry)
		if err != nil {
			return fmt.Errorf("解析 Donated 事件失败：%w", err)
		}
		found = true
		fmt.Println("捐赠者：", event.Donor.Hex())
		fmt.Println("本次金额：", formatWei(event.Amount), "ETH")
		fmt.Println("累计金额：", formatWei(event.TotalDonated), "ETH")
	}
	if !found {
		fmt.Println("这笔交易没有 Donated 事件")
	}
	return nil
}

func safeError(err error) string {
	message := err.Error()
	for _, secret := range []string{os.Getenv("SEPOLIA_RPC_URL"), os.Getenv("ETHERSCAN_API_KEY"), os.Getenv("SEPOLIA_PRIVATE_KEY")} {
		if secret != "" {
			message = strings.ReplaceAll(message, secret, "[REDACTED]")
			message = strings.ReplaceAll(message, url.QueryEscape(secret), "[REDACTED]")
		}
	}
	return message
}

type Command interface {
	Execute(args []string) error
}

type blockCommand struct{}
type txGetCommand struct{}
type txSendCommand struct{}
type accountTxsCommand struct{}
type donationGetCommand struct{}
type donationSendCommand struct{}
type donationEventsCommand struct{}
type helpCommand struct{}

var commands = map[string]Command{
	"block":           blockCommand{},
	"tx get":          txGetCommand{},
	"tx send":         txSendCommand{},
	"account txs":     accountTxsCommand{},
	"donation get":    donationGetCommand{},
	"donation send":   donationSendCommand{},
	"donation events": donationEventsCommand{},
	"help":            helpCommand{},
	"-h":              helpCommand{},
	"--help":          helpCommand{},
}

func run(args []string) error {
	if len(args) == 0 {
		printUsage()
		return nil
	}
	for count := len(args); count > 0; count-- {
		name := strings.Join(args[:count], " ")
		if command, ok := commands[name]; ok {
			return command.Execute(args[count:])
		}
	}
	return fmt.Errorf("未知命令：%s；运行 go run . help 查看用法", strings.Join(args, " "))
}

func (helpCommand) Execute(args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("help 不接受参数：%s", strings.Join(args, " "))
	}
	printUsage()
	return nil
}

func printUsage() {
	fmt.Println(`用法：
  go run . block -number 123456
  go run . tx get -hash 0x...
  go run . tx send -to 0x... -amount 0.01
  go run . account txs -address 0x... [-page 1] [-limit 10]
  go run . donation get -donor 0x...
  go run . donation send -amount 0.01
  go run . donation events -tx 0x...

环境变量：
  SEPOLIA_RPC_URL       区块、交易和捐赠命令所需的 Sepolia RPC 地址
  SEPOLIA_PRIVATE_KEY   tx send 和 donation send 需要的测试账户私钥
  ETHERSCAN_API_KEY     仅 account txs 需要的 Etherscan API Key`)
}

func parseFlags(fs *flag.FlagSet, args []string) error {
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 0 {
		return fmt.Errorf("多余的参数：%s", strings.Join(fs.Args(), " "))
	}
	return nil
}

func dialSepolia(ctx context.Context) (*ethclient.Client, error) {
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	if rpcURL == "" {
		return nil, errors.New("请先设置 SEPOLIA_RPC_URL")
	}
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, errors.New("连接 Sepolia 失败，请检查 RPC 地址和网络")
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("获取链 ID 失败：%w", err)
	}
	if chainID.Cmp(big.NewInt(sepoliaChainID)) != 0 {
		client.Close()
		return nil, fmt.Errorf("RPC 连接到了链 %s，预期是 Sepolia (%d)", chainID, sepoliaChainID)
	}
	return client, nil
}

func (blockCommand) Execute(args []string) error {
	fs := flag.NewFlagSet("block", flag.ContinueOnError)
	number := fs.Uint64("number", 0, "要查询的区块号")
	if err := parseFlags(fs, args); err != nil {
		return err
	}
	found := false
	fs.Visit(func(f *flag.Flag) { found = found || f.Name == "number" })
	if !found {
		return errors.New("请提供 -number 区块号")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := dialSepolia(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(*number))
	if err != nil {
		return fmt.Errorf("查询区块失败：%w", err)
	}
	fmt.Println("区块号：", block.Number())
	fmt.Println("区块哈希：", block.Hash().Hex())
	fmt.Println("时间戳：", block.Time())
	fmt.Println("UTC 时间：", time.Unix(int64(block.Time()), 0).UTC())
	fmt.Println("交易数量：", len(block.Transactions()))
	return nil
}

func (txGetCommand) Execute(args []string) error {
	fs := flag.NewFlagSet("tx get", flag.ContinueOnError)
	hashText := fs.String("hash", "", "交易哈希")
	if err := parseFlags(fs, args); err != nil {
		return err
	}
	if !isHexHash(*hashText) {
		return errors.New("请提供有效的 0x 开头、32 字节的 -hash")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := dialSepolia(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	hash := common.HexToHash(*hashText)
	tx, pending, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		return fmt.Errorf("查询交易失败：%w", err)
	}
	fmt.Println("交易哈希：", tx.Hash().Hex())
	fmt.Println("Nonce：", tx.Nonce())
	fmt.Println("接收地址：", tx.To())
	fmt.Println("金额：", formatWei(tx.Value()), "ETH")
	fmt.Println("Gas 上限：", tx.Gas())
	if pending {
		fmt.Println("状态：待确认")
		return nil
	}
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		return fmt.Errorf("查询交易回执失败：%w", err)
	}
	status := "失败"
	if receipt.Status == types.ReceiptStatusSuccessful {
		status = "成功"
	}
	fmt.Println("状态：", status)
	fmt.Println("区块号：", receipt.BlockNumber)
	fmt.Println("实际 Gas：", receipt.GasUsed)
	return nil
}

func (txSendCommand) Execute(args []string) error {
	fs := flag.NewFlagSet("tx send", flag.ContinueOnError)
	toText := fs.String("to", "", "接收地址")
	amountText := fs.String("amount", "", "ETH 金额，例如 0.01")
	if err := parseFlags(fs, args); err != nil {
		return err
	}
	if !common.IsHexAddress(*toText) {
		return errors.New("请提供有效的 -to 以太坊地址")
	}
	value, err := parseEther(*amountText)
	if err != nil {
		return fmt.Errorf("无效的 -amount：%w", err)
	}
	privateKey, err := loadPrivateKey()
	if err != nil {
		return err
	}
	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	to := common.HexToAddress(*toText)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := dialSepolia(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	hash, err := sendValueTransaction(ctx, client, privateKey, from, to, value, nil)
	if err != nil {
		return err
	}
	fmt.Println("发送地址：", from.Hex())
	fmt.Println("接收地址：", to.Hex())
	fmt.Println("金额：", formatWei(value), "ETH")
	fmt.Println("交易哈希：", hash.Hex())
	fmt.Println("查询状态：go run . tx get -hash", hash.Hex())
	return nil
}

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	privateKeyText := strings.TrimPrefix(os.Getenv("SEPOLIA_PRIVATE_KEY"), "0x")
	if privateKeyText == "" {
		return nil, errors.New("请先设置 SEPOLIA_PRIVATE_KEY")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyText)
	if err != nil {
		return nil, errors.New("SEPOLIA_PRIVATE_KEY 格式无效")
	}
	return privateKey, nil
}

func sendValueTransaction(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey, from, to common.Address, value *big.Int, data []byte) (common.Hash, error) {
	nonce, err := client.PendingNonceAt(ctx, from)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取 Nonce 失败：%w", err)
	}
	gas, err := client.EstimateGas(ctx, ethereum.CallMsg{From: from, To: &to, Value: value, Data: data})
	if err != nil {
		return common.Hash{}, fmt.Errorf("估算 Gas 失败，请检查接收地址是否为拒收 ETH 的合约：%w", err)
	}
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取 Gas 小费失败：%w", err)
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取最新区块失败：%w", err)
	}
	if header.BaseFee == nil {
		return common.Hash{}, errors.New("最新区块缺少 BaseFee")
	}
	feeCap := new(big.Int).Add(new(big.Int).Mul(header.BaseFee, big.NewInt(2)), tip)
	maxGasCost := new(big.Int).Mul(new(big.Int).SetUint64(gas), feeCap)
	needed := new(big.Int).Add(value, maxGasCost)
	balance, err := client.BalanceAt(ctx, from, nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取余额失败：%w", err)
	}
	if balance.Cmp(needed) < 0 {
		return common.Hash{}, fmt.Errorf("余额不足：最多需要 %s ETH（含 Gas 上限），当前有 %s ETH", formatWei(needed), formatWei(balance))
	}
	chainID := big.NewInt(sepoliaChainID)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: feeCap,
		Gas:       gas,
		To:        &to,
		Value:     value,
		Data:      data,
	})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("签名失败：%w", err)
	}
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		return common.Hash{}, fmt.Errorf("发送交易失败：%w", err)
	}
	return signedTx.Hash(), nil
}

type accountTx struct {
	BlockNumber string `json:"blockNumber"`
	TimeStamp   string `json:"timeStamp"`
	Hash        string `json:"hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	IsError     string `json:"isError"`
}

func (accountTxsCommand) Execute(args []string) error {
	fs := flag.NewFlagSet("account txs", flag.ContinueOnError)
	address := fs.String("address", "", "要查询的地址")
	page := fs.Int("page", 1, "页码")
	limit := fs.Int("limit", 10, "每页条数")
	if err := parseFlags(fs, args); err != nil {
		return err
	}
	if !common.IsHexAddress(*address) {
		return errors.New("请提供有效的 -address 以太坊地址")
	}
	if *page < 1 || *limit < 1 || *limit > 100 {
		return errors.New("page 必须大于 0，limit 必须介于 1 到 100")
	}
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		return errors.New("请先设置 ETHERSCAN_API_KEY（Infura Key 不能用于此接口）")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	txs, err := fetchAccountTxs(ctx, apiKey, *address, *page, *limit)
	if err != nil {
		return err
	}
	if len(txs) == 0 {
		fmt.Println("没有找到普通交易记录")
		return nil
	}
	for _, tx := range txs {
		when := "时间未知"
		if seconds, err := strconv.ParseInt(tx.TimeStamp, 10, 64); err == nil {
			when = time.Unix(seconds, 0).UTC().Format(time.RFC3339)
		}
		amount := "?"
		if wei, ok := new(big.Int).SetString(tx.Value, 10); ok {
			amount = formatWei(wei)
		}
		status := "成功"
		if tx.IsError == "1" {
			status = "失败"
		}
		fmt.Printf("%s | 区块 %s | %s | %s ETH | %s -> %s\n  %s\n", when, tx.BlockNumber, status, amount, tx.From, tx.To, tx.Hash)
	}
	return nil
}

func fetchAccountTxs(ctx context.Context, apiKey, address string, page, limit int) ([]accountTx, error) {
	values := url.Values{
		"chainid": {"11155111"},
		"module":  {"account"},
		"action":  {"txlist"},
		"address": {address},
		"page":    {fmt.Sprint(page)},
		"offset":  {fmt.Sprint(limit)},
		"sort":    {"desc"},
		"apikey":  {apiKey},
	}
	requestURL := "https://api.etherscan.io/v2/api?" + values.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, errors.New("创建 Etherscan 请求失败")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求 Etherscan 失败：%w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Etherscan 返回 HTTP %d", resp.StatusCode)
	}
	var result struct {
		Status  string          `json:"status"`
		Message string          `json:"message"`
		Result  json.RawMessage `json:"result"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 4<<20)).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析 Etherscan 响应失败：%w", err)
	}
	if result.Status != "1" {
		var detail string
		_ = json.Unmarshal(result.Result, &detail)
		if strings.EqualFold(detail, "No transactions found") {
			return nil, nil
		}
		return nil, fmt.Errorf("Etherscan 查询失败：%s：%s", result.Message, detail)
	}
	var txs []accountTx
	if err := json.Unmarshal(result.Result, &txs); err != nil {
		return nil, fmt.Errorf("解析交易列表失败：%w", err)
	}
	return txs, nil
}

func parseEther(s string) (*big.Int, error) {
	parts := strings.Split(s, ".")
	if len(parts) > 2 || len(parts) == 0 || !decimalDigits(parts[0]) {
		return nil, errors.New("金额必须是大于 0 的十进制数")
	}
	fraction := ""
	if len(parts) == 2 {
		fraction = parts[1]
		if !decimalDigits(fraction) || len(fraction) > 18 {
			return nil, errors.New("小数部分必须有 1 到 18 位数字")
		}
	}
	whole, _ := new(big.Int).SetString(parts[0], 10)
	wei := new(big.Int).Mul(whole, big.NewInt(params.Ether))
	if fraction != "" {
		fractionWei, _ := new(big.Int).SetString(fraction+strings.Repeat("0", 18-len(fraction)), 10)
		wei.Add(wei, fractionWei)
	}
	if wei.Sign() <= 0 {
		return nil, errors.New("金额必须大于 0")
	}
	return wei, nil
}

func decimalDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func formatWei(wei *big.Int) string {
	unit := big.NewInt(params.Ether)
	whole, fraction := new(big.Int).QuoRem(wei, unit, new(big.Int))
	if fraction.Sign() == 0 {
		return whole.String()
	}
	fractionText := fraction.String()
	return whole.String() + "." + strings.TrimRight(strings.Repeat("0", 18-len(fractionText))+fractionText, "0")
}

func isHexHash(s string) bool {
	if len(s) != 66 || !strings.HasPrefix(s, "0x") {
		return false
	}
	for _, r := range s[2:] {
		if !strings.ContainsRune("0123456789abcdefABCDEF", r) {
			return false
		}
	}
	return true
}
