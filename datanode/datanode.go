package datanode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"shell"
	"sync"
)

type KV struct {
	BlockId string `json:"块ID"`
	DestIp  string `json:"目的IP"`
}

// NodeInfo 用于json和结构体对象的互转
type NodeInfo struct {
	NodeName   string `json:"nodeName"`   //节点hostname 通过配置文件获取
	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
	Port       string `json:"port"`       //节点端口号
	HitRatio   string `json:"hit_ratio"`  //命中率   （每X秒更新）
	//缓存列表 指针
	//命中率的信息
}
type ValueList []KV
type Cache struct {
	ValueList `json:"value_list"` //缓存值列表
	Cap       int                 `json:"cap"` //缓存容量
}
type DataNodeInfo struct {
	NodeInfo `json:"node_info"` //节点信息
	Cache    `json:"cache"`     //缓存
}


func (dn *DataNodeInfo) GetHitRatio(command string,ret *string)error  {
	command="../../loganalyer/src/05-analyer/05-analyer -input ~/Documents/slave1rKV.json"
	*ret=shell.GetCacaheInfo(command)
	return nil
}

//DataNode节点的格式化输出
func (dn DataNodeInfo) String() string {
	nn, _ := json.Marshal(dn)
	var out bytes.Buffer
	json.Indent(&out, nn, "", "\t")
	return out.String()
}

//Node节点的格式化输出
func (node NodeInfo) String() string {
	nn, _ := json.Marshal(node)
	var out bytes.Buffer
	json.Indent(&out, nn, "", "\t")
	return out.String()
}

// GetNode 获取Node
func GetNode(Name, IP, Port string) DataNodeInfo {
	NN := new(DataNodeInfo)
	NN.NodeName = Name
	NN.Port = Port
	NN.NodeIpAddr = IP
	return *NN
}

// StartDNClient 启动DataNode Client
func StartDNClient(dn DataNodeInfo, NameNodeIP string) *rpc.Client {
	var wg sync.WaitGroup
	client, _ := rpc.DialHTTP("tcp", NameNodeIP)
	//var nodelist []NodeInfo
	var datanodelist []NodeInfo
	asyncCall := client.Go("NameNodeInfo.AddNode", dn.NodeInfo, &datanodelist, nil)
	<-asyncCall.Done
	for _, info := range datanodelist {
		fmt.Printf(info.String() + " ")
	}
	fmt.Println()
	wg.Add(1)
	go StartDNRPCServer(dn)
	wg.Wait()
	return client
}

// GetNodeIP 返回值类似于 "127.0.0.1/8|30004"
func GetNodeIP(nodename string, client *rpc.Client) string {
	var ipPort string
	client.Call("NameNodeInfo.GetNodeIP", nodename, &ipPort)
	return ipPort
}
func StartDNRPCServer(dn DataNodeInfo) {
	rpc.Register(dn)
	rpc.HandleHTTP()

	log.Printf("Serving RPC server:%v on port %v", dn.NodeName, dn.Port)
	if err := http.ListenAndServe(":"+dn.Port, nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
