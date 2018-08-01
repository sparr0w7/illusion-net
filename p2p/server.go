package p2p

type NodeInfo struct {
	IpAddress   string
	Port        string
	BlockHeight string
}

var nodelist = make(map[string]NodeInfo)

func AddNode(privateKey string, n NodeInfo) {
	nodelist[privateKey] = n
}

func DeleteNode(privateKey string) {
	delete(nodelist, privateKey)
}

func ReplaceNode(privateKey string, n NodeInfo) {
	nodelist[privateKey] = n
}

func GetNodeInfo(privateKey string) NodeInfo {
	return nodelist[privateKey]
}
