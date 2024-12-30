# Async Task System

这是一个用于管理和执行异步任务的 Go 语言库。它支持通过链式调用配置任务，并提供了异常处理机制。

## 特性

- 支持异步任务的创建和执行
- 使用 `context.Context` 管理任务生命周期
- 提供异常捕获和错误处理
- 支持链式调用配置任务

## 安装

确保你已经安装了 Go 语言环境。建议使用 Go 1.18 或更高版本。然后使用以下命令获取该库：

### 使用 `go get`

你可以使用以下命令获取该库：

    go get -u github.com/YouEvanLi/go-sync


### 使用 `go mod`

在你的项目中，确保 `go.mod` 文件中包含以下内容：

    go 1.18
    require (
        github.com/YouEvanLi/go-sync latest
    )

然后运行：
    
    go mod tidy

## 使用方法

1. **创建任务**：使用 `NewTask` 方法创建一个任务。
2. **配置参数**：使用 `WithParams` 方法设置任务参数。
3. **启动任务**：使用 `Start` 方法启动任务。
4. **运行多个任务**：使用 `Async` 方法并行运行多个任务并等待结果。

## 测试

使用 Go 的 `testing` 包进行测试。运行以下命令执行测试：



## 贡献

欢迎贡献代码！请提交 Pull Request 或报告问题。

## 许可证

该项目使用 MIT 许可证。详情请参阅 LICENSE 文件。