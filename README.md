# 《用Go来做以太坊开发》

xxx
## Step
``` shell
# 初始化以太坊客户端：ganache，简单，方便本地测试； geth：可以用其搭建私有节点或全节点
$ npm install -g ganache  #ganache-cli
# 用于创建本地区块链以快速开发以太坊的工具
# 文档看https://trufflesuite.com/ganache/
# 也可直接使用UI客户端：https://github.com/trufflesuite/ganache-ui，较大
# 运行ganache cli客户端
# ganache-cli
# 可以使用相同的助记词来生成相同序列的公开地址
# ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
$ ganache -m "much repair shock carbon improve miss forget sock include bullet interest solution"  #--detach

# 开始
$ mkdir goethereumbook
$ cd goethereumbook
$ go mod init yunlabs.com/goethereumbook       #初始化 go.mod
# go get -u github.com/spf13/cobra/cobra       #手动使用cobra
# https://github.com/spf13/cobra
# https://github.com/spf13/cobra-cli
$ go install github.com/spf13/cobra-cli@latest #直接使用cobra CLI init
# cobra-cli init --author "George george@yunlabs.com" --license apache --viper
# 或者直接将配置放到~/.cobra.yaml
author: George Liu <george@yunlabs.com>
license: MIT
useViper: true
$ cobra-cli init
$ go run cli/main.go #run root command

# git
$ echo "# goethereumbook" >> README.md
$ git init
# git add ...
$ git commit -m "init"
$ git branch -M main
$ git remote add origin git@github.com:george-taotaome/goethereumbook.git
$ git push -u origin main

# chapter 1
$ go get github.com/ethereum/go-ethereum/ethclient
$ cobra-cli add chapter1
$ go run cli/main.go chapter1

...

# 智能合约
# 安装 solc  https://soliditylang.org/
$ brew update
$ brew tap ethereum/ethereum
$ brew install solidity
$ solc --version

# 安装abigen工具, 学习更复杂的智能合约看truffle framework
$ go get -u github.com/ethereum/go-ethereum
$ cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v1.12.2
$ make
$ make devtools

# Store.sol
$ solc --abi contracts/Store.sol -o contracts/build --overwrite
$ solc --bin contracts/Store.sol -o contracts/build --overwrite
# 合并上面两步，需--optimize，不然提示Runtime error: code size to deposit exceeds maximum code size
# 原因gas费太少，auth.GasLimit加多个0 OK
# 单个交易可以执行的最大 gas 数量是 6,700,000，这对应于大约 250KB 的合约代码大小
$ solc contracts/Store.sol --bin --abi --optimize -o ./contracts/build --overwrite
$ abigen --bin=contracts/build/Store.bin --abi=contracts/build/Store.abi --pkg=store --out=contracts/store/Store.go
$ go run contract/deploy.go
#搞不定，一直提示：2023/09/13 16:22:27 VM Exception while processing transaction: invalid opcode
#先实践 truffle
$ npm install -g truffle
$ npm install -g ganache
$ ganache -m "much repair shock carbon improve miss forget sock include bullet interest solution"  #--detach
#最后确定是ganache-cli不给力，改用ganache后发布合约OK

#创建一个ERC20智能合约
$ solc contracts/ERC20.sol --bin --abi --optimize -o ./contracts/build
$ abigen --bin=contracts/build/ERC20.bin --abi=contracts/build/ERC20.abi --pkg=token --out=contracts/token/ERC20.go


```


## Useful Links
- [用Go来做以太坊开发](https://goethereumbook.org)
