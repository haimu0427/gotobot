# GoBottle - 轻量级Go Web框架

GoBottle是一个受Python Bottle框架启发的轻量级Go Web框架，专为学习和教学目的设计。

## 项目结构

```
gobottle/
├── go.mod              # Go模块定义
├── cmd/
│   └── main.go         # 示例应用入口
├── gobottle/
│   ├── gobottle.go     # 核心框架代码
│   ├── request.go      # 请求封装
│   └── response.go     # 响应封装
└── doc/
    └── bottle_minimal.py  # Python Bottle参考实现
```

## 快速开始

### 安装依赖

```bash
go mod tidy
```

### 运行示例

```bash
go run cmd/main.go
```

访问 http://localhost:8080 查看效果。

## 核心特性

- **简洁的路由系统**: 支持GET和POST请求
- **请求封装**: 提供便捷的查询参数访问
- **响应封装**: 支持HTML、JSON、纯文本响应
- **中间件支持**: 基于Go的http.Handler接口

## 使用示例

```go
package main

import (
    "log"
    "net/http"
    "gobottle/gobottle"
)

func main() {
    // 创建应用实例
    app := gobottle.New()

    // 注册路由
    app.GET("/", func(req *gobottle.Request, res *gobottle.Response) error {
        return res.HTML("<h1>Hello, GoBottle!</h1>")
    })

    app.GET("/hello", func(req *gobottle.Request, res *gobottle.Response) error {
        name := req.Query("name")
        if name == "" {
            name = "World"
        }
        return res.HTML("<h1>Hello, " + name + "!</h1>")
    })

    app.POST("/api/data", func(req *gobottle.Request, res *gobottle.Response) error {
        return res.JSON(map[string]string{"status": "success"})
    })

    // 启动服务器
    log.Println("Server starting on :8080...")
    http.ListenAndServe(":8080", app)
}
```

## API文档

### 路由注册

- `GET(path, handler)` - 注册GET请求处理函数
- `POST(path, handler)` - 注册POST请求处理函数

### 请求对象 (Request)

- `Query(key)` - 获取URL查询参数
- `PostForm(key)` - 获取表单数据
- `Body()` - 获取请求体

### 响应对象 (Response)

- `HTML(html)` - 发送HTML响应
- `JSON(data)` - 发送JSON响应
- `String(format, ...)` - 发送纯文本响应
- `WriteHeader(code)` - 设置状态码

## 学习价值

这个项目是学习Go语言Web开发的绝佳资源：

1. **理解HTTP服务器原理**: 学习如何处理HTTP请求和响应
2. **掌握接口设计**: 了解如何设计简洁易用的API
3. **实践Go语言特性**: 运用结构体、接口、方法等核心概念
4. **Web框架设计**: 理解Web框架的基本架构和工作原理

## 扩展建议

- 添加中间件支持
- 实现路径参数解析
- 增加模板引擎支持
- 添加静态文件服务
- 实现会话管理