## 1.golang特点
    .静态编译
    .垃圾回收
    .简洁的符号和语法
    .平坦的类型系统
    .基于CSP的并发模型
    .高效简单的工具链
    .丰富的标准库

## 2.为什么选择go语言
    .编译型语言，运行速度快
    .静态编译没有依赖
    .自身支持并发，充分利用多核
    .社区活跃，大厂支持

## 3.并行和并发
    .并发（concurrent）不是并行（parallel）
    .比如：node.js具有并发的能力，但不能充分利用多核
    .写出一个能充分利用多核的程序需要很深的系统编程积淀
    .go可以轻松地写出跑满所有CPU的程序

## 4.go的程序结构
    package main
    import(
        "fmt"
        "os"
    )
    func main() {
        fmt.Println("hello world!")
    }

### 4.1package：go通过package组织
    .package关键字
    .放在程序的第一行
    .两种package，一种是库package，一种是二进制package
    .二进制package使用main来表示，库package的名字跟go文件所在目录名一致

#### 4.1.1二进制package
    .以`package main`作为文件的第一行
    .有且只有一个main函数

#### 4.1.2引入package
    .通过关键字import来引入其他package
    .多个package通过括号包含起来
    .引入但没有使用的package会引起编译错误

### 4.2main函数
    .二进制程序的入口
    .