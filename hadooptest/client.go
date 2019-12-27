package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/colinmarc/hdfs"
	"github.com/hortonworks/gohadoop"
	"github.com/hortonworks/gohadoop/hadoop_common"
	"github.com/hortonworks/gohadoop/hadoop_common/ipc/client"
	"github.com/hortonworks/gohadoop/hadoop_yarn"
	"github.com/nu7hatch/gouuid"
	"github.com/sediah/hdfs_ha"
	"log"
	"time"
)

func main() {
	hdfsclient := Hdfs("10.94.90.52:2181,10.94.90.53:2181,10.94.90.54:2181")
	client, err := hdfs.New(hdfsclient.Master())
	//Client(hdfs.ClientOptions{Addresses:[]string{"10.94.90.52:2181","10.94.90.53:2181","10.94.90.54:2181"},
	//	})
	if err != nil {
		beego.Error(err)
	}
	fmt.Println(client.GetBlocksResponseProto())
}

type HdfsUtil interface {
	Master() string
	FsStat()
}

type hdfsStr struct {
	addr string
	ha   *hdfs_ha.HdfsHa
}

func Hdfs(addr string) HdfsUtil {
	ha, _ := hdfs_ha.New(addr, 5*time.Second, "jietiaocluster", true)
	return &hdfsStr{
		ha:   ha,
		addr: addr,
	}
}

func (h *hdfsStr) Master() string {
	server, _ := h.ha.GetActiveNameNode()
	return server
}

func (h *hdfsStr) FsStat() {
	client, _ := hdfs.New(h.Master())
	i, _ := client.StatFs()

	fmt.Println(i)

}

func hortoworks() {
	var err error

	clientId, _ := uuid.NewV4()
	ugi, _ := gohadoop.CreateSimpleUGIProto()
	c := &ipc.Client{ClientId: clientId, Ugi: ugi, ServerAddress: "0.0.0.0:28081"}
	var clientProtocolVersion uint64 = 1
	var methodName string
	var protocolName string

	// ApplicationClientProtocol.getApplications
	methodName = "getApplications"
	protocolName = "org.apache.hadoop.yarn.api.ApplicationClientProtocolPB"
	getAppsRpcProto := hadoop_common.RequestHeaderProto{MethodName: &methodName, DeclaringClassProtocolName: &protocolName, ClientProtocolVersion: &clientProtocolVersion}
	applicationStates := []hadoop_yarn.YarnApplicationStateProto{hadoop_yarn.YarnApplicationStateProto_ACCEPTED, hadoop_yarn.YarnApplicationStateProto_RUNNING, hadoop_yarn.YarnApplicationStateProto_SUBMITTED}
	getAppsReqProto := hadoop_yarn.GetApplicationsRequestProto{ApplicationStates: applicationStates}
	getAppsResProto := hadoop_yarn.GetApplicationsResponseProto{}
	log.Println("Calling rpc method: ", methodName)
	err = c.Call(&getAppsRpcProto, &getAppsReqProto, &getAppsResProto)
	if err != nil {
		log.Fatal("Client.call failed", err)
	}
	log.Println("Returned response: ", getAppsResProto)

	// ApplicationClientProtocol.getNewApplication
	methodName = "getNewApplication"
	protocolName = "org.apache.hadoop.yarn.api.ApplicationClientProtocolPB"
	getNewAppRpcProto := hadoop_common.RequestHeaderProto{MethodName: &methodName, DeclaringClassProtocolName: &protocolName, ClientProtocolVersion: &clientProtocolVersion}
	getNewAppReqProto := hadoop_yarn.GetNewApplicationRequestProto{}
	getNewAppResProto := hadoop_yarn.GetNewApplicationResponseProto{}
	log.Println("Calling rpc method: ", methodName)
	err = c.Call(&getNewAppRpcProto, &getNewAppReqProto, &getNewAppResProto)
	if err != nil {
		log.Fatal("Client.call failed", err)
	}
	log.Println("Returned response: ", getNewAppResProto)
}
