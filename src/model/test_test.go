package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("开始测试")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试Test相关方法")
	//t.Run("测试添加Test:", testAddTest)
	//t.Run("测试查询一条数据", testGetTestById)
	t.Run("测试查询所有数据", testGetTests)
}

func testAddTest(t *testing.T) {
	fmt.Println("测试添加Test")
	test := &Test{}
	test.AddTest()
}

func testGetTestById(t *testing.T) {
	fmt.Println("测试查询一条纪录")
	test := &Test{
		ID: 1,
	}
	u, _ := test.GetTestById()
	fmt.Println("得到Test信息：", u)
}

func testGetTests(t *testing.T) {
	fmt.Println("测试查询所有记录")
	test := &Test{}
	tests, _ := test.GetTests()

	for k, v := range tests {
		fmt.Printf("第%d个用户是%v\n", k + 1, v)

	}

}
