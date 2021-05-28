package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w,"请求头中Accept-Encoding的信息是: ",r.Header["Accept-Encoding"])
	fmt.Fprintln(w,"请求头中Accept-Encoding的属性是: ",r.Header.Get("Accept-Encoding"))
}

func main() {

	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8080", nil)

}
