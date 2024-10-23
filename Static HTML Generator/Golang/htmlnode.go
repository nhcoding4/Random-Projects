package main

import "fmt"

// --------------------------------------------------------------------------------------------------------------------

// htmlNode

// --------------------------------------------------------------------------------------------------------------------

type htmlNode struct {
	tag        string
	properties map[string]string
}

// --------------------------------------------------------------------------------------------------------------------

// Converts the properties in the htmlNode map into an HTML string.

func (h *htmlNode) propsToHtml() string {
	var htmlString string
	i := 0

	for key, value := range h.properties {
		htmlString += fmt.Sprintf("%v=\"%v\"", key, value)
		i++
		if i < len(h.properties) {
			htmlString += " "
		}
	}
	return htmlString
}

// --------------------------------------------------------------------------------------------------------------------
