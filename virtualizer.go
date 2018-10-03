package main

type Node interface {
}

type Props map[string]interface{}

type VNode struct {
	Type     string
	Props    Props
	Children []*VNode
	Text     string
}

func H(tpe string, props Props, children ...*VNode) *VNode {
	return &VNode{
		Type:     tpe,
		Props:    props,
		Children: children,
	}
}
