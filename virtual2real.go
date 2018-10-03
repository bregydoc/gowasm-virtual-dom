package main

import (
	"github.com/dennwc/dom"
	"strings"
)

func SetBooleanProp(target *dom.Element, name string, value interface{}) {
	if value != nil {
		target.SetAttribute(name, value)
		target.SetAttribute(name, true)
	} else {
		target.SetAttribute(name, false)
	}
}

func IsCustomProp(name string) bool {
	return IsEventProp(name) || name == "forceUpdate"
}

func IsEventProp(name string) bool {
	return strings.HasPrefix(name, "on")
}

func ExtractEventName(name string) string {
	return strings.ToLower(name[2:])
}

func RemoveBooleanProp(target *dom.Element, name string) {
	target.RemoveAttribute(name)
	target.SetAttribute(name, false)
}

func RemoveProp(target *dom.Element, name string, value interface{}) {
	_, okIsBool := value.(bool)
	if IsCustomProp(name) {
		return
	} else if name == "className" {
		target.RemoveAttribute("class")
	} else if okIsBool {
		RemoveBooleanProp(target, name)
	} else {
		target.RemoveAttribute(name)
	}
}

func SetProp(target *dom.Element, name string, value interface{}) {
	_, okIsBool := value.(bool)
	if IsCustomProp(name) {
		return
	} else if name == "className" {
		target.SetAttribute("class", value)
	} else if okIsBool {
		SetBooleanProp(target, name, value.(bool))
	} else {
		target.SetAttribute(name, value)
	}

}

func SetProps(target *dom.Element, props Props) {
	for name, value := range props {
		SetProp(target, name, value)
	}
}

func AddEventListeners(target *dom.Element, props Props) {
	for name := range props {
		if IsEventProp(name) {
			target.AddEventListener(ExtractEventName(name), props[name].(dom.EventHandler))
		}
	}
}

func CreateElement(node *VNode) *dom.Element {
	ok := node.Type == "" && node.Text != ""
	if ok {
		return document.CreateTextNode(node.Text)

	}

	element := document.CreateElement(node.Type)
	SetProps(element, node.Props)
	AddEventListeners(element, node.Props)

	for _, c := range node.Children {
		el := CreateElement(c)
		element.AppendChild(el)
	}
	return element

}

func Changed(node1, node2 *VNode) bool {
	return node1.Text != "" && node1 != node2 || node1.Type != node2.Type || node1.Props["forceUpdate"].(bool)
}

func UpdateProp(target *dom.Element, name string, newVal interface{}, oldVal interface{}) {
	if newVal == nil {
		RemoveProp(target, name, oldVal)
	} else if oldVal == nil || newVal != oldVal {
		SetProp(target, name, newVal)
	}
}

func UpdateProps(target *dom.Element, newProps Props, oldProps Props) {
	for n, v := range newProps {
		oldProps[n] = v
	}
	props := oldProps

	for n := range props {
		UpdateProp(target, n, newProps[n], oldProps[n])
	}
}

func UpdateElement(parent *dom.Element, newNode *VNode, oldNode *VNode, index int) {
	if oldNode == nil {
		parent.AppendChild(
			CreateElement(newNode),
		)
	} else if newNode == nil {
		if len(parent.ChildNodes()) > 0 {
			parent.RemoveChild(
				parent.ChildNodes()[index],
			)
		}

	} else if Changed(newNode, oldNode) {
		parent.ReplaceChild(CreateElement(newNode), parent.ChildNodes()[index])
	} else if newNode.Type != "" {
		UpdateProps(parent.ChildNodes()[index], newNode.Props, oldNode.Props)
		newLength := len(newNode.Children)
		oldLength := len(oldNode.Children)
		for i := 0; i < newLength || i < oldLength; i++ {
			UpdateElement(
				parent.ChildNodes()[index],
				newNode.Children[i],
				oldNode.Children[i],
				i,
			)
		}
	}
}
