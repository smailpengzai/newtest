package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

const (
	//host1     = "10.94.10.49"
	//port1     = 54322
	//user1     = "datacenter"
	//password1 = "123456"
	//dbname1   = "postgres"
	host1     = "10.208.52.201"
	port1     = 5432
	user1     = "gpadmin"
	password1 = "360@guoxin!2017"
	dbname1   = "datacenter"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host1, port1, user1, password1, dbname1)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

type execlQureyTable struct {
	SuccessExc []string
	FailExc    []string
}

func main() {
	var execlQureyTable execlQureyTable
	vacuumschema := flag.String("vacuumschema", "", "ampc_withhold_baofu,pc_withhold")
	execSql := flag.String("execsql", "", "select updata ")
	flag.Parse()
	if *vacuumschema != "" {
		vacuumschemas := strings.Split(*vacuumschema, ",")
		fmt.Println(fmt.Sprintf("需要处理的%d个库", len(vacuumschemas)))
		totalTimeStart := time.Now()
		for _, schemaname := range vacuumschemas {
			fmt.Println(fmt.Sprintf("开始处理：%v", schemaname))
			listsql := make([]string, 2)
			listsql = append(listsql, fmt.Sprintf("select 'VACUUM  VERBOSE '    ||schemaname||    '.'    ||tablename||    ';' from pg_tables t inner join pg_namespace d on t.schemaname=d.nspname  where schemaname='%v';", schemaname))
			listsql = append(listsql, fmt.Sprintf("select 'VACUUM ANALYZE VERBOSE '    ||schemaname||    '.'    ||tablename||    ';' from pg_tables t inner join pg_namespace d on t.schemaname=d.nspname  where schemaname='%v';", schemaname))
			for _, vacuumSql := range listsql {
				requestdata := execsql(vacuumSql)
				for _, v := range requestdata.Datas {
					fmt.Println(fmt.Sprintf("开始执行 %v ！", v[0]))
					starttime := time.Now()
					requstdata := execsql(fmt.Sprintf("%v", v[0]))
					if requstdata.Status {
						execlQureyTable.FailExc = append(execlQureyTable.FailExc, fmt.Sprintf("%v", v[0]))
					} else {
						execlQureyTable.SuccessExc = append(execlQureyTable.SuccessExc, fmt.Sprintf("%v", v[0]))
					}
					fmt.Println(fmt.Sprintf(" %v 执行完毕！耗时：%v s \n\n", v[0], time.Now().Sub(starttime).Seconds()))
				}
			}
		}

		fmt.Println(fmt.Sprintf("本次数回收总耗时：%v s \n 成功执行%v 条sql \n \n \n ", time.Now().Sub(totalTimeStart).Seconds(), len(execlQureyTable.SuccessExc)))
		if len(execlQureyTable.FailExc) != 0 {
			fmt.Println(fmt.Sprintf("下面是失败的sql，请手动执行检查其原因：\n"))
			for _, failsql := range execlQureyTable.FailExc {
				fmt.Println(failsql)
			}
		}
	}
	if *execSql != "" {
		fmt.Println("需要执行的sql：", *execSql)
		requestdata := execsql(*execSql)
		fmt.Println("执行结果：", requestdata)
	}
}

type RequstData struct {
	Columns []string
	Datas   [][]interface{}
	Status  bool // false 成功 true 失败
}

func execsql(exesql string) (requstData RequstData) {
	db := connect()
	defer db.Close()
	requst, err := db.Query(exesql)
	if err != nil {
		beego.Error(fmt.Sprintf("错误信息：%v", err))
		requstData.Status = true
	}
	requstData.Columns, _ = requst.Columns()
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
