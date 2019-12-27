package main

import (
	"fmt"
	"time"

	"sync-pg/lib"

	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	datasourcename := "root:123456@tcp(10.94.10.49:3306)/test?charset=utf8"
	db, err := sql.Open("mysql", datasourcename)
	if err != nil {
		fmt.Println("链接失败！")
	}

	for iiint := 1; iiint >= 0; iiint-- {
		tx, _ := db.Begin()
		stmt, terr := tx.Prepare("INSERT INTO ppstesttable_7_14 ( notice_no, order_no, notice_method, " +
			"biz_type, biz_data, priority_level, belong_system, notice_system, notice_num, notice_state, " +
			"date_notice, remark, date_created, created_by, date_updated, updated_by ) " +
			"VALUES (?, ?,?, ?,?, ?,?, ?,?, ?,?, ?,?, ?,?, ?);")
		if terr != nil {
			beego.Error(fmt.Sprintf("失败！%v", terr))
			//stmt.Close()
			tx.Rollback()

		}
		for i := 1000; i >= 0; i-- {
			var ii []interface{}
			ii = append(ii, "NT3661440891703488512")
			ii = append(ii, "big")
			ii = append(ii, "is.api.v3.order.repayfeedback")
			ii = append(ii, "rong360_is.api.v3.order.bindcardfeedback")
			ii = append(ii, "{\"order_no\":\"252193487402763\",\"reason\":\"\",\"bind_status\":1}")
			ii = append(ii, 5)
			ii = append(ii, "PPS")
			ii = append(ii, "rong360")
			ii = append(ii, "1")
			ii = append(ii, "0")
			ii = append(ii, lib.NowTimeToString())
			ii = append(ii, "this is haha !")
			ii = append(ii, lib.NowTimeToString())
			ii = append(ii, "sys")
			ii = append(ii, lib.NowTimeToString())
			ii = append(ii, "sys")
			_, serr := stmt.Exec(ii...)
			if serr != nil {
				beego.Error(fmt.Sprintf("执行失败！%v", serr))
				//stmt.Close()
				tx.Rollback()

			}
		}
		stmt.Close()
		fmt.Println(tx.Commit())
		//fmt.Println(db.Query("show tables"))

		fmt.Println(iiint, "完毕！")
		time.Sleep(2 * time.Second)
	}
}
