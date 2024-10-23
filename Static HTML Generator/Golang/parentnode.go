package main

import "fmt"

// --------------------------------------------------------------------------------------------------------------------

// parentNode and parentNode methods.

// --------------------------------------------------------------------------------------------------------------------

type parentNode struct {
	htmlNode
	parents []parentNode
	leafs   []leafNode
}

// --------------------------------------------------------------------------------------------------------------------

// Convert parentNode and all child nodes into HTML.

func (p *parentNode) toHtml() (string, error) {
	if p.tag == "" {
		return "", fmt.Errorf("parentNode toHtml: no tag found - all parentNodes must have a tag")
	}
	if hasChildren(p.parents, p.leafs) {
		return "", fmt.Errorf("parentNode toHtml: no child nodes found - all parentNodes must have children. Use a leafnode for standalone nodes")
	}

	var htmlString string

	if len(p.properties) == 0 {
		htmlString += fmt.Sprintf("<%v>", p.tag)
	} else {
		htmlString += fmt.Sprintf("<%v %v>", p.tag, p.propsToHtml())
	}

	for _, node := range p.parents {
		data, err := node.toHtml()
		if err != nil {
			return "", err
		}
		htmlString += data
	}

	for _, node := range p.leafs {
		data, err := node.toHtml()
		if err != nil {
			return "", err
		}
		htmlString += data
	}

	return fmt.Sprintf("%v</%v>", htmlString, p.tag), nil
}

// --------------------------------------------------------------------------------------------------------------------

// parentNode helper functions

// --------------------------------------------------------------------------------------------------------------------

// Check for child nodes.

func hasChildren(subParents []parentNode, children []leafNode) bool {
	return len(subParents) == 0 && len(children) == 0
}

// --------------------------------------------------------------------------------------------------------------------
