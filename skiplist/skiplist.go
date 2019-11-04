package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//节点
type Node struct {
	NextNode *Node
	Down     *Node
	data     int
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (node *Node) Compare(node2 *Node) int {
	return node.data - node2.data
}

func (node *Node) Copy() *Node {
	return &Node{data: node.data}
}

//链表
type LinkedQueue struct {
	//层数最底层 为0
	Level       int
	Header      Node
	Tail        *Node
	TailPreNode *Node
}

func NewLinkedQueue() *LinkedQueue {
	linkedQueue := &LinkedQueue{}
	linkedQueue.Tail = &linkedQueue.Header
	return linkedQueue
}

func (s *LinkedQueue) Insert(node *Node) {
	temp := s.Header.NextNode
	preNode := &s.Header

	if s.Tail != &s.Header && s.Tail.Compare(node) <= 0 {
		s.Tail.NextNode = node
		s.Tail = node
		return
	}

	for ; temp != nil; temp = temp.NextNode {
		if temp.Compare(node) >= 0 {
			node.NextNode = preNode.NextNode
			preNode.NextNode = node
			break
		}
		preNode = temp
	}

	//说明已经是最大的
	if temp == nil {
		s.Tail.NextNode = node
		s.Tail = node
	}
}

//在一个节点后面插入
func (s *LinkedQueue) InsertAfter(after, node *Node) {
	if after.NextNode == nil {
		if after != s.Tail {
			panic("after.NextNode is nil but not is s.tail")
		}
		//之前忽略了这重置尾指针 ,导致数据不全
		s.Tail.NextNode = node
		s.Tail = node
		return
	}

	temp := after.NextNode
	preNode := after

	if s.Tail != &s.Header && s.Tail.Compare(node) <= 0 {
		s.Tail.NextNode = node
		s.Tail = node
		return
	}

	for ; temp != nil; temp = temp.NextNode {
		if temp.Compare(node) >= 0 {
			node.NextNode = preNode.NextNode
			preNode.NextNode = node
			break
		}
		preNode = temp
	}

	//说明已经是最大的
	if temp == nil {
		s.Tail.NextNode = node
		s.Tail = node
	}
}

//从开始找到第一个大于searchNode的节点的前一个节点 ,或者等于searchNode节点
func (s *LinkedQueue) FindIndexNode(searchNode *Node) *Node {
	if searchNode == nil {
		panic("searchNode cant be nil")
	}
	//是否为空队列
	if s.Tail == &s.Header {
		return nil
	}

	temp := s.Header.NextNode
	preNode := &s.Header
	//如果队第一个元素 大于待搜索节点返回 nil, 说明后面的搜索 ,插入都可以从 header 开始
	if temp.Compare(searchNode) >= 0 {
		return nil
	}

	if s.Tail.Compare(searchNode) <= 0 {
		return s.Tail
	}

	for ; temp != nil; temp = temp.NextNode {
		if temp.Compare(searchNode) == 0 {
			return temp
		}
		if temp.Compare(searchNode) >= 0 {
			return preNode
		}
		preNode = temp
	}

	return s.Tail
}

//从afterNode从开始找到第一个大于searchNode的节点的前一个节点, ,或者等于searchNode节点
func (s *LinkedQueue) FindIndexAfterNode(afterNode, searchNode *Node) *Node {
	if afterNode == nil {
		return s.FindIndexNode(searchNode)
	}
	if s.Tail == &s.Header {
		log.Print(" s.Tail == &s.Header cannot appear at FindIndexAfterNode")
		return nil
	}

	if afterNode.NextNode == nil {
		return afterNode
	}

	temp := afterNode
	preNode := afterNode

	if s.Tail.Compare(searchNode) <= 0 {
		return s.Tail
	}

	for ; temp != nil; temp = temp.NextNode {
		if temp.Compare(searchNode) == 0 {
			return temp
		}
		if temp.Compare(searchNode) > 0 {
			return preNode
		}
		preNode = temp
	}
	return s.Tail
}

//从afterNode从开始找到第一个大于searchNode的节点的前一个节点, ,或者等于searchNode节点
func (s *LinkedQueue) FindAfterNode(afterNode, searchNode *Node) *Node {
	if afterNode == nil {
		return s.FindIndexNode(searchNode)
	}
	if s.Tail == &s.Header {
		log.Print(" s.Tail == &s.Header cannot appear at FindIndexAfterNode")
		return nil
	}
	temp := afterNode
	if s.Tail.Compare(searchNode) <= 0 {
		return s.Tail
	}

	for ; temp != nil; temp = temp.NextNode {
		if temp.Compare(searchNode) == 0 {
			return temp
		}
		if temp.Compare(searchNode) > 0 {
			return nil
		}
	}
	return s.Tail
}

func (s *LinkedQueue) Display() {
	temp := s.Header.NextNode
	for ; temp != nil; temp = temp.NextNode {
		fmt.Printf("[%d],", temp.data)
	}
	fmt.Println()
}

//跳表
const MaxLevel = 10

type Skiplist struct {
	TopLevel int
	Queues   []*LinkedQueue
}

func NewSkiplis() *Skiplist {
	var queues []*LinkedQueue
	//最底层是数据层  [1 - MaxLevel] 是索引层
	for i := 0; i < MaxLevel+1; i++ {
		queues = append(queues, NewLinkedQueue())
	}
	return &Skiplist{Queues: queues, TopLevel: 1}
}

func (s *Skiplist) Display() {
	for index, queue := range s.Queues {
		fmt.Printf("level%d:\t", index)
		queue.Display()
		fmt.Println()
	}
}

func (s *Skiplist) Insert(node *Node) {
	if node == nil {
		return
	}

	newNodeLevel := RandLevel()
	if s.TopLevel < newNodeLevel {
		s.TopLevel = newNodeLevel
	}

	newNodes := make([]*Node, newNodeLevel+1)
	//node.level = 0
	newNodes[0] = node
	for i := 1; i < newNodeLevel+1; i++ {
		newNodes[i] = NewNode(node.data)
		newNodes[i].Down = newNodes[i-1]
		//newNodes[i].level = i
		//fmt.Printf("newNodes:%+v\n", newNodes[i])
	}
	//查询索引
	var indexNode, preIndexNode *Node
	//从 level 最顶层开始查找   , level 0 相当于是数据
	for i := s.TopLevel; i > 0; i-- {
		//查找node  在当前层的索引
		if preIndexNode == nil {
			indexNode = s.Queues[i].FindIndexNode(node)
		} else {
			indexNode = s.Queues[i].FindIndexAfterNode(preIndexNode.Down, node)
		}
		//查找node  在当前层的索引
		if newNodeLevel >= i {
			if indexNode == nil {
				s.Queues[i].Insert(newNodes[i])
			} else {
				s.Queues[i].InsertAfter(indexNode, newNodes[i])
			}
		}
		preIndexNode = indexNode
	}

	//插入数据
	if indexNode != nil {
		//fmt.Printf(" %+v \n", indexNode.Down)
		s.Queues[0].InsertAfter(indexNode.Down, node)
	} else {
		s.Queues[0].Insert(node)
	}
}

func (s *Skiplist) FindNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	//查询索引
	var indexNode, preIndexNode *Node
	//从 level 最顶层开始查找   , level 0 相当于是数据
	for i := s.TopLevel - 1; i > 0; i-- {
		//查找node  在当前层的索引
		if preIndexNode == nil {
			indexNode = s.Queues[i].FindIndexNode(node)
		} else {
			indexNode = s.Queues[i].FindIndexAfterNode(preIndexNode.Down, node)
		}
		preIndexNode = indexNode
	}
	var retNode *Node
	if indexNode != nil {
		// 注意这里 indexNode 是上层索引的节点
		retNode = s.Queues[0].FindIndexAfterNode(indexNode.Down, node)
	} else {
		retNode = s.Queues[0].FindIndexNode(node)
	}

	if retNode == nil || retNode.data != node.data {
		return nil
	}
	return retNode
}

func RandLevel() int {
	var max = 0x1 << (MaxLevel + 1)
	rand.Seed(time.Now().UnixNano())
	v := uint(rand.Int31n(int32(max)))

	return GetLevel(v, MaxLevel)
}

func GetLevel(num uint, maxLevel uint) int {
	var tmp uint = 0x1<<(maxLevel-1) - 1
	var tmp2 uint = 0x1<<(maxLevel) - 1
	var mask uint = 0xFFFFFFFF - tmp

	level := 0
	for ; num > 0; {
		if num&mask == 0 {
			break
		}
		level++
		num = num << 1
		num = num & tmp2
	}
	return level
}
