package selector

import (
	"golang.org/x/net/html"
	"strings"
)

// Inherits directly from Node in golang.org/x/net/html
type Node struct {
	*html.Node
}

// Represents a key-value attribute of a Node
type Attribute struct {
	Key   string
	Value string
}

// Used to conduct a search for a node.
type Query struct {
	Tag        string
	Class      string
	Id         string
	Attributes []Attribute
}

// Returns Node used to call QuerySelector
func NewNode(n *html.Node) *Node {
	return &Node{n}
}

// Returns string representation of the Node
func (n *Node) String() (output string) {
	if n.Type == html.TextNode {
		return strings.Trim(n.Data, " \r\n\t")
	}

	output = "<" + n.Data + ""

	for _, attr := range n.Attr {
		output += " " + attr.Key + "=\"" + attr.Val + "\""
	}

	return output + ">" + n.Inner() + "</" + strings.Trim(n.Data, " \n\t\r") + ">"
}

// Returns string representation of the inner content of the Node
func (n *Node) Inner() (output string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		child := &Node{c}
		output += child.String()
	}

	return output
}

// Retrives an attribute of the node by name
func (n *Node) getAttribute(name string) string {
	for _, attr := range n.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

// Checks if the node passes the given query
func (n *Node) passesQuery(query *Query) bool {
	if n.Type != html.ElementNode {
		return false
	}

	if query.Tag != "" && n.Data != query.Tag {
		return false
	}

	// Currently only supports a single class name in query
	if query.Class != "" {
		classes := strings.Split(n.getAttribute("class"), " ")
		matches := false

		for _, class := range classes {
			if class == query.Class {
				matches = true
			}
		}

		if !matches {
			return false
		}
	}

	// Check if ID matches
	if query.Id != "" {
		id := n.getAttribute("id")

		if query.Id != id {
			return false
		}
	}

	// Check that all Attributes check out
	for _, attr := range query.Attributes {
		if n.getAttribute(attr.Key) != attr.Value {
			return false
		}
	}

	return true
}

// Searches all children of Node and returns first one that satisfies the query.
func (n *Node) QuerySelector(query *Query) (*Node, error) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		child := &Node{c}
		if child.passesQuery(query) {
			return child, nil
		}

		node, err := child.QuerySelector(query)
		if err != nil || node == nil {
			continue
		}

		return node, nil
	}

	return nil, nil
}

// Searches all children of Node and returns all nodes satisfying the query.
func (n *Node) QuerySelectorAll(query *Query) ([]*Node, error) {
	results := make([]*Node, 0)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		child := &Node{c}
		if child.passesQuery(query) {
			results = append(results, child)
		}

		nodes, err := child.QuerySelectorAll(query)
		if err != nil {
			continue
		}

		results = append(results, nodes...)
	}

	return results, nil
}
