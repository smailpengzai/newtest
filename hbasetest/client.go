package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tsuna/gohbase/filter"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var host = flag.String("host", "10.94.90.52:2181,10.94.90.53:2181,10.94.90.54:2181", "The location where HBase is running")

//var host = flag.String("host", "10.94.90.51:60010", "The location where HBase is running")

func main() {
	///**
	//10.94.90.51 hadoopa
	// 10.94.90.52 hadoops
	// 10.94.90.53 hadoop1
	// 10.94.90.54 hadoop2
	// 10.94.90.55 hadoop4.add.bjyt.qihoo.net
	//**/
	//beego.Notice("===")
	//Hclient := newAdminClient()
	//beego.Notice("===",)
	//beego.Notice(Hclient)
	// status,_ := Hclient.ClusterStatus()
	//fmt.Println(status)
	//分裂
	//httpget("http://10.94.90.51:60010/table.jsp?action=split&name=gx&key=0e3513b7dc366000bac13bc2dd3b24f3")
	//压缩
	//httpget("http://10.94.90.51:60010/table.jsp?action=compact&name=gx&key=1bc16731a1d0fe9c98a38ca45cd3cdcd")

	rsp, str := GetHbaseIndex("", "", "gx", 100)

	beego.Notice(str)

	for _, v := range rsp {
		fmt.Println(v)
	}
}

func GetHbaseIndex(startRow, stopRow, table string, limit int64) (rsp []*hrpc.Result, str string) {
	var nextIndex string
	temScanner, err := PagedQuery(table, startRow, stopRow, limit)
	if err != nil {
		fmt.Println("GetHbaseIndex with limit: ", err)
	}
	if int64(len(temScanner)) <= limit {
		nextIndex = ""
		return temScanner, nextIndex
	} else {
		nextIndex = string(temScanner[limit].Cells[0].Row)
		return temScanner[:len(temScanner)-1], nextIndex
	}
}

func PagedQuery(table, startRow, stopRow string, limit int64) (rsp []*hrpc.Result, err error) {
	var (
		scanRequest *hrpc.Scan
		res         *hrpc.Result
	)
	pFilter := filter.NewPageFilter(limit + 1)
	scanRequest, err = hrpc.NewScanRangeStr(context.Background(), table, startRow, stopRow, hrpc.Reversed(), hrpc.Filters(pFilter))
	if err != nil {
		beego.Error("hrpc.NewScanStr: %s", err.Error())
	}
	scanner := newclient().Scan(scanRequest)
	for {
		res, err = scanner.Next()
		if err == io.EOF || res == nil {
			break
		}
		if err != nil {
			beego.Error("hrpc.Scan: %s", err.Error())
		}
		rsp = append(rsp, res)
	}
	return rsp, err
}

func newAdminClient() gohbase.AdminClient {
	beego.Notice(*host)
	return gohbase.NewAdminClient(*host)
}
func newclient() gohbase.Client {
	return gohbase.NewClient(*host, gohbase.ZookeeperRoot("/"))
}

func httpget(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	resCode := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println("状态码：", resCode)
	fmt.Println(string(body))
}
