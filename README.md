# 安装环境

sudo add-apt-repository ppa:longsleep/golang-backports

sudo apt-get update

sudo apt-get install golang-go

~$ go version

go version go1.14 linux/amd6

# 环境配置

国内无法访问国外的情况下，我们需要对环境作简单的配置，才能访问到开源的第三方的库或者工具

使用国内七牛云的 go module 镜像。

参考 https://github.com/goproxy/goproxy.cn。

golang 1.14 可以直接执行：

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

除了七牛云，还可以使用阿里云的 golang 国内镜像。

https://mirrors.aliyun.com/goproxy/

设置方法

go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct


# 学习内容

通过观看golang Google的资深工程师讲解的视频教程来学习，结合项目的需要，其中可能穿插有docker等工具的使用，我们学习过程中有使用到。