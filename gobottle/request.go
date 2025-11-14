package gobottle

import (
	"io"
	"net/http"
)

// Request 封装了 http.Request，提供便捷的方法来访问请求数据。
type Request struct {
	*http.Request
}

// NewRequest 创建一个新的 Request 实例。
func NewRequest(r *http.Request) *Request {
	return &Request{r}
}

// Query 返回 URL 查询参数。
// 例如，对于路径 /hello?name=world，调用 Query("name") 将返回 "world"。
func (r *Request) Query(key string) string {
	return r.URL.Query().Get(key)
}

// PostForm 返回解析后的表单数据。
// 注意：此方法应在调用 ParseForm 之后使用，通常在处理函数内部自动完成。
func (r *Request) PostForm(key string) string {
	return r.FormValue(key)
}

// Body 返回请求体的 io.ReadCloser。
// 调用者负责在读取完毕后关闭它。
func (r *Request) Body() io.ReadCloser {
	return r.Request.Body
}
