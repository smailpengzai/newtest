package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/pb"
	"os"
)

type HbaseUtil interface {
	Master() string
	BackupMasters() (result []string)
	DeadServers() (result []string)
	LiveServers() (result []string)
	Reginservers()
	TableSplit()
	Close()
}

type hbaseStr struct {
	zkAddr        string
	clusterStatus *pb.ClusterStatus
	client        gohbase.Client
}

func Hbase(addr string) HbaseUtil {
	cs, _ := gohbase.NewAdminClient(addr).ClusterStatus()
	c := gohbase.NewClient(addr)
	return &hbaseStr{
		zkAddr:        addr,
		clusterStatus: cs,
		client:        c,
	}
}

func (h *hbaseStr) Master() string {
	return h.clusterStatus.Master.GetHostName()
}

func (h *hbaseStr) Close() {
	h.client.Close()
}

func (h *hbaseStr) BackupMasters() (result []string) {
	for _, each := range h.clusterStatus.BackupMasters {
		result = append(result, each.GetHostName())
	}
	return
}

func (h *hbaseStr) DeadServers() (result []string) {
	for _, each := range h.clusterStatus.DeadServers {
		result = append(result, each.GetHostName())
	}
	return
}

func (h *hbaseStr) LiveServers() (result []string) {
	for _, each := range h.clusterStatus.LiveServers {
		result = append(result, each.GetServer().GetHostName())
	}
	return
}

//type HbaseInfo struct {
// Master string
// MasterPort string
// Standby []string
// StandbyPort []string
// Regionservers []
//}

func (h *hbaseStr) Reginservers() {

	for _, each := range h.clusterStatus.LiveServers {

		for _, i := range each.GetServerLoad().RegionLoads {
			fmt.Println(i)
			fmt.Println(string(i.RegionSpecifier.Value), i.GetStorefileSize_MB(), i.GetReadRequestsCount(),
				i.GetWriteRequestsCount())

		}
		//result = append(result, each.RegionState.RegionInfo.TableName.String())
	}
	return
}

func (h *hbaseStr) TableSplit() {
	//hrpc.NewGet()
	//h.client.Get()
}
func main() {

	// 以Stdout为输出，代替默认的stderr
	logrus.SetOutput(os.Stdout)
	// 设置日志等级
	logrus.SetLevel(logrus.DebugLevel)

	h := Hbase("10.94.90.52:2181,10.94.90.53:2181,10.94.90.54:2181")
	l := h.LiveServers()
	d := h.DeadServers()
	fmt.Println(l, d)
}
