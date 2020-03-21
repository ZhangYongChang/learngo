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

git config --global credential.helper store


# 程序调优

```go
package main


func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func lengthOfNonRepeatingSubStr2(s string) int {
	lastOccurred := make([]int, 0xffff)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI > start {
			start = lastI
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}
```

## 使用pprof进行程序的性能测试
函数lengthOfNonRepeatingSubStr性能测试的结果

```go
package main

import "testing"

func BenchmarkSubStr(b *testing.B) {
	s := "为是年解放萨的房间内克的思考国家考虑过开放嗯规划法国当局的黑色风格世纪的法国三大和"
	ans := 20
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
```

使用go test来实现性能测试

```shell script
yczhang@yczhang:~/go/src/github.com/ZhangYongChang/learngo/nonrepeatingsubstr$ go test -bench . -cpuprofile cpu.out
goos: linux
goarch: amd64
pkg: github.com/ZhangYongChang/learngo/nonrepeatingsubstr
BenchmarkSubStr-8         202452              5416 ns/op
PASS
ok      github.com/ZhangYongChang/learngo/nonrepeatingsubstr    1.306s
```

分析结果可视化

```shell script
yczhang@yczhang:~/go/src/github.com/ZhangYongChang/learngo/nonrepeatingsubstr$ go tool pprof cpu.out 
File: nonrepeatingsubstr.test
Type: cpu
Time: Feb 27, 2020 at 9:44pm (CST)
Duration: 1.30s, Total samples = 1.25s (96.04%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) web
(pprof) quit
```
![avatar](https://github.com/ZhangYongChang/learngo/tree/master/images/pprof001.svg)

```go
package main

import "testing"

func BenchmarkSubStr(b *testing.B) {
	s := "为是年解放萨的房间内克的思考国家考虑过开放嗯规划法国当局的黑色风格世纪的法国三大和"
	ans := 20
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr2(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
```

```shell script
yczhang@yczhang:~/go/src/github.com/ZhangYongChang/learngo/nonrepeatingsubstr$ go test -bench . -cpuprofile cpu.out
goos: linux
goarch: amd64
pkg: github.com/ZhangYongChang/learngo/nonrepeatingsubstr
BenchmarkSubStr-8          14737             83185 ns/op
PASS
ok      github.com/ZhangYongChang/learngo/nonrepeatingsubstr    2.215s

```

![avatar](https://github.com/ZhangYongChang/learngo/tree/master/images/pprof002.svg)