package gobottle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response 封装了 http.ResponseWriter，提供便捷的方法来构建响应。
type Response struct {
	http.ResponseWriter
	statusCode int
}

// NewResponse 创建一个新的 Response 实例。
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{ResponseWriter: w, statusCode: http.StatusOK}
}

// Header 返回响应头映射，用于设置自定义头。
func (r *Response) Header() http.Header {
	return r.ResponseWriter.Header()
}

// WriteHeader 发送一个带有状态码的响应头。
func (r *Response) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

// Write 写入响应体数据。
func (r *Response) Write(data []byte) (int, error) {
	return r.ResponseWriter.Write(data)
}

// String 以字符串形式写入响应体，并设置 Content-Type 为 text/plain。
func (r *Response) String(format string, values ...interface{}) (int, error) {
	if r.Header().Get("Content-Type") == "" {
		r.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}
	return fmt.Fprintf(r, format, values...)
}

// HTML 以 HTML 形式写入响应体，并设置 Content-Type 为 text/html。
func (r *Response) HTML(html string) (int, error) {
	r.Header().Set("Content-Type", "text/html; charset=utf-8")
	return r.Write([]byte(html))
}

// JSON 将数据编码为 JSON 格式写入响应体，并设置 Content-Type 为 application/json。
func (r *Response) JSON(data interface{}) error {
	r.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(r).Encode(data)
}

// Status 返回当前响应的状态码。
func (r *Response) Status() int {
	return r.statusCode
}