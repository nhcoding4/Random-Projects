package main

import "fmt"

// --------------------------------------------------------------------------------------------------------------------

// leafNode and leafNode methods.

// --------------------------------------------------------------------------------------------------------------------

type leafNode struct {
	htmlNode
	value string
}

// --------------------------------------------------------------------------------------------------------------------

// Convert LeafNode into HTML.

func (l *leafNode) toHtml() (string, error) {
	if l.value == "" {
		return "", fmt.Errorf("leafNode toHtml: leafNode has no value - all leafnodes require " +
			"a value")
	}

	if l.tag == "img" {
		return fmt.Sprintf("<%v %v>", l.tag, l.propsToHtml()), nil
	}

	if l.tag == "" {
		if len(l.properties) == 0 {
			return l.value, nil
		} else {
			return fmt.Sprintf("%v %v", l.value, l.propsToHtml()), nil
		}
	}

	if len(l.properties) == 0 {
		return fmt.Sprintf("<%v>%v</%v>", l.tag, l.value, l.tag), nil
	} else {
		return fmt.Sprintf("<%v %v>%v</%v>", l.tag, l.propsToHtml(), l.value, l.tag), nil
	}
}

// --------------------------------------------------------------------------------------------------------------------
