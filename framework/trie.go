package framework

import (
	"errors"
	"strings"
)

// trie 前缀树
type tree struct {
	mName HttpMethod
	root  *node
}

func NewTree(mName HttpMethod) *tree {
	return &tree{mName: mName, root: newNode()}
}

// trie 前缀树节点
type node struct {
	isLast  bool
	segment string
	handler ControllerHandler
	childes []*node
}

func newNode() *node {
	return &node{
		isLast:  false,
		childes: []*node{},
		segment: "",
	}
}

// 是否以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, "/:")
}

// 寻找满足segment的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childes) <= 0 {
		return nil
	}
	if isWildSegment(segment) {
		return n.childes
	}
	ansNodes := make([]*node, 0, len(n.childes))
	for i := 0; i < len(n.childes); i++ {
		if isWildSegment(n.childes[i].segment) {
			ansNodes = append(ansNodes, n.childes[i])
		} else if n.childes[i].segment == segment {
			ansNodes = append(ansNodes, n.childes[i])
		}
	}
	return ansNodes
}

// 寻找匹配uri的节点
func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 3)
	segment := "/"+segments[1]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	nodes := n.filterChildNodes(segment)
	if len(nodes) <= 0 {
		return nil
	}
	if len(segments) == 2 {
		for i := 0; i < len(nodes); i++ {
			if nodes[i].isLast {
				return nodes[i]
			}
		}
		return nil
	}
	for i := 0; i < len(nodes); i++ {
		if cNode := nodes[i].matchNode("/"+segments[2]); cNode != nil {
			return cNode
		}
	}
	return nil
}

// 添加节点路由
func (t *tree) AddRouter(uri string, handler ControllerHandler) error {
	if mNode := t.root.matchNode(uri); mNode != nil {
		return errors.New("route exist:" + uri)
	}
	segments := strings.Split(uri, "/")
	n := t.root
	for i, s := range segments {
		if i == 0{
			continue
		}
		s = "/"+s
		isLast := false
		if i == len(segments)-1 {
			isLast = true
		}
		if !isWildSegment(s) {
			s = strings.ToUpper(s)
		}
		nodes := n.filterChildNodes(s)
		var cNode *node
		for _, item := range nodes {
			if item.segment == s {
				cNode = item
				break
			}
		}
		if cNode == nil {
			cNode = &node{childes: make([]*node, 0), segment: s}
			if isLast {
				cNode.isLast = true
				cNode.handler = handler
			}
			n.childes = append(n.childes, cNode)
		}
		n = cNode
	}
	return nil
}

// 寻找路由对应的控制器
func (t *tree) FindHandler(uri string) ControllerHandler {
	matN := t.root.matchNode(uri)
	if matN != nil {
		return matN.handler
	}
	return nil
}
