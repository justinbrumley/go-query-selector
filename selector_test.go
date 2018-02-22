package selector

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"testing"
)

func TestQuerySelector(t *testing.T) {
	file, err := os.Open("./test.html")
	if err != nil {
		t.Error(err)
		return
	}

	doc, err := html.Parse(file)
	if err != nil {
		t.Error(err)
		return
	}

	node := NewNode(doc)
	query := &Query{
		Id: "test",
	}
	testIdNode, err := node.QuerySelector(query)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Found node by id: \n%v\n\n", testIdNode)

	query = &Query{
		Class: "test-class-2",
	}
	classNode, err := node.QuerySelector(query)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Found node by class: \n%v\n\n", classNode)
	fmt.Printf("Inner Content: \n%v\n\n", classNode.Inner())

	query = &Query{
		Class: "nested-class",
	}
	nestedNode, err := node.QuerySelector(query)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Found nested node by class: \n%v\n\n", nestedNode)

	query = &Query{
		Attributes: []Attribute{
			{
				Key:   "type",
				Value: "checkbox",
			},
		},
	}
	attrNode, err := node.QuerySelector(query)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Found node by attribute [type=\"checkbox\"]: \n%v\n\n", attrNode)

	query = &Query{
		Tag: "input",
	}
	tagNode, err := node.QuerySelector(query)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Found node by Tag \"input\": \n%v\n\n", tagNode)

	query = &Query{
		Class: "test-class",
	}
	nodes, err := node.QuerySelectorAll(query)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Found all nodes with class \"test-class\": \n%v\n\n", nodes)
}
