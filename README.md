# Sepolia 区块链交互示例

这个 Go 命令行程序可以查询 Sepolia 区块、查询单笔交易、发送测试 ETH，以及查询某个地址的普通交易记录。

## 准备环境

1. 安装 Go，并在项目目录下载依赖：

   ```bash
   go mod download
   ```

2. 按要使用的命令，在当前终端配置环境变量：

   | 环境变量 | 用途 |
   | --- | --- |
   | `SEPOLIA_RPC_URL` | Sepolia RPC 地址；`block`、`tx get`、`tx send` 需要 |
   | `SEPOLIA_PRIVATE_KEY` | 发送交易的测试账户私钥；仅 `tx send` 需要，可带或不带 `0x` 前缀 |
   | `ETHERSCAN_API_KEY` | Etherscan API Key；仅 `account txs` 需要，不能用 Infura Key 代替 |

   例如，`SEPOLIA_RPC_URL` 可以使用 Infura 提供的 Sepolia HTTPS RPC 地址。私钥只用于专门的测试账户，不要写进代码、README 或 Git 提交，也不要分享给别人。

3. 查看命令帮助：

   ```bash
   go run . help
   ```

## 查询区块

```bash
go run . block -number 123456
```

将 `123456` 换成要查询的 Sepolia 区块号。程序输出区块号、哈希、时间戳、UTC 时间和交易数量。区块号 `0` 也可以查询。

## 查询单笔交易

```bash
go run . tx get -hash 0x你的交易哈希
```

交易哈希必须是 `0x` 开头的 32 字节十六进制值。程序输出金额、接收地址、Gas 上限和交易状态；已确认的交易还会输出区块号和实际 Gas 用量。

## 发送测试 ETH

先确认发送账户在 Sepolia 上有足够的测试 ETH，接收地址也确实属于 Sepolia 上的目标账户。转账金额之外还需要 Gas 费用。

```bash
go run . tx send -to 0x接收地址 -amount 0.01
```

`-amount` 的单位是 ETH，最多支持 18 位小数。程序会从 `SEPOLIA_PRIVATE_KEY` 推导发送地址，估算 Gas、检查余额、签名并发送交易，然后输出交易哈希。**执行这条命令会实际向 Sepolia 广播交易**；重复执行会再次发送。

发送后可以用 `tx get` 命令查询该哈希的状态。

## 查询地址的历史交易

```bash
go run . account txs -address 0x要查询的地址
go run . account txs -address 0x要查询的地址 -page 2 -limit 20
```

该命令使用 [Etherscan V2 的 `account/txlist` 接口](https://docs.etherscan.io/api-reference/endpoint/txlist)，按时间从新到旧输出普通交易。默认查询第 1 页，每页 10 条；`-limit` 可设为 1 到 100。它不列出合约内部转账或 ERC-20 代币转账事件。

## 检查代码

```bash
go test ./...
go vet ./...
```
