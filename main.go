package main

import (
	"fmt"
	"illusion-net/p2p"
)

func main() {
	var nodeinfo p2p.NodeInfo
	nodeinfo.IpAddress = "192.168.168.1"
	nodeinfo.Port = "9999"
	nodeinfo.BlockHeight = "1"
	p2p.AddNode("test1", nodeinfo)

	if p2p.GetNodeInfo("test12") == (p2p.NodeInfo{}) {
		fmt.Println("empty")
	}

}
