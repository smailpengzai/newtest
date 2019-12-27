package main

import (
	"fmt"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"

	"log"
)

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "10.94.10.49:3306"
	cfg.User = "root"
	cfg.Password = "123456"
	c, err := canal.NewCanal(cfg)

	if err != nil {
		fmt.Println(err)
	}
	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{canal.DummyEventHandler{}})

	// Start canal
	c.RunFrom(mysql.Position{Name: "mysql-bin.000001", Pos: 4})

}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {

	//log.Println(fmt.Sprintf("%s %s %v\n 个数：%v", e.Action, e.Table,e.Rows,len(e.Rows[0])))
	log.Println("OnRow", e.Action)
	for _, datatmp := range e.Rows {
		for index, value := range e.Table.Columns {
			if len(datatmp) == len(e.Table.Columns) {
				log.Println(len(datatmp), len(e.Table.Columns))
				log.Println(e.Table.Name, value.Name, ":", datatmp[index])
			}
		}
	}
	return nil
}

func (h *MyEventHandler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	log.Println(string(queryEvent.Query))

	return nil
}
func (h *MyEventHandler) OnRotate(roateEvent *replication.RotateEvent) error {
	var err error
	//log.Println("OnRotate", roateEvent.NextLogName)
	return err
}
func (h *MyEventHandler) OnTableChanged(schema string, table string) error {
	var err error
	//log.Println("OnTableChanged", schema, table)
	return err
}
func (h *MyEventHandler) OnXID(nextPos mysql.Position) error {
	var err error
	log.Println("OnXID nextPOS;-------", nextPos.Name, nextPos.String())
	return err
}

func (h *MyEventHandler) OnGTID(gtid mysql.GTIDSet) error {
	var err error
	//log.Println("OnGTID", gtid.String())
	return err
}
func (h *MyEventHandler) OnPosSynced(pos mysql.Position, force bool) error {
	var err error
	log.Println("OnPosSynced", pos.String(), force)
	return err
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}
