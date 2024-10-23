package main

import (
	"fmt"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------

// Block and block types.

// --------------------------------------------------------------------------------------------------------------------

// Block types used for defining blocks of markdown.

// --------------------------------------------------------------------------------------------------------------------

type blockType int

const (
	paragraph blockType = iota
	heading
	blockCode
	quote
	unorderedList
	orderedList
)

// --------------------------------------------------------------------------------------------------------------------

// blocks and block methods.

// --------------------------------------------------------------------------------------------------------------------

type block struct {
	contents  string
	blockType blockType
}

// --------------------------------------------------------------------------------------------------------------------

// Turns blockNodes into htmlNodes

func (b *block) blockToHtml() (parentNode, error) {
	root := parentNode{htmlNode: htmlNode{tag: "div"}}

	var blockContents parentNode

	switch b.blockType {

	case heading:
		htmlBlock, err := convertHeadingBlock(b.contents)
		if err != nil {
			return root, nil
		}

		blockContents = htmlBlock

	case quote:
		htmlBlock, err := convertHelper("blockquote", b.contents, ">", 1)
		if err != nil {
			return root, err
		}

		blockContents = htmlBlock

	case paragraph:
		htmlBlock, err := convertHelper("p", b.contents, "", 0)
		if err != nil {
			return root, err
		}

		blockContents = htmlBlock

	case blockCode:

		htmlBlock, err := convertHelper("code", b.contents, "```", 2)
		if err != nil {
			return root, err
		}

		preNode := parentNode{htmlNode: htmlNode{tag: "pre"}}
		preNode.parents = append(preNode.parents, htmlBlock)
		blockContents = preNode

	case orderedList:
		htmlBlock, err := convertHelperLists("li", b.contents, ".")
		if err != nil {
			return root, err
		}

		listRoot := parentNode{htmlNode: htmlNode{tag: "ol"}}
		listRoot.parents = append(listRoot.parents, htmlBlock...)
		blockContents = listRoot

	case unorderedList:
		htmlBlock, err := convertHelperLists("li", b.contents, "-")
		if err != nil {
			return root, err
		}

		listRoot := parentNode{htmlNode: htmlNode{tag: "ul"}}
		listRoot.parents = append(listRoot.parents, htmlBlock...)
		blockContents = listRoot
	}

	root.parents = append(root.parents, blockContents)
	return root, nil
}

// --------------------------------------------------------------------------------------------------------------------

// block helper functions.

// --------------------------------------------------------------------------------------------------------------------

// Turns a slice of strings into a slice of block structs.

func blockToBlockType(blockStrings []string) []block {
	var newBlocks []block

	for _, subBlock := range blockStrings {
		if len(subBlock) > 1 {
			if isValidCodeBlock(subBlock) {
				newBlocks = append(newBlocks, block{contents: subBlock, blockType: blockCode})
			} else {
				switch subBlock[:1] {
				case "#":
					newBlocks = append(newBlocks, block{contents: subBlock, blockType: heading})
				case ">":
					newBlocks = append(newBlocks, block{contents: subBlock, blockType: quote})
				case "*":
					newBlocks = append(newBlocks, block{contents: subBlock, blockType: unorderedList})
				case "-":
					newBlocks = append(newBlocks, block{contents: subBlock, blockType: unorderedList})
				case ".":
					newBlocks = append(newBlocks, block{contents: subBlock, blockType: orderedList})
				default:
					newBlocks = append(newBlocks, block{contents: subBlock, blockType: paragraph})
				}
			}
		}
	}
	return newBlocks
}

// --------------------------------------------------------------------------------------------------------------------

// Turns a raw string into markdown blocks.

func markdownToBlocks(markdown string) []string {
	var previousChar string
	var newString string
	var blocks []string

	for _, char := range markdown {
		currentChar := string(char)

		if isLineBlank(previousChar, currentChar) {
			newString += currentChar
			blocks = append(blocks, newString)
			newString = ""
			previousChar = currentChar
			continue
		}

		newString += currentChar
		previousChar = currentChar
	}

	blocks = append(blocks, newString)

	for i := range blocks {
		blocks[i] = strings.TrimSpace(blocks[i])
	}

	return blocks
}

// --------------------------------------------------------------------------------------------------------------------

// Used for converting raw text strings into htmlNodes.

func convertHelper(tag, nodeContents, markdownSymbol string, replaceAmount int) (parentNode, error) {

	var blockContents parentNode

	blockContents.tag = tag
	if replaceAmount > 0 {
		nodeContents = strings.Replace(nodeContents, markdownSymbol, "", replaceAmount)
	}
	nodeContent := strings.TrimSpace(nodeContents)

	textNodes, err := textToTextNodes(nodeContent)
	if err != nil {
		return blockContents, err
	}

	for _, node := range textNodes {
		newLeaf, err := node.toHtmlNode()
		if err != nil {
			return blockContents, err
		}

		blockContents.leafs = append(blockContents.leafs, newLeaf)
	}
	return blockContents, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Converts markdown lists into htmlNodes.

func convertHelperLists(tag, nodeContents, markdownSymbol string) ([]parentNode, error) {

	stringParts := strings.Split(nodeContents, markdownSymbol)
	var scrubbedParts []string

	for _, part := range stringParts {
		if len(part) > 0 {
			scrubbed := strings.TrimSpace(part)
			scrubbedParts = append(scrubbedParts, scrubbed)
		}
	}

	var textNodes []textNode

	for _, data := range scrubbedParts {
		newTextNodes, err := textToTextNodes(data)
		if err != nil {
			return []parentNode{}, err
		}

		textNodes = append(textNodes, newTextNodes...)
	}

	var listItems []parentNode

	for _, node := range textNodes {
		newLeaf, err := node.toHtmlNode()
		if err != nil {
			return []parentNode{}, err
		}

		listEntry := parentNode{htmlNode: htmlNode{tag: tag}, leafs: []leafNode{newLeaf}}
		listItems = append(listItems, listEntry)
	}

	return listItems, nil

}

// --------------------------------------------------------------------------------------------------------------------

// Coverts heading markdown into a valid block.

func convertHeadingBlock(blockContent string) (parentNode, error) {

	headingType, tagAmount, err := func(text string) (string, int, error) {
		for i, char := range text {
			if char == ' ' {
				return fmt.Sprintf("h%v", i), i + 1, nil
			}
		}
		return " ", 0, fmt.Errorf("blockToHtml: invalid block type: %v\n(please ensure a space is between block tag and content)", text)
	}(blockContent)

	if err != nil {
		return parentNode{}, err
	}

	htmlBlock, err := convertHelper(headingType, blockContent, "#", tagAmount)
	if err != nil {
		return parentNode{}, err
	}

	return htmlBlock, nil
}

// --------------------------------------------------------------------------------------------------------------------

func isValidCodeBlock(blockString string) bool {
	return len(blockString) >= 6 && strings.Contains(blockString[:3], "```") && strings.Contains(blockString[len(blockString)-4:], "```")
}

// --------------------------------------------------------------------------------------------------------------------

func isLineBlank(previousChar, currentChar string) bool {
	return previousChar == "\n" && len(strings.TrimSpace(currentChar)) == 0
}

// --------------------------------------------------------------------------------------------------------------------

// Handler function for converting a block of raw markdown into html.

func textToHtml(rawText string) ([]string, error) {
	textBlocks := markdownToBlocks(rawText)
	markdownBlocks := blockToBlockType(textBlocks)

	var htmlStrings []string

	for _, currentBlock := range markdownBlocks {
		newNode, err := currentBlock.blockToHtml()
		if err != nil {
			return []string{}, fmt.Errorf("textToHtml: error while converting block to html: %v", err)
		}

		htmlString, err := newNode.toHtml()
		if err != nil {
			return []string{}, fmt.Errorf("textToHtml: error while parsing htmlNode: %v", err)
		}

		htmlStrings = append(htmlStrings, htmlString)
	}

	return htmlStrings, nil
}

// --------------------------------------------------------------------------------------------------------------------
