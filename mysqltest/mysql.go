package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http/pprof"
)

func main() {

	//// 以Stdout为输出，代替默认的stderr
	//logrus.SetOutput(os.Stdout)
	//// 设置日志等级
	//logrus.SetLevel(logrus.DebugLevel)
	//requstData := mysqlexecsql("select * from user")
	//fmt.Println(fmt.Sprintf("返回数据的个数：%v 大小：%v", len(requstData.Datas), unsafe.Sizeof(requstData.Datas)),requstData)
	//log.Println(http.ListenAndServe("0.0.0.0:10010", nil))
	//select {}
	//beforetime := time.Now()
	//for ii := 200000 ; ii>=0 ;ii-- {
	//mysqlexecSql("root", "123456", "10.94.10.49", "3306", "test", "select * from ppstesttable_7_14")
	mysqlInsertSql("root", "S6r7@qHoGr*KmHTNRVROT", "192.168.31.189", "3306", "test",
		`insert into test (Name) values("xiongba3");`)
	//time.Sleep(1000*time.Millisecond)
	//}
	//beego.Notice("耗时：", time.Now().Sub(beforetime).Seconds())
}

type RequstData struct {
	Columns []string
	Datas   [][]interface{}
}

func mysqlexecsql(exesql string) (requstData RequstData) {

	datasourcename := "root:123456@tcp(10.94.10.49:3306)/test?charset=utf8"
	db, err := sql.Open("mysql", datasourcename)
	if err != nil {
		fmt.Println("链接失败！")
	}
	requst, _ := db.Query(exesql)
	requstData.Columns, _ = requst.Columns()

	for _, v := range requstData.Columns {
		fmt.Println(v)
	}
	for requst.Next() {
		requsttmp := make([]interface{}, len(requstData.Columns))
		values := make([]json.Number, len(requstData.Columns))
		for i := range values {
			requsttmp[i] = &values[i]
		}
		requst.Scan(requsttmp...)
		requstData.Datas = append(requstData.Datas, requsttmp)
	}
	return
}

func mysqlexecSql(username, password, host, port, dbname, exesql string) (requstData RequstData) {

	datasourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbname)
	db, err := sql.Open("mysql", datasourcename)
	if err != nil {
		fmt.Println("链接失败！")
	}
	requst, _ := db.Query(exesql)
	requstData.Columns, _ = requst.Columns()
	//
	//for _, v := range requstData.Columns {
	//	fmt.Println(v)
	//}
	requsttmp := make([]interface{}, len(requstData.Columns))
	values := make([]json.Number, len(requstData.Columns))
	for i := range values {
		requsttmp[i] = &values[i]
	}
	var count int = 0
	for requst.Next() {
		requst.Scan(requsttmp...)
		//requstData.Datas = append(requstData.Datas, requsttmp)
		count++
	}
	beego.Notice(fmt.Sprintf("该表拥有数据的大小：%d  ,最后一条数据：%v", count, requsttmp))
	return
}
func mysqlInsertSql(username, password, host, port, dbname, exesql string) (requstData RequstData) {

	datasourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbname)
	beego.Debug(datasourcename)
	db, err := sql.Open("mysql", datasourcename)
	defer db.Close()
	if err != nil {
		fmt.Println("链接失败！")
	}
	requst, err := db.Exec(exesql)
	if err != nil {
		beego.Error("插入失败：", err)
	}

	beego.Notice(requst.LastInsertId())
	beego.Notice(requst.RowsAffected())
	return
}
