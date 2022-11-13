package client

import (
	DataNode "datanode"
	"fmt"
	"net/rpc"
	"strings"
)

// StartClient 启动Client
func StartClient(message string, NameNodeIP string) {
	client, _ := rpc.DialHTTP("tcp", NameNodeIP)
	ops:=strings.Split(message," ")
	switch ops[0] {
	case "getDNList": //获取DataNode 列表
		var datanodelist []DataNode.NodeInfo
		asyncCall := client.Go("NameNodeInfo.GetNodeList", "", &datanodelist, nil)
		<-asyncCall.Done
		for _, info := range datanodelist {
			fmt.Printf(info.String() + " ")
		}
	case "getCacheHitRatio"://获取当前节点的命中率
		var output string
		asyncCall := client.Go("NameNodeInfo.GetHitRatio", "", &output, nil)
		<-asyncCall.Done
		fmt.Println(output)
	default:

	}
}
