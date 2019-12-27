package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/lib/pq"
	"strings"
	"time"
)

const (
	host     = "10.94.10.49"
	port     = 54322
	user     = "datacenter"
	password = "123456"
	dbname   = "postgres"
)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
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

func main() {
	beego.SetLogFuncCall(true)
	//fmt.Println(pq.CopyIn("test", "content"))
	//fmt.Println(pq.CopyInSchema("lps", "test", "content"))
	pgCopyInsert(Connect())
}
func pgCopyInsert(db *sql.DB) {
	lock := true
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(0)

	txn, err := db.Begin()
	if err != nil {
		beego.Error(err)
	}
	t1 := time.Now()
	stmt, _ := txn.Prepare(pq.CopyIn("ppstesttable_7_14", "id", "notice_no", "order_no", "notice_method", "biz_type", "biz_data", "priority_level", "belong_system", "notice_system", "notice_num", "notice_state", "date_notice", "remark", "date_created", "created_by", "date_updated", "updated_by"))
	for i := 30000000; i < 40000000; i++ {
		_, err := stmt.Exec(i, `NT3661440891703488512`, `big`, `is.api.v3.order.repayfeedback`, `rong360_is.api.v3.order.bindcardfeedback`, `{\"order_no\":\"252193487402763\",\"reason\":\"\",\"bind_status\":1}`, `5`, `PPS`, `rong360`, `1`, `0`, `2019-05-24 14:26:01`, `this is haha !`, `2019-05-24 14:26:01`, `sys`, `2019-05-24 14:26:01`, `sys`)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				beego.Error(err)
			}
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			beego.Error(err)
			pgDelete(db)
			pgCopyInsert(db)
			lock = false
		}
	}
	if lock {
		err = stmt.Close()
		if err != nil {
			beego.Error(err)
		}
		err = txn.Commit()
		if err != nil {
			beego.Error(err)
		}

	}
	elapsed := time.Since(t1)
	fmt.Println("copy insert time: ", elapsed)
}

func pgDelete(db *sql.DB) {
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(0)

	txn, err := db.Query("truncate table datatest_table")
	fmt.Println("truncate table ", txn, err)
}
