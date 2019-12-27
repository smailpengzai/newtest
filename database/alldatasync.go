package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
)

type requstData struct {
	tableName string
	Columns   []string
	Datas     []interface{}
}

func main() {
	var (
		dataChan chan requstData
		endChan  chan bool
	)
	dataChan = make(chan requstData, 1)
	endChan = make(chan bool, 1)
	go getResourceData("root", "123456", "10.94.10.49", "3306", "test", "ppstesttable_7_14", dataChan, endChan)
	go insertToPg("datacenter", "123456", "10.94.10.49", "54322", "postgres", dataChan, endChan)
	select {}
}

//取数据
func getResourceData(username, password, host, port, dbname, tablename string, dataChan chan requstData, endChan chan bool) {
	var requstData requstData
	datasourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbname)
	db, err := sql.Open("mysql", datasourcename)
	if err != nil {
		fmt.Println("链接失败！")
	}
	requst, _ := db.Query(fmt.Sprintf("select * from %s", tablename))
	requstData.tableName = tablename
	requstData.Columns, _ = requst.Columns()
	//
	//for _, v := range requstData.Columns {
	//	fmt.Println(v)
	//}

	for requst.Next() {
		requsttmp := make([]interface{}, len(requstData.Columns))
		values := make([]json.Number, len(requstData.Columns))
		for i := range values {
			requsttmp[i] = &values[i]
		}
		requst.Scan(requsttmp...)
		requstData.Datas = requsttmp
		dataChan <- requstData

	}
	endChan <- false
	return
}
func insertToPg(username, password, host, port, dbname string, dataChan chan requstData, endChan chan bool) {
	var (
		count int = 0
		datas []requstData
	)
	for {

		datatmp := <-dataChan
		datas = append(datas, datatmp)
		if count == 1000000 {
			fmt.Println("入数据")
			//数据入库
			execDataInsert(username, password, host, port, dbname, datas)
			datas = nil
		}
		if len(endChan) == 1 {
			//数据入库
			execDataInsert(username, password, host, port, dbname, datas)
			return
		}
	}
}

func execDataInsert(username, password, host, port, dbname string, datas []requstData) {
	if len(datas) == 0 {
		return
	}
	db := Connect(host, port, username, password, dbname)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(0)

	txn, err := db.Begin()
	if err != nil {
		fmt.Errorf("开启数据库失败！错误信息：%v", err)
	}
	stmt, _ := txn.Prepare(pq.CopyIn(datas[0].tableName, datas[0].Columns...))
	for _, datatmp := range datas {
		_, err := stmt.Exec(datatmp.Datas...)
		if err != nil {
			fmt.Errorf("放入数据失败，请检查数据的准确性。错误为：%v", err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Errorf("预编译执行失败：%v", err)
	}

	err = stmt.Close()
	if err != nil {
		fmt.Errorf("关闭失败，请检查。错误为：%v", err)
	}
	err = txn.Commit()
	if err != nil {
		fmt.Errorf("提交失败，请检查。错误为：%v", err)
	}

}

func Connect(host, port, user, password, dbname string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
