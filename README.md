
# task-queue

一个简单的任务队列

<!-- PROJECT SHIELDS -->

本篇README.md面向开发者

## 目录

- [上手指南](#上手指南)
    - [安装步骤](#安装步骤)
- [使用demo](#创建worker对象)
    - [创建worker对象](#创建worker对象)
    - [添加任务](#添加任务)
    - [结束任务](#结束任务)

### 上手指南
###### **安装步骤**


```sh
go get github.com/wei-yg/task-queue
```

### 使用demo
###### **创建worker对象**
```go
worker := NewWorker()
```
###### **添加任务**
```go
go func(){  // 这里为什么要加另外的协程，请看 task_test.go 有注释说明
    worker.Push(func () error {
        // 你的任务1 注意 push进来的是一个返回error的函数，当有一个任务返回error的时候，整个worker也将停止
    })

    worker.Push(func() error{
        // 你的任务2 
    })
}()
```

###### **结束任务**
```go
worker.Stop()
```
