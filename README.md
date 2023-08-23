# 《用Go来做以太坊开发》

## Step
``` shell
# 初始化以太坊客户端：ganache，简单，方便本地测试； geth：可以用其搭建私有节点或全节点
$ npm install -g ganache-cli
# 用于创建本地区块链以快速开发以太坊的工具
# 文档看https://trufflesuite.com/ganache/
# 也可直接使用UI客户端：https://github.com/trufflesuite/ganache-ui，较大
# 运行ganache cli客户端
# ganache-cli
# 可以使用相同的助记词来生成相同序列的公开地址
$ ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"

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
$ go run main.go #run root command

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
$ go run main.go chapter1

...

```


## Useful Links
- [用Go来做以太坊开发](https://goethereumbook.org)
