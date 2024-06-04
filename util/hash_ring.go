package util

import (
	"hash/crc32"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	// DefaultVirtualNodeSize 默认虚拟节点
	DefaultVirtualNodeSize = 400
)

// NewHashRingNode 创建 一个 哈希环 节点
func NewHashRingNode[T any](nodeKey string, obj T, weight int) *HashRingNode[T] {
	if weight <= 0 {
		weight = 1
	}

	h := &HashRingNode[T]{
		node:    obj,
		nodeKey: nodeKey,
		weight:  weight,
	}
	return h
}

type HashRingNode[T any] struct {
	node    T
	weight  int
	nodeKey string
}

func (this_ *HashRingNode[T]) GetNode() (res T) {
	return this_.node
}

type HashRingVirtualNode[T any] struct {
	hashRingNode   *HashRingNode[T]
	virtualNodeKey string
	hash           uint32
}

type hashRingNodesArray[T any] []*HashRingVirtualNode[T]

func (p hashRingNodesArray[T]) Len() int           { return len(p) }
func (p hashRingNodesArray[T]) Less(i, j int) bool { return p[i].hash < p[j].hash }
func (p hashRingNodesArray[T]) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p hashRingNodesArray[T]) Sort()              { sort.Sort(p) }

// HashRing 哈希环 存放 Node
type HashRing[T any] struct {
	virtualNodeSize int
	nodes           hashRingNodesArray[T]
	nodeCache       map[string]*HashRingNode[T]
	mu              sync.RWMutex
	Hash            func(s string) uint32
}

// NewHashRing 创建 哈希环 指定 虚拟节点数量  默认使用 DefaultVirtualNodeSize
func NewHashRing[T any](virtualNodeSize int, obj T) *HashRing[T] {
	if virtualNodeSize <= 0 {
		virtualNodeSize = DefaultVirtualNodeSize
	}

	h := &HashRing[T]{
		virtualNodeSize: virtualNodeSize,
		nodeCache:       make(map[string]*HashRingNode[T]),
	}
	return h
}

// AddNodes add nodes to hash ring
func (h *HashRing[T]) AddNodes(nodes []*HashRingNode[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, node := range nodes {
		h.nodeCache[node.nodeKey] = node
	}
	h.generate()
}

// AddNode add node to hash ring
func (h *HashRing[T]) AddNode(node *HashRingNode[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.nodeCache[node.nodeKey] = node
	h.generate()
}

// RemoveNode remove node
func (h *HashRing[T]) RemoveNode(node *HashRingNode[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.nodeCache, node.nodeKey)
	h.generate()
}

// UpdateNode update node with weight
func (h *HashRing[T]) UpdateNode(node *HashRingNode[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.nodeCache[node.nodeKey] = node
	h.generate()
}

func (h *HashRing[T]) generate() {
	var totalW int
	for _, n := range h.nodeCache {
		totalW += n.weight
	}

	totalVirtualNodeSize := h.virtualNodeSize * len(h.nodeCache)
	h.nodes = hashRingNodesArray[T]{}

	for nodeKey, node := range h.nodeCache {
		virtualNodeSize := int(math.Floor(float64(node.weight) / float64(totalW) * float64(totalVirtualNodeSize)))
		for i := 0; i < virtualNodeSize; i++ {

			virtualNode := &HashRingVirtualNode[T]{
				hashRingNode: node,
			}
			virtualNode.virtualNodeKey = nodeKey + "#VN-" + strconv.Itoa(i)

			virtualNode.hash = h.GetHash(virtualNode.virtualNodeKey)

			h.nodes = append(h.nodes, virtualNode)
		}
	}
	h.nodes.Sort()
}

// GetNode get node with key
func (h *HashRing[T]) GetNode(s string) (node *HashRingVirtualNode[T]) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.nodes) == 0 {
		return
	}

	v := h.GetHash(s)
	i := sort.Search(len(h.nodes), func(i int) bool { return h.nodes[i].hash >= v })

	if i == len(h.nodes) {
		i = 0
	}

	return h.nodes[i]
}

func (h *HashRing[T]) GetHash(s string) (res uint32) {

	if h.Hash != nil {
		res = h.Hash(s)
	} else {
		res = crc32.ChecksumIEEE([]byte(s))
	}
	return
}
