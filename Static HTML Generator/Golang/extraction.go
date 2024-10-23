package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------

// Functions used to extract and sort data.

// --------------------------------------------------------------------------------------------------------------------

// Split an HTML node into raw text and delimited type.

func splitNodeDelimiter(oldNodes []textNode, delimiter string, nodeType textNodeType) ([]textNode, error) {
	var newNodes []textNode

	for i, node := range oldNodes {
		if strings.Contains(node.text, delimiter) {
			oldNodes[i].nodeType = nodeType
		}
	}

	for _, node := range oldNodes {
		if node.nodeType != nodeType {
			newNodes = append(newNodes, node)
			continue
		}

		nodes, err := delimitText(node.text, delimiter, nodeType)
		if err != nil {
			return newNodes, err
		}

		newNodes = append(newNodes, nodes...)

	}
	return newNodes, nil
}

// --------------------------------------------------------------------------------------------------------------------

//  Keeps creating nodes of X type and raw textNodes until there is no more delimited text of that type.

func delimitText(nodeText string, delimiter string, nodeType textNodeType) ([]textNode, error) {
	var newTextNodes []textNode

	if !strings.Contains(nodeText, delimiter) {
		return newTextNodes, fmt.Errorf("splitNodeDelimiter: no delimiter in textNode: %v", nodeText)
	}

	for {
		if !strings.Contains(nodeText, delimiter) {
			if len(nodeText) > 0 {
				newTextNodes = append(newTextNodes, textNode{text: nodeText, nodeType: text})
			}
			break
		}

		before, after, _ := strings.Cut(nodeText, delimiter)
		if len(before) > 0 {
			newTextNodes = append(newTextNodes, textNode{text: before, nodeType: text})
		}

		delimitedString, remaining, found := strings.Cut(after, delimiter)
		if !found {
			return newTextNodes, fmt.Errorf("splitNodeDelimiter: delimiter has not been closed: %v", after)
		}

		newTextNodes = append(newTextNodes, textNode{text: delimitedString, nodeType: nodeType})
		nodeText = remaining
	}
	return newTextNodes, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Extracts images or urls out of textNodes.

func splitNodesForLinksImages(oldNodes []textNode, newNodeType textNodeType) ([]textNode, error) {
	var newNodes []textNode

	for _, old := range oldNodes {
		var links []string

		if newNodeType == image {
			links = extractMarkdownImages(old.text)
		} else {
			links = extractMarkdownLinks(old.text)
		}

		if len(links) == 0 {
			newNodes = append(newNodes, old)

		} else {
			extractedNodes, err := extractImageOrLink(old.text, links, newNodeType)
			if err != nil {
				return newNodes, err
			}

			newNodes = append(newNodes, extractedNodes...)
		}
	}
	return newNodes, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Extracts all image data from a string

func extractMarkdownImages(text string) []string {
	re := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	return re.FindAllString(text, -1)
}

// --------------------------------------------------------------------------------------------------------------------

// Extracts all link data from a string.

func extractMarkdownLinks(text string) []string {
	re := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	return re.FindAllString(text, -1)
}

// --------------------------------------------------------------------------------------------------------------------

// Extracts images and links from strings and turns them into textNodes.

func extractImageOrLink(currentNodeText string, nodeStrings []string, newNodeType textNodeType) ([]textNode, error) {
	raw := currentNodeText
	var newNodes []textNode

	for _, item := range nodeStrings {
		sections := strings.Split(raw, item)

		if len(sections) != 2 {
			return newNodes, fmt.Errorf("splitNodeLinks - invalid markdown passed to function: %v", sections)
		}

		if len(sections[0]) != 0 {
			newNodes = append(newNodes, textNode{text: sections[0], nodeType: text})
		}

		imageLinkStrings := getImageOrLink(item)
		newNodes = append(newNodes, textNode{text: imageLinkStrings[0], url: imageLinkStrings[1], nodeType: newNodeType})
		raw = sections[1]
	}

	if len(raw) > 0 {
		newNodes = append(newNodes, textNode{text: raw, nodeType: text})
	}

	return newNodes, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Pulls an image src or url from a string.

func getImageOrLink(word string) []string {
	words := regexp.MustCompile("[(*?)]").Split(word, -1)

	var assets []string

	for i := range words {
		if len(words[i]) >= 1 {
			toAppend := strings.TrimSpace(words[i])
			assets = append(assets, toAppend)
		}
	}

	return assets
}

// --------------------------------------------------------------------------------------------------------------------

// Checks for and finds html elements from a list of html strings.

func extractTitleHtml(htmlStrings []string) (string, error) {
	for _, htmlItem := range htmlStrings {
		if strings.Contains(htmlItem, "h1") {
			htmlItem = strings.Replace(htmlItem, "<div>", "", 1)
			htmlItem = strings.Replace(htmlItem, "</div>", "", 1)
			htmlItem = strings.Replace(htmlItem, "<h1>", "", 1)
			htmlItem = strings.Replace(htmlItem, "</h1>", "", 1)
			return htmlItem, nil
		}
	}

	return "", errors.New("extractTitleHtml: no h1 found in html elements")
}

// --------------------------------------------------------------------------------------------------------------------
