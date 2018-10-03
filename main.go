package main

import "github.com/dennwc/dom"

var document *dom.Document

func init() {
	document = dom.GetDocument()
}

func main() {
	a := &VNode{
		Type: "ul",
		Props: Props{
			"className": "main",
			"style":     "background-color: red;",
		},
		Children: []*VNode{
			{
				Type: "li",
				Children: []*VNode{
					{
						Text: "Item 1",
					},
				},
			},
			{
				Type: "li",
				Children: []*VNode{
					{
						Text: "Item 2",
					},
				},
			},
		},
	}

	b := &VNode{
		Type: "ul",

		Children: []*VNode{
			{
				Type: "li",
				Children: []*VNode{
					{
						Text: "Item 1",
					},
				},
			},
			{
				Type: "li",
				Children: []*VNode{
					{
						Text: "Hello World!",
					},
				},
			},
		},
	}

	c := make(chan struct{}, 0)
	root := document.GetElementById("root")
	reload := document.GetElementById("reload")

	UpdateElement(root, a, nil, 0)

	reload.AddEventListener("click", func(event dom.Event) {
		println(event.Type())
		UpdateElement(root, b, a, 0)
	})

	<-c
}
