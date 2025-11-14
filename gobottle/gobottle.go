package gobottle

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义了路由处理函数的签名。
// 它接收一个封装后的 Request 和 Response，并返回一个错误。
type HandlerFunc func(*Request, *Response) error

// GoBottle 是框架的核心结构，它管理路由并处理 HTTP 请求。
type GoBottle struct {
	// routes 存储了所有注册的路由。
	// 键是 HTTP 方法（如 "GET"），值是一个映射，将路径模式映射到对应的处理函数。
	routes map[string]map[string]HandlerFunc
}

// New 创建并返回一个新的 GoBottle 应用实例。
func New() *GoBottle {
	return &GoBottle{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

// ServeHTTP 实现了 http.Handler 接口。
// 这是 Go 的 HTTP 服务器调用的入口点，用于处理每个传入的请求。
func (g *GoBottle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1. 封装请求和响应对象
	request := NewRequest(r)
	response := NewResponse(w)

	// 2. 查找并调用匹配的路由处理函数
	handler, err := g.findHandler(r.Method, r.URL.Path)
	if err != nil {
		// 如果没有找到匹配的路由，返回 404 Not Found
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, "404 Not Found: %s\n", r.URL.Path)
		return
	}

	// 3. 执行处理函数
	if err := handler(request, response); err != nil {
		// 如果处理函数返回错误，返回 500 Internal Server Error
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Internal Server Error: %v\n", err)
	}
}

// findHandler 根据 HTTP 方法和路径查找对应的处理函数。
// 目前只支持静态路径匹配。
func (g *GoBottle) findHandler(method, path string) (HandlerFunc, error) {
	// 获取该 HTTP 方法对应的所有路由
	methodRoutes, exists := g.routes[method]
	if !exists {
		return nil, fmt.Errorf("method %s not allowed", method)
	}

	// 查找精确匹配的路径
	handler, exists := methodRoutes[path]
	if !exists {
		return nil, fmt.Errorf("path %s not found", path)
	}

	return handler, nil
}

// GET 是一个便捷方法，用于注册处理 GET 请求的路由。
func (g *GoBottle) GET(path string, handler HandlerFunc) {
	g.addRoute("GET", path, handler)
}

// POST 是一个便捷方法，用于注册处理 POST 请求的路由。
func (g *GoBottle) POST(path string, handler HandlerFunc) {
	g.addRoute("POST", path, handler)
}

// addRoute 向应用中添加一个新的路由。
func (g *GoBottle) addRoute(method, path string, handler HandlerFunc) {
	// 延迟初始化该 HTTP 方法的路由映射
	if g.routes[method] == nil {
		g.routes[method] = make(map[string]HandlerFunc)
	}
	g.routes[method][path] = handler
}