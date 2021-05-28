package model

import (
	"untils"
	"fmt"
)

type Test struct {
	ID   int
	Name string
}

func (test *Test) AddTest() error {
	sqlStr := "insert into sys_test(id,name) values (?,?)"

	inStmt, e := untils.Db.Prepare(sqlStr)

	if e != nil {
		fmt.Println("预编译出现错误", e)
	}

	_, sqlError := inStmt.Exec("5", "yang")

	if sqlError != nil {
		fmt.Println("执行出现异常", sqlError)
		return sqlError
	}
	return nil
}

func (test *Test) GetTestById() (*Test, error) {

	sqlStr := "select id,name from sys_test where id = ?"

	row := untils.Db.QueryRow(sqlStr, test.ID)

	var id int
	var name string

	e := row.Scan(&id, &name)

	if e != nil {
		return nil, e
	}

	u := &Test{
		ID:   id,
		Name: name,
	}
	return u, e
}

func (test *Test) GetTests() ([] *Test, error) {
	sqlStr := "select id,name from sys_test"

	rows, e := untils.Db.Query(sqlStr)

	if e != nil {
		return nil, e
	}

	var tests []*Test
	for rows.Next() {
		var id int
		var name string

		e := rows.Scan(&id, &name)

		if e != nil {
			return nil, e
		}

		u := &Test{
			ID:   id,
			Name: name,
		}
		tests = append(tests, u)
	}
	return tests, nil
}
