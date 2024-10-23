package main

// --------------------------------------------------------------------------------------------------------------------

// Types and Functions Related to textNodes

// --------------------------------------------------------------------------------------------------------------------

// Custom type used for denoting the type of textNode.

// --------------------------------------------------------------------------------------------------------------------

type textNodeType int

const (
	text textNodeType = iota
	bold
	italic
	code
	link
	image
)

// --------------------------------------------------------------------------------------------------------------------

// TextNode and TextNode methods.

// --------------------------------------------------------------------------------------------------------------------

type textNode struct {
	text     string
	url      string
	nodeType textNodeType
}

// --------------------------------------------------------------------------------------------------------------------

// Creates a htmlNode from a textNode.

func (t *textNode) toHtmlNode() (leafNode, error) {
	var returnNode leafNode

	switch t.nodeType {
	case text:
		returnNode = leafNode{value: t.text}
	case bold:
		returnNode = leafNode{value: t.text, htmlNode: htmlNode{tag: "b"}}
	case italic:
		returnNode = leafNode{value: t.text, htmlNode: htmlNode{tag: "i"}}
	case code:
		returnNode = leafNode{value: t.text, htmlNode: htmlNode{tag: "code"}}
	case link:
		returnNode = leafNode{value: t.text, htmlNode: htmlNode{tag: "a", properties: map[string]string{"href": t.url}}}
	case image:
		returnNode = leafNode{value: t.text, htmlNode: htmlNode{
			tag: "img", properties: map[string]string{"src": t.url, "alt": t.text}}}
	}

	return returnNode, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Helper functions related to textNodes.

// --------------------------------------------------------------------------------------------------------------------

// Turns a string into an array of textNodes by running the node through each function.

func textToTextNodes(userInput string) ([]textNode, error) {
	newTextNode := textNode{text: userInput, nodeType: text}

	nodeSlice, err := splitNodeDelimiter([]textNode{newTextNode}, "**", bold)
	if err != nil {
		return []textNode{}, err
	}

	nodeSlice, err = splitNodeDelimiter(nodeSlice, "*", italic)
	if err != nil {
		return []textNode{}, err
	}

	nodeSlice, err = splitNodeDelimiter(nodeSlice, "`", code)
	if err != nil {
		return []textNode{}, err
	}

	nodeSlice, err = splitNodesForLinksImages(nodeSlice, image)
	if err != nil {
		return []textNode{}, err
	}

	nodeSlice, err = splitNodesForLinksImages(nodeSlice, link)
	if err != nil {
		return []textNode{}, err
	}

	return nodeSlice, nil
}

// --------------------------------------------------------------------------------------------------------------------
