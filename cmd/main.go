package main

import (
	"log"
	"net/http"

	"gobottle/gobottle" // 导入我们自定义的框架包
)

func main() {
	// 创建一个新的 GoBottle 应用实例
	app := gobottle.New()

	// 注册一个处理 GET 请求的路由
	app.GET("/", func(req *gobottle.Request, res *gobottle.Response) error {
		_, err := res.HTML("<h1>Hello, GoBottle!</h1>")
		return err
	})

	// 注册另一个 GET 路由，演示如何使用查询参数
	app.GET("/hello", func(req *gobottle.Request, res *gobottle.Response) error {
		name := req.Query("name")
		if name == "" {
			name = "World"
		}
		_, err := res.HTML("<h1>Hello, " + name + "!</h1>")
		return err
	})

	// 注册一个 POST 路由
	app.POST("/submit", func(req *gobottle.Request, res *gobottle.Response) error {
		// 注意：这里简化处理，实际应用中需要解析表单数据
		return res.JSON(map[string]string{"status": "received"})
	})

	// 启动 HTTP 服务器，监听 8080 端口
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", app); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}