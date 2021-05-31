package main

import (
	"net/http"
	"fmt"
	"model"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是: ", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性是: ", r.Header.Get("Accept-Encoding"))

	len := r.ContentLength

	body := make([]byte, len)

	r.Body.Read(body)

	fmt.Fprintln(w, "请求体中的内容有：", string(body))
}

func handForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintln(w, "请求参数为：", r.Form)
	fmt.Fprintln(w, "Post请求参数为：", r.PostForm)
}

func handlerFormValue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过FormValue:", r.FormValue("username"))
	fmt.Fprintln(w, "通过PostFormVlaue:", r.PostFormValue("password"))
}

func handerWrite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("你的请求我已经收到"))
}

func handlerWriteHtml(w http.ResponseWriter, r *http.Request) {
	html := `<html>
				<head>
					<title>测试响应内容为网页</title>
					<meta charset="utf-8"/>
				</head>
				<body>
					<h1>我是以网页的形式响应过来的</h1>
				</body>
			</html>`

	w.Write([]byte(html))
}

func handlerJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	test := model.Test{
		ID:   1,
		Name: "test",
	}

	json, _ := json.Marshal(test)

	w.Write([]byte(json))
}

func handlerRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {

	http.HandleFunc("/hello", handler)

	http.HandleFunc("/form", handForm)

	http.HandleFunc("/formValue", handlerFormValue)

	http.HandleFunc("/hw", handerWrite)

	http.HandleFunc("/html", handlerWriteHtml)

	http.HandleFunc("/json", handlerJson)

	http.HandleFunc("/redirect", handlerRedirect)

	http.ListenAndServe(":8080", nil)

}
