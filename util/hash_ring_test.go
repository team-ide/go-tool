package util

import (
	"fmt"
	"strconv"
	"testing"
)

type TestNode struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

func TestHashRing(t *testing.T) {
	var hashRing *HashRing[*TestNode]
	var hashRingNodeObj = &TestNode{}
	hashRing = NewHashRing(10, hashRingNodeObj)
	hashRing.AddNode(NewHashRingNode("chat-node-1", &TestNode{
		Name: "chat-node-1",
		Ip:   "192.168.0.1",
		Port: 11011,
	}, 1))

	fmt.Println("#### 只有 `chat-node-1` 节点时候")
	fmt.Println()
	outNode(hashRing)

	fmt.Println()
	outSession(hashRing)
	fmt.Println()

	fmt.Println("#### 新增 `chat-node-2` 节点时候")

	hashRing.AddNode(NewHashRingNode("chat-node-2", &TestNode{
		Name: "chat-node-2",
		Ip:   "192.168.0.2",
		Port: 11011,
	}, 1))

	fmt.Println()
	outNode(hashRing)
	fmt.Println()

	outSession(hashRing)
	fmt.Println()

	fmt.Println("#### 新增 `chat-node-3` 节点时候")

	hashRing.AddNode(NewHashRingNode("chat-node-3", &TestNode{
		Name: "chat-node-3",
		Ip:   "192.168.0.3",
		Port: 11011,
	}, 1))

	fmt.Println()
	outNode(hashRing)
	fmt.Println()

	outSession(hashRing)
	fmt.Println()

}

func outNode(hashRing *HashRing[*TestNode]) {
	var size = len(hashRing.nodes)

	fmt.Println()
	fmt.Println("* 哈希环")
	fmt.Println()
	fmt.Println("|  节点   | 虚拟节点 | Hash 区间| 数量(万) | ")
	fmt.Println("|:-----:|:----:|:----:|:----:|")
	var nodeCount = map[string]uint32{}
	for i := 0; i < size; i++ {
		node := hashRing.nodes[i]
		var startHash = node.hash
		var endHash uint32
		if i == size-1 {
			lastNode := hashRing.nodes[0]
			endHash = lastNode.hash - 1
		} else {
			endHash = hashRing.nodes[i+1].hash - 1
		}
		var count uint32
		if endHash < startHash {
			count = (uint32(4294967295) - startHash) + endHash + 1
		} else {
			count = endHash - startHash + 1
		}
		nodeCount[node.hashRingNode.nodeKey] += count
		cStr := strconv.FormatFloat(float64(count)/10000, 'f', 2, 64)
		if endHash < startHash {
			fmt.Println("|", node.hashRingNode.nodeKey, "|", node.virtualNodeKey, "|", startHash, "~ 4294967295,0 ~", endHash, "|", cStr, "|")
		} else {
			fmt.Println("|", node.hashRingNode.nodeKey, "|", node.virtualNodeKey, "|", startHash, "~", endHash, "|", cStr, "|")
		}
	}
	for k, v := range nodeCount {
		fmt.Println("* 节点:", k, " 总数量:", v)
	}
}

func outSession(hashRing *HashRing[*TestNode]) {

	fmt.Println()
	fmt.Println("* 会话信息")
	fmt.Println()
	fmt.Println("|  会话信息   | Hash 值 | 所属节点| 虚拟节点 | 状态 |")
	fmt.Println("|:-----:|:----:|:----:|:----:|:----:|")

	for i := 1; i <= 10; i++ {
		key := fmt.Sprintf("p2p_%d_%d", i, i+1)
		hash := hashRing.GetHash(key)
		node := hashRing.GetNode(key)
		cStr := "是当前节点，正常"
		if node.hashRingNode.nodeKey != "chat-node-1" {
			cStr = "非当前节点，标记清理"
		}
		fmt.Println("|", key, "|", hash, "|", node.hashRingNode.nodeKey, "|", node.virtualNodeKey, "|", cStr, "|")
	}
}
