package main

import (
	"fmt"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/mysql"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func main() {
	sql := "select * from wn where id = 2"

	pa := parser.New()

	//set charsetInfo and collation
	charsetInfo, collation := "utf8", "utf8_general_ci"

	pa.SetSQLMode(mysql.ModeNone)

	stmtNodes, _, _ := pa.Parse(sql, charsetInfo, collation)

	fmt.Println(stmtNodes[0])
}
