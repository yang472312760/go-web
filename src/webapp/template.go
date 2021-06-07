package main

import (
	"net/http"
	"html/template"
	"model"
)

func handlerTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/webapp/index.html", "src/webapp/hello.html"))
	t.ExecuteTemplate(w, "hello.html", "我是Hello模板")
	//t.Execute(w, "Hello World")
}

func ageTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/webapp/age.html"))

	var age int = 20

	t.Execute(w, age > 18)
}

func rangeTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/webapp/range.html"))
	var tests []*model.Test
	test := &model.Test{
		ID:   1,
		Name: "yang1",
	}
	tests = append(tests, test)
	test1 := &model.Test{
		ID:   2,
		Name: "yang2",
	}
	tests = append(tests, test1)
	test2 := &model.Test{
		ID:   3,
		Name: "yang3",
	}
	tests = append(tests, test2)

	t.Execute(w, tests)
}

func testWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/webapp/with.html"))

	t.Execute(w, "狸猫")
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/webapp/template1.html", "src/webapp/template2.html"))
	t.Execute(w, "我能将数据传入到两个template吗？")
}

func testModel(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/webapp/model2.html"))
	t.ExecuteTemplate(w, "model2", "model2")
}
func testModel1(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/webapp/model1.html"))
	t.ExecuteTemplate(w, "model", "model2")
}

func main() {
	http.HandleFunc("/template", handlerTemplate)
	http.HandleFunc("/age", ageTemplate)
	http.HandleFunc("/range", rangeTemplate)
	http.HandleFunc("/with", testWith)
	http.HandleFunc("/testTemplate", testTemplate)
	http.HandleFunc("/testModel", testModel)
	http.HandleFunc("/testModel1", testModel1)
	http.ListenAndServe(":8080", nil)
}
